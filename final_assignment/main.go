package main

import (
	"fmt"
	"hacktiv8-course/final_assignment/cmd/server"
	"runtime"
)

func main() {
	appName := "Final Assignment - Hacktiv8 Indra Octama"
	goVersion := runtime.Version()
	osArch := fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH)

	fmt.Println("App Name	:", appName)
	fmt.Println("Go Version	:", goVersion)
	fmt.Println("OS Arch	:", osArch)

	server.StartServer()
}
