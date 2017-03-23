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
	
	"syscall"
	"unsafe"
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

const MB_ICONINFORMATION = 0x00000040

func (n notification) PopMessageBox() error {
	var mod = syscall.NewLazyDLL("user32.dll")
	var proc = mod.NewProc("MessageBoxW")

	proc.Call(0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(n.Message))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(n.Title))),
		uintptr(MB_ICONINFORMATION))

	return nil
}

func (n notification) Notice() error {
	var err error
	if winVer := os.Getenv("WINVER"); winVer == "win10" {
		err = n.PushNotification()
	} else {
		err = n.PopMessageBox()
	}

	return err
}
