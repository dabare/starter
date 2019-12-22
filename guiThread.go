package main

import (
	"github.com/mitchellh/go-ps"
	"log"
	"os/exec"
	"time"
)

var root = `starter\inventory.exe`
var url = `http://server:7778`

func checkGUI() {
	pid := startGUI()
	process, _ := ps.FindProcess(pid)
	for true {
		if pid == -1 || process == nil {
			log.Println("GUI was killed")
			pid = startGUI()
		}
		//log.Println(pid)
		process, _ = ps.FindProcess(pid)
		//log.Println(process)

		time.Sleep(1 * time.Second)
	}
}

//`--kiosk-printing` ,`--disable-print-preview`, `--use-system-default-printer`,  `--incognito`,
func startGUI() int {
	//	cmd := exec.Command(root, `--new-window`, `--kiosk`, `--suppress-message-center-popups`, `--user-data-dir=`+root+`Data`, url)
	cmd := exec.Command(root, `--new-window`, `--kiosk`, `--kiosk-printing`, `--suppress-message-center-popups`, `--user-data-dir=Data`, url)
	if err := cmd.Run(); err != nil {
		log.Println("Error:", err)
		return -1
	}
	log.Println("GUI started")
	return cmd.Process.Pid
}
