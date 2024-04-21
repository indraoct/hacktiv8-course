package main

import (
	"fmt"
	"hacktiv8-course/assignment2/cmd/server"
	"runtime"
)

func main() {
	appName := "Assignment 2 - Hactiv8"
	goVersion := runtime.Version()
	osArch := fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH)

	fmt.Println("App Name	:", appName)
	fmt.Println("Go Version	:", goVersion)
	fmt.Println("OS Arch	:", osArch)

	server.StartServer()
}
