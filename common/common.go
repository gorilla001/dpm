package common

import (
	"fmt"
)

func ShowAvailableServices() {
	fmt.Println(" - redis")
	fmt.Println(" - rmq")
	fmt.Println(" - mysql")
	fmt.Println(" - influxdb")
	fmt.Println(" - elasticsearch")
	fmt.Println(" - logstash")
}
