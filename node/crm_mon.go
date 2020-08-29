package node

import (
	"encoding/xml"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

type crmMonResult struct {
	Summary struct {
		Stack struct {
			Type string `xml:"type,attr"`
		} `xml:"stack"`
		CurrentDc struct {
			Present    string `xml:"present,attr"`
			Version    string `xml:"version,attr"`
			Name       string `xml:"name,attr"`
			ID         string `xml:"id,attr"`
			WithQuorum string `xml:"with_quorum,attr"`
		} `xml:"current_dc"`
		LastUpdate struct {
			Time string `xml:"time,attr"`
		} `xml:"last_update"`
		LastChange struct {
			Time   string `xml:"time,attr"`
			User   string `xml:"user,attr"`
			Client string `xml:"client,attr"`
			Origin string `xml:"origin,attr"`
		} `xml:"last_change"`
		NodesConfigured struct {
			Number string `xml:"number,attr"`
		} `xml:"nodes_configured"`
		ResourcesConfigured struct {
			Number   string `xml:"number,attr"`
			Disabled string `xml:"disabled,attr"`
			Blocked  string `xml:"blocked,attr"`
		} `xml:"resources_configured"`
		ClusterOptions struct {
			StonithEnabled   string `xml:"stonith-enabled,attr"`
			SymmetricCluster string `xml:"symmetric-cluster,attr"`
			NoQuorumPolicy   string `xml:"no-quorum-policy,attr"`
			MaintenanceMode  string `xml:"maintenance-mode,attr"`
		} `xml:"cluster_options"`
	} `xml:"summary"`
	Nodes struct {
		Node []struct {
			Name             string `xml:"name,attr"`
			ID               string `xml:"id,attr"`
			Online           string `xml:"online,attr"`
			Standby          string `xml:"standby,attr"`
			StandbyOnfail    string `xml:"standby_onfail,attr"`
			Maintenance      string `xml:"maintenance,attr"`
			Pending          string `xml:"pending,attr"`
			Unclean          string `xml:"unclean,attr"`
			Shutdown         string `xml:"shutdown,attr"`
			ExpectedUp       string `xml:"expected_up,attr"`
			IsDc             string `xml:"is_dc,attr"`
			ResourcesRunning string `xml:"resources_running,attr"`
			Type             string `xml:"type,attr"`
		} `xml:"node"`
	} `xml:"nodes"`
	Resources struct {
		Resource []struct {
			ID             string `xml:"id,attr"`
			ResourceAgent  string `xml:"resource_agent,attr"`
			Role           string `xml:"role,attr"`
			Active         string `xml:"active,attr"`
			Orphaned       string `xml:"orphaned,attr"`
			Blocked        string `xml:"blocked,attr"`
			Managed        string `xml:"managed,attr"`
			Failed         string `xml:"failed,attr"`
			FailureIgnored string `xml:"failure_ignored,attr"`
			NodesRunningOn string `xml:"nodes_running_on,attr"`
			TargetRole     string `xml:"target_role,attr"`
			Node           struct {
				Name   string `xml:"name,attr"`
				ID     string `xml:"id,attr"`
				Cached string `xml:"cached,attr"`
			} `xml:"node"`
		} `xml:"resource"`
		Clone []struct {
			ID             string `xml:"id,attr"`
			MultiState     string `xml:"multi_state,attr"`
			Unique         string `xml:"unique,attr"`
			Managed        string `xml:"managed,attr"`
			Failed         string `xml:"failed,attr"`
			FailureIgnored string `xml:"failure_ignored,attr"`
			Resource       []struct {
				ID             string `xml:"id,attr"`
				ResourceAgent  string `xml:"resource_agent,attr"`
				Role           string `xml:"role,attr"`
				Active         string `xml:"active,attr"`
				Orphaned       string `xml:"orphaned,attr"`
				Blocked        string `xml:"blocked,attr"`
				Managed        string `xml:"managed,attr"`
				Failed         string `xml:"failed,attr"`
				FailureIgnored string `xml:"failure_ignored,attr"`
				NodesRunningOn string `xml:"nodes_running_on,attr"`
				Pending        string `xml:"pending,attr"`
				Node           struct {
					Name   string `xml:"name,attr"`
					ID     string `xml:"id,attr"`
					Cached string `xml:"cached,attr"`
				} `xml:"node"`
			} `xml:"resource"`
		} `xml:"clone"`
	} `xml:"resources"`
	NodeAttributes struct {
		Node []struct {
			Name      string `xml:"name,attr"`
			Attribute []struct {
				Name  string `xml:"name,attr"`
				Value string `xml:"value,attr"`
			} `xml:"attribute"`
		} `xml:"node"`
	} `xml:"node_attributes"`
	NodeHistory struct {
		Node []struct {
			Name            string `xml:"name,attr"`
			ResourceHistory []struct {
				ID                 string `xml:"id,attr"`
				Orphan             string `xml:"orphan,attr"`
				MigrationThreshold string `xml:"migration-threshold,attr"`
				FailCount          string `xml:"fail-count,attr"`
				LastFailure        string `xml:"last-failure,attr"`
				OperationHistory   []struct {
					Call         string `xml:"call,attr"`
					Task         string `xml:"task,attr"`
					LastRcChange string `xml:"last-rc-change,attr"`
					LastRun      string `xml:"last-run,attr"`
					ExecTime     string `xml:"exec-time,attr"`
					QueueTime    string `xml:"queue-time,attr"`
					Rc           string `xml:"rc,attr"`
					RcText       string `xml:"rc_text,attr"`
					Interval     string `xml:"interval,attr"`
				} `xml:"operation_history"`
			} `xml:"resource_history"`
		} `xml:"node"`
	} `xml:"node_history"`
	Failures struct {
		Failure []struct {
			OpKey        string `xml:"op_key,attr"`
			Node         string `xml:"node,attr"`
			Exitstatus   string `xml:"exitstatus,attr"`
			Exitreason   string `xml:"exitreason,attr"`
			Exitcode     string `xml:"exitcode,attr"`
			Call         string `xml:"call,attr"`
			Status       string `xml:"status,attr"`
			LastRcChange string `xml:"last-rc-change,attr"`
			Queued       string `xml:"queued,attr"`
			Exec         string `xml:"exec,attr"`
			Interval     string `xml:"interval,attr"`
			Task         string `xml:"task,attr"`
		} `xml:"failure"`
	} `xml:"failures"`
	Status struct {
		Code    string `xml:"code,attr"`
		Message string `xml:"message,attr"`
	} `xml:"status"`
}

// CheckCrmMonStatus with crm_mon interface, usefull for resource check etc
func CheckCrmMonStatus() {
	var crmMon *crmMonResult
	crmMonXML, err := exec.Command("/usr/sbin/crm_mon", "-X", "--inactive").Output()
	if err != nil {
		log.Errorf("error while executing crm_mon")
	}

	err = xml.Unmarshal(crmMonXML, &crmMon)
	if err != nil {
		log.Errorf("error while parsing crm_mon XML output")
	}
	log.Infof("%s", crmMon)
}
