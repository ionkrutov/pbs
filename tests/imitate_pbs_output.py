import time
import sys

s1 = """
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
                "PBS_O_LANG":"en_US.UTF-8",
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
"""

s2 = """
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
                    "PBS_O_LANG":"en_US.UTF-8",
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
                    "PBS_O_LANG":"en_US.UTF-8",
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
    }
"""
s_selected = """
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
	}
"""
qstat_table = """Job id            Name             User              Time Use S Queue
----------------  ---------------- ----------------  -------- - -----
2008.g3           STDIN            ivan                     0 Q workq           
2009.g3           STDIN            ivan                     0 Q workq           
2010.g3           STDIN            ivan                     0 Q workq           
2011.g3           STDIN            ivan                     0 Q workq           
2012.g3           STDIN            ivan                     0 Q workq           
2013.g3           STDIN            ivan                     0 Q workq           
2014.g3           STDIN            ivan                     0 Q workq           
2015.g3           STDIN            ivan                     0 Q workq
"""

qstat_unknow ="""qstat: Unknown Job Id 2001.g3
qstat: Unknown Job Id 2002.g3
qstat: Unknown Job Id 2003.g3
"""


cmd_line = ' '.join(sys.argv[1:])

if cmd_line == "qstat -f -F json 1008":
    print(s1)
    sys.exit(0)
elif cmd_line == "qstat -f -F json":
    print(s2)
    sys.exit(0)
elif cmd_line == "qsub -v USER_EMAIL=user@mail.ru, PYTHON_SCRIPT=script.py, QUOTA_CPU=1, QUOTA_MEMORY=1000, QUOTA_TIME=3600, QUOTA_PRIORITY=1 post_task.py":
    print('1013.vm-priboy')
    sys.exit(0)
elif cmd_line == "qstat -f -F json 2007 2009 2010":
    print(s_selected)
    sys.stderr.write(qstat_unknow)
    sys.exit(0)
elif cmd_line == "qstat":
    print(qstat_table)
    sys.exit(0)
else:
    print("Unknown command tra-ta-ta")
    sys.exit(13)
