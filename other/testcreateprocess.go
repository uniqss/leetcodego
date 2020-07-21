package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func CProcess() {
	//args := []string{"/Game/Maps/Highrise", "-log"}

	var args []string
	args = append(args, "/Game/Maps/Highrise")
	args = append(args, "-log")

	var attr = &os.ProcAttr{}
	attr.Dir = "E:\\work\\ShooterGame\\PackageTest\\WindowsNoEditor\\ShooterGame\\Binaries\\Win64"

	attr.Files = []*os.File{os.Stdin,
		os.Stdout, os.Stderr}

	_, err := os.StartProcess("ShooterServer-Win64-DebugGame.exe", args, attr)
	if err != nil {
		log.Error("proc create fail ", err)
		return
	}
}

func main() {
	CProcess()
}
