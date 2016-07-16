package pull 

import (
	"fmt"
	"os/exec"
)

const (
	SourceCodeDirectory = "/var/lib/dpm/src"
)

const (
	AppRepository = "https://github.com/Dataman-Cloud/omega-app.git"
	BillingRepository = "https://github.com/Dataman-Cloud/omega-billing.git"
	ClusterRepository = "https://github.com/Dataman-Cloud/omega-cluster.git"
	EsRepository = "https://github.com/Dataman-Cloud/omega-es.git"
	LoggingRepository = "https://github.com/Dataman-Cloud/omega-logging.git"
	MetricsRepository = "https://github.com/Dataman-Cloud/omega-metrics.git"
	AlertRepository = "https://github.com/Dataman-Cloud/sryun-alert.git"
	DroneRepository = "https://github.com/Dataman-Cloud/drone.git"
	FrontendRepository = "https://github.com/Dataman-Cloud/frontend.git"
	WebPageRepository = "https://github.com/Dataman-Cloud/webpage.git"
	HarborRepository = "https://github.com/Dataman-Cloud/harbor.git"
)

const (
	PullOmegaAppRepositoryCommand = "git clone -b master %s %s    
)

func PullRepository(service string) {
	var cmd *exec.Cmd
	switch service {
	case "app":
		cmd = exec.Command("/bin/bash", "-c", fmt.Sprintf(PullOmegaAppRepositoryCommand,OmegaAppRepository) 	
	default:
		fmt.Println("Unknown service")
		return
	}
	return
}
