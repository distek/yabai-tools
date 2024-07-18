package internal

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"sync"
)

// RunCmd - Returns line delimited slice of yabai command and error
func RunCmd(command string) ([]string, error) {
	var ret []string

	cmd := exec.Command("sh", "-c", fmt.Sprintf("yabai %s", command))

	resp, err := cmd.CombinedOutput()

	ret = strings.Split(string(resp), "\n")

	return ret, err
}

// RunCmdBytes - Returns byte array of command output. Useful for JSON parsing
func RunCmdBytes(command string) ([]byte, error) {
	cmd := exec.Command("sh", "-c", fmt.Sprintf("yabai %s", command))

	resp, err := cmd.CombinedOutput()

	return resp, err
}

// StartCmd - Starts command without checking any output (for async runs)
func StartCmd(wg *sync.WaitGroup, command string) {
	cmd := exec.Command("sh", "-c", fmt.Sprintf("yabai %s", command))

	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}

	wg.Done()
}

// BlindCmd - Starts command without checking any output (for async runs)
func BlindCmd(command string) {
	cmd := exec.Command("sh", "-c", command)

	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}

func ExtraCmd(command string, wg *sync.WaitGroup) {
	cmd := exec.Command("sh", "-c", command)

	err := cmd.Run()
	if err != nil {
		log.Println("yabai", command)
		log.Println(err)
	}

	wg.Done()
}
