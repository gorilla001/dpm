package install

import (
	"dpm/common"
	"fmt"
	"log"
	"os/exec"
)

const (
	PullRedisImageCmd         = "docker pull demoregistry.dataman-inc.com/srypoc/redis:3.0.5"
	PullMysqlImageCmd         = "docker pull demoregistry.dataman-inc.com/srypoc/mysql:5.6"
	PullRmqImageCmd           = "docker pull demoregistry.dataman-inc.com/srypoc/rabbitmq:3.6.0-management"
	PullInfluxdbImageCmd      = "docker pull demoregistry.dataman-inc.com/srypoc/influxdb:0.10"
	PullElasticsearchImageCmd = "docker pull demoregistry.dataman-inc.com/srypoc/centos7-jdk7-elasticsearch-1.4.5-alone:20160522230210"
	PullLogstashImageCmd      = "docker pull demoregistry.dataman-inc.com/srypoc/logstash:1.5.6"
)

func InstallService(service string) {
	var cmd *exec.Cmd
	switch service {

	case "redis":
		cmd = exec.Command("/bin/bash", "-c", PullRedisImageCmd)
	case "mysql":
		cmd = exec.Command("/bin/bash", "-c", PullMysqlImageCmd)
	case "rmq":
		cmd = exec.Command("/bin/bash", "-c", PullRmqImageCmd)
	case "influxdb":
		cmd = exec.Command("/bin/bash", "-c", PullInfluxdbImageCmd)
	case "elasticsearch":
		cmd = exec.Command("/bin/bash", "-c", PullElasticsearchImageCmd)
	case "logstash":
		cmd = exec.Command("/bin/bash", "-c", PullLogstashImageCmd)
	default:
		fmt.Println("Unknown service! available services are:")
		common.ShowAvailableServices()
		return
	}
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalln(fmt.Sprintf("%s:%s", err, output))
		return
	}
	log.Println(string(output))
	return
}
