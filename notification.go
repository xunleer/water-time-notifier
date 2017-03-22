package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

type notification struct {
	Application string `json:"application,omitempty"`
	Title       string `json:"title,omitempty"`
	Message     string `json:"message,omitempty"`
}

func runCommand(file string) error {
	return exec.Command("PowerShell", "-ExecutionPolicy", "Bypass", "-File", file).Run()
}

func (n notification) PushNotification() error {
	tpl, err := template.ParseFiles("notification.ps1")
	if err != nil {
		log.Fatalln(err)
		return err
	}

	var scriptContent bytes.Buffer
	err = tpl.Execute(&scriptContent, n)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	r := rand.New(rand.NewSource(99))
	fileName := fmt.Sprintf("nofifier-%d", r.Uint32())
	file := filepath.Join(os.TempDir(), fileName+".ps1")
	defer os.Remove(file)

	err = ioutil.WriteFile(file, []byte(scriptContent.String()), 0600)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return runCommand(file)
}
