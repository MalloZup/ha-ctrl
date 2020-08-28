package main

import (
	ps "github.com/shirou/gopsutil/process"
	log "github.com/sirupsen/logrus"
)

func checkPacemakerProcesses(proc *ps.Process) {
	pacemakerProcNames := []string{
		// this might be different across versions..
		"/usr/sbin/pacemakerd",
		"/usr/lib/pacemaker/pacemaker-based",
		"/usr/lib/pacemaker/pacemaker-fenced",
		"/usr/lib/pacemaker/pacemaker-execd",
		"/usr/lib/pacemaker/pacemaker-attrd",
		"/usr/lib/pacemaker/pacemaker-schedulerd",
		"/usr/lib/pacemaker/pacemaker-controld",
	}
	procExec, _ := proc.Exe()
	log.Debugf("%d\t%s\n", proc.Pid, procExec)
	// go trough the list of our process and check their status
	for _, p := range pacemakerProcNames {
		if procExec == p {
			procStatus, _ := proc.Status()
			log.Infof("%d\t%s\t%s\n", proc.Pid, procExec, procStatus)
			// R: Running S: Sleep T: Stop I: Idle
			// Z: Zombie W: Wait L: Lock
			// tollerate running and sleeping. Otherwise print warning
			if procStatus != "R" && procStatus != "S" {
				log.Warnf("Process %s\t state is in a not expected status: %d\t%s\n", procExec, proc.Pid, procStatus)
			}
		}
	}

}

func main() {

	log.Info("Starting ha-control..")

	processList, err := ps.Processes()
	if err != nil {
		log.Println("ps.Processes() Failed, are you using windows?")
		return
	}

	for x := range processList {
		proc := processList[x]
		procExec, _ := proc.Exe()
		log.Debugf("%d\t%s\n", proc.Pid, procExec)
		// go trough the list
		checkPacemakerProcesses(proc)

	}

	log.Info("Health check sucessfully completed!")

}
