package node

import (
	"os/exec"

	log "github.com/sirupsen/logrus"
)

// CheckClusterServices check all systemd units required
func CheckClusterServices() {

	clusterServices := []string{
		"pacemaker",
		"corosync",
	}

	for _, service := range clusterServices {
		systemctlStatus(service)
	}

}

// SystemctlStatus call systemctl status on service
func systemctlStatus(service string) error {
	out, err := exec.Command("/usr/bin/systemctl", "status", service).CombinedOutput()
	if err != nil {
		log.Errorf("service %s is not running correctly. %s \n error %s", service, out, err)
		return err
	}
	log.Infof("service %s is up and running", service)
	return nil
}
