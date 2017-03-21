package main

import (
	"log"
	"os/exec"
	"time"
)

func run(file string) error {
	return exec.Command("PowerShell", "-ExecutionPolicy", "Bypass", "-File", file).Run()
}

func main() {

	timer1 := time.NewTimer(time.Hour * 1)

	for {
		<-timer1.C
		err := run("notification.ps1")
		if err != nil {
			log.Fatalln(err)
		}

		timer1 = time.NewTimer(time.Hour * 1)
	}
}
