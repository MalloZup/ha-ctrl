package main

import (
	"github.com/MalloZup/ha-ctrl/node"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.Info("Starting ha-control..")

	log.Infoln("Checking relavant HA Cluster process")
	node.CheckPacemakerProcesses()

	log.Infoln("Checking services status for HA cluster")
	node.CheckClusterServices()

	log.Infoln("Checking cluster status with help of crm_mon")
	node.CheckClusterNodes()

	log.Infoln("Checking if stonith is enabled")
	node.DoStonithChecks()
}
