package main

import (
	"github.com/Mallozup/ha-ctrl/node"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.Info("Starting ha-control..")

	// 01) Check process
	log.Infoln("Checking relavant HA Cluster process")
	node.CheckPacemakerProcesses()
}
