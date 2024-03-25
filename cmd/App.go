package main

import "service_manager/core"

func main() {
	s := core.NewDeploy("D:/Projects/GolandProjects/service_manager/deploy.toml")
	s.Run()
}
