package pbs

import (
	"encoding/json"
	"log"
	"testing"
	"time"
)

var json_map map[string]interface{}
var location *time.Location
var json_map_multi_tasks map[string]interface{}
var json_map_selected_tasks map[string]interface{}

func init() {

	// set global location
	location, _ = time.LoadLocation("Europe/Moscow")
	if _, err := time.LoadLocation("Europe/Moscow"); err != nil {
		log.Fatal(err)
	}

	// prepare a data for one task
	empJson := `
	{
		"timestamp":1618432467,
		"pbs_version":"20.0.1",
		"pbs_server":"qot-calc0",
		"Jobs":{
			"1008.qot-calc0":{
				"Job_Name":"ntest",
				"Job_Owner":"qoot@qot-calc0",
				"resources_used":{
					"cpupercent":0,
					"cput":"00:00:00",
					"mem":"0kb",
					"ncpus":5,
					"vmem":"0kb",
					"walltime":"00:00:00"
				},
				"job_state":"R",
				"queue":"atom-sim",
				"server":"qot-calc0",
				"Checkpoint":"u",
				"ctime":"Wed Apr 14 23:34:17 2021",
				"Error_Path":"qot-calc0:/home/qoot/ntest.e1008",
				"exec_host":"qot-calc0/10+qot-calc0/11+qot-calc0/12+qot-calc0/13+qot-calc0/14",
				"exec_vnode":"(qot-calc0:ncpus=1)+(qot-calc0:ncpus=1)+(qot-calc0:ncpus=1)+(qot-calc0:ncpus=1)+(qot-calc0:ncpus=1)",
				"Hold_Types":"n",
				"Join_Path":"n",
				"Keep_Files":"n",
				"Mail_Points":"a",
				"mtime":"Wed Apr 14 23:34:18 2021",
				"Output_Path":"qot-calc0:/home/qoot/ntest.o1008",
				"Priority":0,
				"qtime":"Wed Apr 14 23:34:17 2021",
				"Rerunable":"True",
				"Resource_List":{
					"ncpus":5,
					"nodect":5,
					"place":"free",
					"select":"5:ncpus=1",
					"walltime":"04:00:00"
				},
				"stime":"Wed Apr 14 23:34:19 2021",
				"session_id":16971,
				"jobdir":"/home/qoot",
				"substate":42,
				"Variable_List":{
					"PBS_O_HOME":"/home/qoot",
					"PBS_O_LANG":"en_US.Utime_format-8",
					"PBS_O_LOGNAME":"qoot",
					"PBS_O_PATH":"/home/qoot/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/snap/bin:/opt/pbs/bin",
					"PBS_O_SHELL":"/bin/bash",
					"PBS_O_WORKDIR":"/home/qoot",
					"PBS_O_SYSTEM":"Linux",
					"PYTHON_SCRIPT":"myscript.py",
					"PBS_O_QUEUE":"default",
					"PBS_O_HOST":"qot-calc0"
				},
				"comment":"Job run at Wed Apr 14 at 23:34 on (qot-calc0:ncpus=1)+(qot-calc0:ncpus=1)+(qot-calc0:ncpus=1)+(qot-calc0:ncpus=1)+(qot-calc0:ncpus=1)",
				"etime":"Wed Apr 14 23:34:17 2021",
				"run_count":1,
				"Submit_arguments":"-v PYTHON_SCRIPT=myscript.py test.pbs",
				"project":"_pbs_project_default",
				"Submit_Host":"qot-calc0"
			}
		}
	}
	`
	json.Unmarshal([]byte(empJson), &json_map)

	// prepare a data for all tasks
	empJson = `
	{
		"timestamp":1618431756,
		"pbs_version":"20.0.1",
		"pbs_server":"qot-calc0",
		"Jobs":{
			"1006.qot-calc0":{
				"Job_Name":"ntest",
				"Job_Owner":"qoot@qot-calc0",
				"resources_used":{
					"cpupercent":0,
					"cput":"00:00:00",
					"mem":"12272kb",
					"ncpus":5,
					"vmem":"19796kb",
					"walltime":"00:01:50"
				},
				"job_state":"R",
				"queue":"atom",
				"server":"qot-calc0",
				"Checkpoint":"u",
				"ctime":"Wed Apr 14 23:20:07 2021",
				"Error_Path":"qot-calc0:/home/qoot/ntest.e1006",
				"exec_host":"qot-calc0/0+qot-calc0/1+qot-calc0/2+qot-calc0/3+qot-calc0/4",
				"exec_vnode":"(qot-calc0:ncpus=1)+(qot-calc0:ncpus=1)+(qot-calc0:ncpus=1)+(qot-calc0:ncpus=1)+(qot-calc0:ncpus=1)",
				"Hold_Types":"n",
				"Join_Path":"n",
				"Keep_Files":"n",
				"Mail_Points":"a",
				"mtime":"Wed Apr 14 23:21:57 2021",
				"Output_Path":"qot-calc0:/home/qoot/ntest.o1006",
				"Priority":0,
				"qtime":"Wed Apr 14 23:20:07 2021",
				"Rerunable":"True",
				"Resource_List":{
					"ncpus":5,
					"nodect":5,
					"place":"free",
					"select":"5:ncpus=1",
					"walltime":"04:00:00"
				},
				"stime":"Wed Apr 14 23:20:07 2021",
				"session_id":16769,
				"jobdir":"/home/qoot",
				"substate":42,
				"Variable_List":{
					"PBS_O_HOME":"/home/qoot",
					"PBS_O_LANG":"en_US.Utime_format-8",
					"PBS_O_LOGNAME":"qoot",
					"PBS_O_PATH":"/home/qoot/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/snap/bin:/opt/pbs/bin",
					"PBS_O_SHELL":"/bin/bash",
					"PBS_O_WORKDIR":"/home/qoot",
					"PBS_O_SYSTEM":"Linux",
					"PYTHON_SCRIPT":"myscript1.py",
					"PBS_O_QUEUE":"default",
					"PBS_O_HOST":"qot-calc0"
				},
				"comment":"Job run at Wed Apr 14 at 23:20 on (qot-calc0:ncpus=1)+(qot-calc0:ncpus=1)+(qot-calc0:ncpus=1)+(qot-calc0:ncpus=1)+(qot-calc0:ncpus=1)",
				"etime":"Wed Apr 14 23:20:07 2021",
				"run_count":1,
				"Submit_arguments":"test.pbs",
				"project":"_pbs_project_default",
				"Submit_Host":"qot-calc0"
			},
			"1007.qot-calc0":{
				"Job_Name":"ntest",
				"Job_Owner":"qoot@qot-calc0",
				"resources_used":{
					"cpupercent":0,
					"cput":"00:00:00",
					"mem":"0kb",
					"ncpus":5,
					"vmem":"0kb",
					"walltime":"00:00:00"
				},
				"job_state":"R",
				"queue":"linopt-sim",
				"server":"qot-calc0",
				"Checkpoint":"u",
				"ctime":"Wed Apr 14 23:22:31 2021",
				"Error_Path":"qot-calc0:/home/qoot/ntest.e1007",
				"exec_host":"qot-calc0/5+qot-calc0/6+qot-calc0/7+qot-calc0/8+qot-calc0/9",
				"exec_vnode":"(qot-calc0:ncpus=1)+(qot-calc0:ncpus=1)+(qot-calc0:ncpus=1)+(qot-calc0:ncpus=1)+(qot-calc0:ncpus=1)",
				"Hold_Types":"n",
				"Join_Path":"n",
				"Keep_Files":"n",
				"Mail_Points":"a",
				"mtime":"Wed Apr 14 23:22:31 2021",
				"Output_Path":"qot-calc0:/home/qoot/ntest.o1007",
				"Priority":0,
				"qtime":"Wed Apr 14 23:22:31 2021",
				"Rerunable":"True",
				"Resource_List":{
					"ncpus":5,
					"nodect":5,
					"place":"free",
					"select":"5:ncpus=1",
					"walltime":"04:00:00"
				},
				"stime":"Wed Apr 14 23:22:31 2021",
				"session_id":16829,
				"jobdir":"/home/qoot",
				"substate":42,
				"Variable_List":{
					"PBS_O_HOME":"/home/qoot",
					"PBS_O_LANG":"en_US.Utime_format-8",
					"PBS_O_LOGNAME":"qoot",
					"PBS_O_PATH":"/home/qoot/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/snap/bin:/opt/pbs/bin",
					"PBS_O_SHELL":"/bin/bash",
					"PBS_O_WORKDIR":"/home/qoot",
					"PBS_O_SYSTEM":"Linux",
					"PYTHON_SCRIPT":"myscript2.py",
					"PBS_O_QUEUE":"default",
					"PBS_O_HOST":"qot-calc0"
				},
				"comment":"Job run at Wed Apr 14 at 23:22 on (qot-calc0:ncpus=1)+(qot-calc0:ncpus=1)+(qot-calc0:ncpus=1)+(qot-calc0:ncpus=1)+(qot-calc0:ncpus=1)",
				"etime":"Wed Apr 14 23:22:31 2021",
				"run_count":1,
				"Submit_arguments":"test.pbs",
				"project":"_pbs_project_default",
				"Submit_Host":"qot-calc0"
			}
		}
	}`
	json.Unmarshal([]byte(empJson), &json_map_multi_tasks)

	// prepare a data for selected tasks
	empJson = `
	{
		"timestamp":1619121423,
		"pbs_version":"20.0.1",
		"pbs_server":"vm-priboy",
		"Jobs":{
			"2007.vm-priboy":{
				"Job_Name":"nsleep",
				"Job_Owner":"priboy@vm-priboy",
				"resources_used":{
					"cpupercent":0,
					"cput":"00:00:00",
					"mem":"6436kb",
					"ncpus":1,
					"vmem":"10508kb",
					"walltime":"00:00:12"
				},
				"job_state":"R",
				"queue":"atom-sim",
				"server":"vm-priboy",
				"Checkpoint":"u",
				"ctime":"Thu Apr 22 22:56:43 2021",
				"Error_Path":"vm-priboy:/home/priboy/nsleep.e2007",
				"exec_host":"vm-priboy/0",
				"exec_vnode":"(vm-priboy:ncpus=1)",
				"Hold_Types":"n",
				"Join_Path":"n",
				"Keep_Files":"n",
				"Mail_Points":"a",
				"mtime":"Thu Apr 22 22:56:55 2021",
				"Output_Path":"vm-priboy:/home/priboy/nsleep.o2007",
				"Priority":0,
				"qtime":"Thu Apr 22 22:56:43 2021",
				"Rerunable":"True",
				"Resource_List":{
					"ncpus":1,
					"nodect":1,
					"place":"pack",
					"select":"1:ncpus=1"
				},
				"stime":"Thu Apr 22 22:56:43 2021",
				"session_id":38459,
				"jobdir":"/home/priboy",
				"substate":42,
				"Variable_List":{
					"PBS_O_HOME":"/home/priboy",
					"PBS_O_LANG":"en_US.Utime_format-8",
					"PBS_O_LOGNAME":"priboy",
					"PBS_O_PATH":"/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/snap/bin:/opt/pbs/bin",
					"PBS_O_SHELL":"/bin/bash",
					"PBS_O_WORKDIR":"/home/priboy",
					"PYTHON_SCRIPT":"myscript1.py",
					"PBS_O_SYSTEM":"Linux",
					"PBS_O_QUEUE":"workq",
					"PBS_O_HOST":"vm-priboy"
				},
				"comment":"Job run at Thu Apr 22 at 22:56 on (vm-priboy:ncpus=1)",
				"etime":"Thu Apr 22 22:56:43 2021",
				"run_count":1,
				"Submit_arguments":"sleep.pbs",
				"project":"_pbs_project_default",
				"Submit_Host":"vm-priboy"
			},
			"2009.vm-priboy":{
				"Job_Name":"nsleep",
				"Job_Owner":"priboy@vm-priboy",
				"resources_used":{
					"cpupercent":0,
					"cput":"00:00:00",
					"mem":"6624kb",
					"ncpus":1,
					"vmem":"10508kb",
					"walltime":"00:00:10"
				},
				"job_state":"R",
				"queue":"atom-sim",
				"server":"vm-priboy",
				"Checkpoint":"u",
				"ctime":"Thu Apr 22 22:56:45 2021",
				"Error_Path":"vm-priboy:/home/priboy/nsleep.e2009",
				"exec_host":"vm-priboy/2",
				"exec_vnode":"(vm-priboy:ncpus=1)",
				"Hold_Types":"n",
				"Join_Path":"n",
				"Keep_Files":"n",
				"Mail_Points":"a",
				"mtime":"Thu Apr 22 22:56:55 2021",
				"Output_Path":"vm-priboy:/home/priboy/nsleep.o2009",
				"Priority":0,
				"qtime":"Thu Apr 22 22:56:45 2021",
				"Rerunable":"True",
				"Resource_List":{
					"ncpus":1,
					"nodect":1,
					"place":"pack",
					"select":"1:ncpus=1"
				},
				"stime":"Thu Apr 22 22:56:45 2021",
				"session_id":38483,
				"jobdir":"/home/priboy",
				"substate":42,
				"Variable_List":{
					"PBS_O_HOME":"/home/priboy",
					"PBS_O_LANG":"en_US.Utime_format-8",
					"PBS_O_LOGNAME":"priboy",
					"PBS_O_PATH":"/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/snap/bin:/opt/pbs/bin",
					"PBS_O_SHELL":"/bin/bash",
					"PBS_O_WORKDIR":"/home/priboy",
					"PYTHON_SCRIPT":"myscript2.py",
					"PBS_O_SYSTEM":"Linux",
					"PBS_O_QUEUE":"workq",
					"PBS_O_HOST":"vm-priboy"
				},
				"comment":"Job run at Thu Apr 22 at 22:56 on (vm-priboy:ncpus=1)",
				"etime":"Thu Apr 22 22:56:45 2021",
				"run_count":1,
				"Submit_arguments":"sleep.pbs",
				"project":"_pbs_project_default",
				"Submit_Host":"vm-priboy"
			},
			"2010.vm-priboy":{
				"Job_Name":"nsleep",
				"Job_Owner":"priboy@vm-priboy",
				"resources_used":{
					"cpupercent":0,
					"cput":"00:00:00",
					"mem":"6748kb",
					"ncpus":1,
					"vmem":"10508kb",
					"walltime":"00:00:10"
				},
				"job_state":"R",
				"queue":"linopt-gates",
				"server":"vm-priboy",
				"Checkpoint":"u",
				"ctime":"Thu Apr 22 22:56:45 2021",
				"Error_Path":"vm-priboy:/home/priboy/nsleep.e2010",
				"exec_host":"vm-priboy/3",
				"exec_vnode":"(vm-priboy:ncpus=1)",
				"Hold_Types":"n",
				"Join_Path":"n",
				"Keep_Files":"n",
				"Mail_Points":"a",
				"mtime":"Thu Apr 22 22:56:55 2021",
				"Output_Path":"vm-priboy:/home/priboy/nsleep.o2010",
				"Priority":0,
				"qtime":"Thu Apr 22 22:56:45 2021",
				"Rerunable":"True",
				"Resource_List":{
					"ncpus":1,
					"nodect":1,
					"place":"pack",
					"select":"1:ncpus=1"
				},
				"stime":"Thu Apr 22 22:56:45 2021",
				"session_id":38495,
				"jobdir":"/home/priboy",
				"substate":42,
				"Variable_List":{
					"PBS_O_HOME":"/home/priboy",
					"PBS_O_LANG":"en_US.Utime_format-8",
					"PBS_O_LOGNAME":"priboy",
					"PBS_O_PATH":"/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/snap/bin:/opt/pbs/bin",
					"PBS_O_SHELL":"/bin/bash",
					"PBS_O_WORKDIR":"/home/priboy",
					"PYTHON_SCRIPT":"myscript3.py",
					"PBS_O_SYSTEM":"Linux",
					"PBS_O_QUEUE":"workq",
					"PBS_O_HOST":"vm-priboy"
				},
				"comment":"Job run at Thu Apr 22 at 22:56 on (vm-priboy:ncpus=1)",
				"etime":"Thu Apr 22 22:56:45 2021",
				"run_count":1,
				"Submit_arguments":"sleep.pbs",
				"project":"_pbs_project_default",
				"Submit_Host":"vm-priboy"
			}
		}
	}`
	json.Unmarshal([]byte(empJson), &json_map_selected_tasks)
}

func CompareTasks(expectedTask *Task, actualTask *Task, t *testing.T) {
	if expectedTask.ProcessId != actualTask.ProcessId {
		t.Errorf("Expected Sid  != Actual Sid  [%v != %v]", expectedTask.ProcessId, actualTask.ProcessId)
	}
	if expectedTask.FileName != actualTask.FileName {
		t.Errorf("Expected FileName != Actual FileName  [%v != %v]", expectedTask.FileName, actualTask.FileName)
	}
	if expectedTask.Status != actualTask.Status {
		t.Errorf("Expected Status != Actual Status  [%v != %v]", expectedTask.Status, actualTask.Status)
	}
	if !expectedTask.CreationTime.Equal(actualTask.CreationTime) {
		t.Errorf("Expected CreationTime != Actual CreationTime  [%v != %v]",
			expectedTask.CreationTime, actualTask.CreationTime)
	}
	if !expectedTask.UpdateTime.Equal(actualTask.UpdateTime) {
		t.Errorf("Expected UpdateTime != Actual UpdateTime  [%v != %v]", expectedTask.UpdateTime, actualTask.UpdateTime)
	}
	if !expectedTask.LaunchTime.Equal(actualTask.LaunchTime) {
		t.Errorf("Expected LaunchTime != Actual LaunchTime  [%v != %v]", expectedTask.LaunchTime, actualTask.LaunchTime)
	}
	if expectedTask.ServiceId != actualTask.ServiceId {
		t.Errorf("Expected ServiceId != Actual ServiceId  [%v != %v]", expectedTask.ServiceId, actualTask.ServiceId)
	}
}

func TestParseQSTATAllOutput(t *testing.T) {
	var time_format = "2006-01-02 15:04:05"
	tasks_map := make(map[int]*Task)
	// prepare first task
	qtime, _ := time.ParseInLocation(time_format, "2021-04-14 23:20:07", location)
	mtime, _ := time.ParseInLocation(time_format, "2021-04-14 23:21:57", location)
	stime, _ := time.ParseInLocation(time_format, "2021-04-14 23:20:07", location)
	actualTask1 := Task{ProcessId: 1006}
	expectedTask1 := Task{
		ProcessId:    1006,
		FileName:     "myscript1.py",
		Status:       "running",
		CreationTime: qtime,
		UpdateTime:   mtime,
		LaunchTime:   stime,
		ServiceId:    1}

	qtime, _ = time.ParseInLocation(time_format, "2021-04-14 23:22:31", location)
	mtime, _ = time.ParseInLocation(time_format, "2021-04-14 23:22:31", location)
	stime, _ = time.ParseInLocation(time_format, "2021-04-14 23:22:31", location)
	actualTask2 := Task{ProcessId: 1007}
	expectedTask2 := Task{
		ProcessId:    1007,
		FileName:     "myscript2.py",
		Status:       "running",
		CreationTime: qtime,
		UpdateTime:   mtime,
		LaunchTime:   stime,
		ServiceId:    5}

	tasks_map[actualTask1.ProcessId] = &actualTask1
	tasks_map[actualTask2.ProcessId] = &actualTask2

	ParseQSTAT(&json_map_multi_tasks, &tasks_map)
	CompareTasks(&expectedTask1, &actualTask1, t)
	CompareTasks(&expectedTask2, &actualTask2, t)
}

func TestParseQSTATOutput(t *testing.T) {
	var time_format = "2006-01-02 15:04:05"
	qtime, _ := time.ParseInLocation(time_format, "2021-04-14 23:34:17", location)
	mtime, _ := time.ParseInLocation(time_format, "2021-04-14 23:34:18", location)
	stime, _ := time.ParseInLocation(time_format, "2021-04-14 23:34:19", location)
	var actualTask Task
	actualTask.ProcessId = 1008
	tasks_map := make(map[int]*Task)
	tasks_map[actualTask.ProcessId] = &actualTask
	expectedTask := Task{
		ProcessId:    1008,
		FileName:     "myscript.py",
		Status:       "running",
		CreationTime: qtime,
		UpdateTime:   mtime,
		LaunchTime:   stime,
		ServiceId:    2}

	ParseQSTAT(&json_map, &tasks_map)
	CompareTasks(&expectedTask, &actualTask, t)
}

func TestRunPBSCommand(t *testing.T) {
	var time_format = "2006-01-02 15:04:05"
	qtime, _ := time.ParseInLocation(time_format, "2021-04-14 23:34:17", location)
	mtime, _ := time.ParseInLocation(time_format, "2021-04-14 23:34:18", location)
	stime, _ := time.ParseInLocation(time_format, "2021-04-14 23:34:19", location)
	actualTask := Task{ProcessId: 1008}
	tasks_map := make(map[int]*Task)
	tasks_map[actualTask.ProcessId] = &actualTask
	expectedTask := Task{
		ProcessId:    1008,
		FileName:     "myscript.py",
		Status:       "running",
		CreationTime: qtime,
		UpdateTime:   mtime,
		LaunchTime:   stime,
		ServiceId:    2}
	args := []string{"tests/imitate_pbs_output.py", "qstat", "-f", "-F", "json", "1008"}
	RunQSTATCommand(&tasks_map, "python3", args)
	CompareTasks(&expectedTask, &actualTask, t)
}

func TestParseQSUB(t *testing.T) {
	var actualTask Task
	expectedTask := Task{ProcessId: 1013}
	ParseQSUB("1013.vm-priboy", &actualTask)
	CompareTasks(&expectedTask, &actualTask, t)
}

func TestRunQSUBCommand(t *testing.T) {
	var actualTask Task
	expectedTask := Task{ProcessId: 1013}
	RunQSUBCommand(&actualTask, "python3", "tests/imitate_pbs_output.py",
		"qsub -v USER_EMAIL=user@mail.ru, PYTHON_SCRIPT=script.py, QUOTA_CPU=1, "+
			"QUOTA_MEMORY=1000, QUOTA_TIME=3600, QUOTA_PRIORITY=1 post_task.py")

	CompareTasks(&expectedTask, &actualTask, t)
}

func TestRunIncorrectQSUBCommand(t *testing.T) {
	var actualTask Task
	// expectedTask := Task{ProcessId: 1013}
	RunQSUBCommand(&actualTask, "python3", "tests/imitate_pbs_output.py",
		"qsub -q linopt post_task.pbs")

	// CompareTasks(&expectedTask, &actualTask, t)
}

func TestParseQSTATSelectedSid(t *testing.T) {
	var time_format = "2006-01-02 15:04:05"
	tasks_map := make(map[int]*Task)

	// prepare the first task
	qtime, _ := time.ParseInLocation(time_format, "2021-04-22 22:56:43", location)
	mtime, _ := time.ParseInLocation(time_format, "2021-04-22 22:56:55", location)
	stime, _ := time.ParseInLocation(time_format, "2021-04-22 22:56:43", location)
	actualTask1 := Task{ProcessId: 2007}
	expectedTask1 := Task{
		ProcessId:    2007,
		FileName:     "myscript1.py",
		Status:       "running",
		CreationTime: qtime,
		UpdateTime:   mtime,
		LaunchTime:   stime,
		ServiceId:    2}

	// prepare the second task
	qtime, _ = time.ParseInLocation(time_format, "2021-04-22 22:56:45", location)
	mtime, _ = time.ParseInLocation(time_format, "2021-04-22 22:56:55", location)
	stime, _ = time.ParseInLocation(time_format, "2021-04-22 22:56:45", location)
	actualTask2 := Task{ProcessId: 2009}
	expectedTask2 := Task{
		ProcessId:    2009,
		FileName:     "myscript2.py",
		Status:       "running",
		CreationTime: qtime,
		UpdateTime:   mtime,
		LaunchTime:   stime,
		ServiceId:    2}

	// prepare the third task
	qtime, _ = time.ParseInLocation(time_format, "2021-04-22 22:56:45", location)
	mtime, _ = time.ParseInLocation(time_format, "2021-04-22 22:56:55", location)
	stime, _ = time.ParseInLocation(time_format, "2021-04-22 22:56:45", location)
	actualTask3 := Task{ProcessId: 2010}
	expectedTask3 := Task{
		ProcessId:    2010,
		FileName:     "myscript3.py",
		Status:       "running",
		CreationTime: qtime,
		UpdateTime:   mtime,
		LaunchTime:   stime,
		ServiceId:    4}

	actualTask4 := Task{ProcessId: 2001}
	actualTask5 := Task{ProcessId: 2002}
	actualTask6 := Task{ProcessId: 2003}
	tasks_map[actualTask1.ProcessId] = &actualTask1
	tasks_map[actualTask2.ProcessId] = &actualTask2
	tasks_map[actualTask3.ProcessId] = &actualTask3
	tasks_map[actualTask4.ProcessId] = &actualTask4
	tasks_map[actualTask5.ProcessId] = &actualTask5
	tasks_map[actualTask6.ProcessId] = &actualTask6

	args := []string{"tests/imitate_pbs_output.py", "qstat", "-f", "-F", "json", "2007", "2009", "2010"}
	RunQSTATCommand(&tasks_map, "python3", args)
	CompareTasks(&expectedTask1, &actualTask1, t)
	CompareTasks(&expectedTask2, &actualTask2, t)
	CompareTasks(&expectedTask3, &actualTask3, t)
	if actualTask4.Status != Completed {
		t.Error("Wrong actual status for actualTask4")
	}
	if actualTask5.Status != Completed {
		t.Error("Wrong actual status for actualTask6")
	}
	if actualTask6.Status != Completed {
		t.Error("Wrong actual status for actualTask6")
	}

}

func TestParseQSTATTable(t *testing.T) {
	qstat_output := `Job id            Name             User              Time Use S Queue
					 ----------------  ---------------- ----------------  -------- - -----
					 2008.g3           STDIN            ivan                     0 Q workq           
					 2009.g3           STDIN            ivan                     0 Q workq           
					 2010.g3           STDIN            ivan                     0 Q workq           
					 2011.g3           STDIN            ivan                     0 Q workq           
					 2012.g3           STDIN            ivan                     0 Q workq           
					 2013.g3           STDIN            ivan                     0 Q workq           
					 2014.g3           STDIN            ivan                     0 Q workq           
					 2015.g3           STDIN            ivan                     0 Q workq
	`
	expected_array := []int{2008, 2009, 2010, 2011, 2012, 2013, 2014, 2015}
	expected_map := make(map[int]int)
	for _, value := range expected_array {
		expected_map[value] = value
	}

	actual_map := ParseQSTATTable(qstat_output)

	for i := range expected_map {
		if _, ok := (*actual_map)[i]; !ok {
			t.Errorf("value = %v  not found in actual_map", i)
		}
	}
}

func TestGetAllQSTATTasks(t *testing.T) {
	expected_array := []int{2008, 2009, 2010, 2011, 2012, 2013, 2014, 2015}
	expected_map := make(map[int]int)
	for _, value := range expected_array {
		expected_map[value] = value
	}
	actual_tasks_map, ok := GetAllQSTATTasks("python3", "tests/imitate_pbs_output.py", "qstat")
	if ok {
		for i := range expected_map {
			if _, ok := (*actual_tasks_map)[i]; !ok {
				t.Errorf("value = %v  not found in actual_map", i)
			}
		}
	} else {
		t.Error("Empty map was returned by GetAllQSTATTasks")
	}
}
