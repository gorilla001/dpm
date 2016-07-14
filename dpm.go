package main 

import "fmt"
import "flag"
import "dpm/install"

func ShowUsage() {
	fmt.Println("Name:")
	fmt.Println("    dpm - dataman cloud package manager")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("    dpm [ install | build | cleanup | pull ]")
	fmt.Println("")
}

func show_available_services() {
	fmt.Println(" - redis")
	fmt.Println(" - rmq")
	fmt.Println(" - mysql")
	fmt.Println(" - influxdb")
	fmt.Println(" - elasticsearch")
	fmt.Println(" - logstash")
}

func main() {

	flag.Parse()

	if flag.NArg() == 0 || flag.Args()[0] == "--help" {
		ShowUsage()
		return
	}

	switch flag.Args()[0] {

		case "install":
			if flag.NArg() == 1 {
				fmt.Println("No service specified! available services are:")
				show_available_services()
				return	
			}
			switch flag.Args()[1] {

				case "redis":
					install.InstallService("redis")
				case "mysql":
					install.InstallService("mysql")
				case "rmq":
					install.InstallService("rmq")
				case "influxdb":
					install.InstallService("influxdb")
				case "elasticsearch":
					install.InstallService("elasticsearch")
				case "logstash":
					install.InstallService("logstash")
				default:
					fmt.Println("Unknown service! available services are:")
					show_available_services()
					return
			}
					
		case "build":
			if flag.NArg() == 1 {
				fmt.Println("No service specified! run dpm --help for details")
				return	
			}
			fmt.Println(flag.Args()[1])
	}
	return
}
