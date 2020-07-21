package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {

	var args []string
	args = append(args, "E:\\work\\ShooterGame\\PackageTest\\WindowsNoEditor\\ShooterGame\\Binaries\\Win64\\ShooterServer-Win64-DebugGame.exe")

	args = append(args, "/Game/Maps/Highrise")
	args = append(args, "-log")

	var attr = &os.ProcAttr{}
	attr.Dir = "E:\\work\\ShooterGame\\PackageTest\\WindowsNoEditor\\ShooterGame\\Binaries\\Win64"
	//attr.Dir = "F:\\leetcodecpp\\Debug"

	attr.Files = []*os.File{os.Stdin,
		os.Stdout, os.Stderr}

	_, err := os.StartProcess("ShooterServer-Win64-DebugGame.exe", args, attr)
	//_, err := os.StartProcess("echoparams.exe", args, attr)
	if err != nil {
		log.Error("proc create fail ", err)
		return
	}
}
