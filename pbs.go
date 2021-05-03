package pbs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func init() {
	fileForLog, err := os.OpenFile(LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.SetOutput(fileForLog)
	log.SetFlags(log.Ldate | log.Lmicroseconds)
	log.Println("in init() of package pbs")
}

func ParseTime(node_name string, map_ptr *map[string]interface{}) time.Time {

	qtime, ok := (*map_ptr)[node_name]
	if !ok {
		log.Println(" can't find "+node_name+" node. Set time = ", time.Now())
		return time.Now()
	}

	location, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		log.Fatal(err)
	}
	timeFormatFrom := "Mon Jan  2 15:04:05 2006"
	t, err := time.ParseInLocation(timeFormatFrom, qtime.(string), location)
	if err != nil {
		log.Println(err)
	}
	return t
}

// This function parse the output of qstat which look like:
// Job id            Name             User              Time Use S Queue
// ----------------  ---------------- ----------------  -------- - -----
// 2008.g3           STDIN            ivan                     0 Q workq
// 2009.g3           STDIN            ivan                     0 Q workq
// 2010.g3           STDIN            ivan                     0 Q workq
func ParseQSTATTable(str string) *map[int]int {
	log.Println(" In ParseQSTATTable")

	// The map is used here because finding an element in it has complexity O(1).
	map_of_id := make(map[int]int)
	for i, line := range strings.Split(strings.TrimSuffix(str, "\n"), "\n") {
		if i >= 2 {
			string_arr := strings.Split(line, ".")
			if num, err := strconv.Atoi(strings.TrimSpace(string_arr[0])); err == nil {
				map_of_id[num] = num
			} else {
				log.Println(err)
			}
		}
	}
	return &map_of_id
}

// This function parese the stderr string after running qstat and update a tasks
// qstat: Unknown Job Id 2002.g3
// qstat: Unknown Job Id 2003.g3
func ParseStdErr(stderr string, tasks_map *map[int]*Task) {
	log.Printf(" In ParseStdErr. stderr = %v", stderr)
	if len(stderr) == 0 {
		return
	}
	for _, line := range strings.Split(strings.TrimSuffix(stderr, "\n"), "\n") {
		string_arr := strings.Split(line, " ")
		id_str := strings.Split(string_arr[4], ".")
		if num, err := strconv.Atoi(id_str[0]); err == nil {
			(*tasks_map)[num].Status = Completed
			log.Printf(" In ParseStdErr task with process id = %v is completed", num)
		} else {
			log.Println(err)
		}
	}
}

// Parse the output of the command <qstat -f -F json 1000>.
func ParseQSTAT(json_map *map[string]interface{}, tasks_map *map[int]*Task) {

	log.Println(" In function ParseQSTAT")

	jobs, ok := (*json_map)["Jobs"].(map[string]interface{})
	if !ok {
		log.Println(" can't finde Jobs node")
		return
	}

	for key := range jobs {
		qot_calc := jobs[key].(map[string]interface{})

		str_arry := strings.Split(key, ".")
		process_id, err := strconv.Atoi(str_arry[0])
		if err != nil {
			log.Println(err)
			continue
		}
		task, ok := (*tasks_map)[process_id]
		if !ok {
			log.Printf(" the task with process_id = %v not found\n", process_id)
			continue
		}

		job_state, ok := qot_calc["job_state"]
		if !ok {
			log.Println(" can't find job_state node")
		}
		// For more detail see https://www.altair.com/pdfs/pbsworks/PBSReferenceGuide2020.1.pdf [pp. 189].
		switch job_state {
		case "B": // Job array has started execution
			task.Status = Running
		case "E": // The Exiting state
			task.Status = Running
		case "F": // The Finished stat
			task.Status = Completed
		case "H": // The Held state
			task.Status = Queued
		case "M": // The Moved state
			task.Status = Running
		case "Q": // The Queued state
			task.Status = Queued
		case "R": // The Running state
			task.Status = Running
		case "S": // The Suspended state
			task.Status = Running
		case "T": // The Transiting state
			task.Status = Running
		case "U": // Job suspended due to workstation user activity
			task.Status = Stopped
		case "W": // The Waiting state
			task.Status = Queued
		case "X": // The eXited state. Subjobs only
			task.Status = Unknown
		default:
			task.Status = Unknown
			log.Println(" unknown job state")
		}

		queue, ok := qot_calc["queue"]
		if !ok {
			log.Println(" can't find queue node")
		}
		switch queue {
		case "atom":
			task.ServiceId = 1
		case "atom-sim":
			task.ServiceId = 2
		case "linopt":
			task.ServiceId = 3
		case "linopt-gates":
			task.ServiceId = 4
		case "linopt-sim":
			task.ServiceId = 5
		}
		task.CreationTime = ParseTime("qtime", &qot_calc)
		task.UpdateTime = ParseTime("mtime", &qot_calc)
		task.LaunchTime = ParseTime("stime", &qot_calc)

		// find PYTHON_SCRIPT
		variable_list, ok := qot_calc["Variable_List"].(map[string]interface{})
		if !ok {
			log.Println("can't find Variable_List node")
		}
		python_script, ok := variable_list["PYTHON_SCRIPT"]
		if !ok {
			log.Println("can't find PYTHON_SCRIPT node")
		}
		task.FileName = python_script.(string)
	}
}

// Parse the result of the  qsub command.
func ParseQSUB(str string, task *Task) {
	log.Printf(" In ParseQSUB string = %v", str)
	string_arr := strings.Split(str, ".")
	if num, err := strconv.Atoi(string_arr[0]); err == nil {
		task.ProcessId = num
	} else {
		log.Println(err)
	}
}

//
func RunQSTATCommand(tasks_map *map[int]*Task, cmdToRun string, args []string) error {
	log.Println(" RunQSTATCommand cmdToRun = ", cmdToRun, args)
	if len(*tasks_map) == 0 { // Nothing to be done for empty list of tasks
		return nil
	}
	cmd := exec.Command(cmdToRun, args...)
	stdout, err := cmd.StdoutPipe()
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err != nil {
		log.Println("cmd.StdoutPipe()", err)
		return err
	}
	if err := cmd.Start(); err != nil {
		log.Println("cmd.Start()", err)
		return err
	}

	var stdout_json map[string]interface{}
	if err := json.NewDecoder(stdout).Decode(&stdout_json); err != nil {
		log.Println("json.NewDecoder", err)
		return err
	}

	if err := cmd.Wait(); err != nil {
		log.Println(err)
	}
	ParseStdErr(stderr.String(), tasks_map)
	ParseQSTAT(&stdout_json, tasks_map)
	return nil
}

// This function will run "qstat" command.
func GetAllQSTATTasks(cmdToRun string, args ...string) (*map[int]int, bool) {
	log.Println(" GetAllQSTATTasks cmdToRun = ", cmdToRun)
	empty_map := make(map[int]int)
	cmd := exec.Command(cmdToRun, args...)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stdout

	if err := cmd.Start(); err != nil {
		log.Println(err)
		return &empty_map, false
	}

	if err := cmd.Wait(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				log.Printf("Exit status: %d in run qsub command = %v", status.ExitStatus(), cmdToRun)
				return &empty_map, false
			}
		} else {
			log.Printf("cmd.Wait: %v", err)
			return &empty_map, false
		}
	}

	log.Println(" in GetAllQSTATTasks stdout = ", stdout.String())
	map_with_all_queue := ParseQSTATTable(stdout.String())
	return map_with_all_queue, true
}

// This command starts script execution and assigns the required ProcessId to the task.
func RunQSUBCommand(task *Task, cmdToRun string, args ...string) error {
	log.Println(" In func RunQSUBCommand() cmdToRun = ", cmdToRun, args)
	cmd := exec.Command(cmdToRun, args...)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stdout

	if err := cmd.Start(); err != nil {
		log.Println(err)
		return err
	}

	if err := cmd.Wait(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				log.Printf("Exit status: %d in run qsub command with args %v %v", status.ExitStatus(), cmdToRun, args)
				return exiterr
			}
		} else {
			log.Printf("cmd.Wait: %v", err)
			return err
		}
	}

	log.Println(" in RunQSUBCommand stdout = ", stdout.String())
	ParseQSUB(stdout.String(), task)
	log.Println(" in qsub. Task.ProcessId = ", task.ProcessId)
	return nil
}

// Delete a task with process id = processID from pbs queue.
func RunQDELCommand(cmdToRun string, processId int) error {
	log.Println(" in func RunQDELCommand() cmdToRun = ", cmdToRun, "processId = ", processId)
	cmd := exec.Command(cmdToRun, strconv.Itoa(processId))
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stdout

	if err := cmd.Start(); err != nil {
		log.Println(err)
		return err
	}

	if err := cmd.Wait(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				log.Printf("Exit status: %d for run qdel command with processID = %v", status.ExitStatus(), processId)
				return exiterr
			}
		} else {
			log.Printf("cmd.Wait: %v", err)
			return err
		}
	}
	log.Println(" RunQDELCommand completed with stdout & stderr = ", stdout.String())
	return nil
}
