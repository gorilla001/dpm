package common

import (
	"fmt"
	"os"
)

func ShowAvailableServices() {
	fmt.Println(" - redis")
	fmt.Println(" - rmq")
	fmt.Println(" - mysql")
	fmt.Println(" - influxdb")
	fmt.Println(" - elasticsearch")
	fmt.Println(" - logstash")
}

func Exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
