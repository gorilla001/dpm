package pull

import (
	"dpm/common"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
)

const (
	SourceCodeBaseDirectory    = "/usr/local/go/src/github.com/Dataman-Cloud/"
	AppSourceCodeDirectory     = "/usr/local/go/src/github.com/Dataman-Cloud/omega-app"
	BillingSourceCodeDirectory = "/usr/local/go/src/github.com/Dataman-Cloud/omega-billing"
	ClusterSourceCodeDirectory = "/usr/local/go/src/github.com/Dataman-Cloud/omega-cluster"
)

var SourceCodeDirectoryMap = map[string]string{
	"app":     "/usr/local/go/src/github.com/Dataman-Cloud/omega-app",
	"billing": "/usr/local/go/src/github.com/Dataman-Cloud/omega-billing",
	"cluster": "/usr/local/go/src/github.com/Dataman-Cloud/omega-cluster",
	"es":      "/usr/local/go/src/github.com/Dataman-Cloud/omega-es",
	"logging": "/usr/local/go/src/github.com/Dataman-Cloud/omega-logging",
}

const (
	AppRepository      = "https://github.com/Dataman-Cloud/omega-app.git"
	BillingRepository  = "https://github.com/Dataman-Cloud/omega-billing.git"
	ClusterRepository  = "https://github.com/Dataman-Cloud/omega-cluster.git"
	EsRepository       = "https://github.com/Dataman-Cloud/omega-es.git"
	LoggingRepository  = "https://github.com/Dataman-Cloud/omega-logging.git"
	MetricsRepository  = "https://github.com/Dataman-Cloud/omega-metrics.git"
	AlertRepository    = "https://github.com/Dataman-Cloud/sryun-alert.git"
	DroneRepository    = "https://github.com/Dataman-Cloud/drone.git"
	FrontendRepository = "https://github.com/Dataman-Cloud/frontend.git"
	WebPageRepository  = "https://github.com/Dataman-Cloud/webpage.git"
	HarborRepository   = "https://github.com/Dataman-Cloud/harbor.git"
)

var (
	PullAppRepositoryCommand     = fmt.Sprintf("%s %s %s", "git clone -b master", AppRepository, AppSourceCodeDirectory)
	PullBillingRepositoryCommand = fmt.Sprintf("%s %s %s", "git clone -b master", BillingRepository, BillingSourceCodeDirectory)
	PullClusterRepositoryCommand = fmt.Sprintf("%s %s %s", "git clone -b master", ClusterRepository, ClusterSourceCodeDirectory)
	PullEsRepositoryCommand      = fmt.Sprintf("%s %s %s", "git clone -b master", EsRepository, SourceCodeDirectoryMap["es"])
	PullLoggingRepositoryCommand = fmt.Sprintf("%s %s %s", "git clone -b master", LoggingRepository, SourceCodeDirectoryMap["logging"])
)

func PullRepository(service string) {

	if common.Exists(path.Join(SourceCodeDirectoryMap[service], ".git")) {
		fmt.Println("Pulling canceled")
		fmt.Println(fmt.Sprintf("%s has already exists", SourceCodeDirectoryMap[service]))
		return
	}

	var cmd *exec.Cmd
	switch service {
	case "app":
		cmd = exec.Command("/bin/bash", "-c", PullAppRepositoryCommand)
	case "billing":
		cmd = exec.Command("/bin/bash", "-c", PullBillingRepositoryCommand)
	case "cluster":
		cmd = exec.Command("/bin/bash", "-c", PullClusterRepositoryCommand)
	case "es":
		cmd = exec.Command("/bin/bash", "-c", PullEsRepositoryCommand)
	case "logging":
		cmd = exec.Command("/bin/bash", "-c", PullLoggingRepositoryCommand)
	default:
		fmt.Println("Unknown service")
		return
	}
	log.Println("Pulling...")
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	if err := cmd.Run(); err != nil {
		log.Println("Pulling...falied")
		return
	}
	log.Println("Pulling...succeed")
	return
}
