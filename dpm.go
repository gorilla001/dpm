package main 

import "fmt"
import "flag"
import "dpm/install"
import "dpm/common"

func ShowUsage() {
	fmt.Println("Name:")
	fmt.Println("    dpm - dataman cloud package manager")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("    dpm [ install | build | cleanup | pull ]")
	fmt.Println("")
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
				common.ShowAvailableServices()
				return	
			}
			install.InstallService(flag.Args()[1])
					
		case "build":
			if flag.NArg() == 1 {
				fmt.Println("No service specified! run dpm --help for details")
				return	
			}
			fmt.Println(flag.Args()[1])
	}
	return
}
