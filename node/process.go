package node

import (
	ps "github.com/shirou/gopsutil/process"
	log "github.com/sirupsen/logrus"
)

// CheckPacemakerProcesses check pacemaker process
func CheckPacemakerProcesses() {

	pacemakerProcNames := map[string]bool{
		// this might be different across versions..
		"/usr/sbin/pacemakerd":                    false,
		"/usr/lib/pacemaker/pacemaker-based":      false,
		"/usr/lib/pacemaker/pacemaker-fenced":     false,
		"/usr/lib/pacemaker/pacemaker-execd":      false,
		"/usr/lib/pacemaker/pacemaker-attrd":      false,
		"/usr/lib/pacemaker/pacemaker-schedulerd": false,
		"/usr/lib/pacemaker/pacemaker-controld":   false,
	}

	processList, err := ps.Processes()
	if err != nil {
		log.Println("ps.Processes() Failed, are you using windows?")
		return
	}

	for x := range processList {
		proc := processList[x]
		procExec, _ := proc.Exe()
		log.Debugf("%d\t%s\n", proc.Pid, procExec)

		// go trough the list of our process and check their status
		for key := range pacemakerProcNames {
			if procExec == key {
				// process expected was found
				pacemakerProcNames[procExec] = true

				procStatus, _ := proc.Status()
				log.Infof("%d\t%s\t%s\n", proc.Pid, procExec, procStatus)
				// R: Running S: Sleep T: Stop I: Idle  Z: Zombie W: Wait L: Lock
				// tollerate running and sleeping. Otherwise print warning
				if procStatus != "R" && procStatus != "S" {
					log.Warnf("Process %s\t state is in a not expected status: %d\t%s\n", procExec, proc.Pid, procStatus)
				}
			}
		}
	}

	// check if we don't have some process not active
	for key, value := range pacemakerProcNames {
		if value == false {
			log.Errorf("Process %s needed for HA was not found", key)
		}
	}
}
