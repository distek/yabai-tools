package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/distek/yabai-tools/internal"
	"github.com/distek/yabai-tools/structs"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/pflag"
	"gopkg.in/yaml.v3"
)

var (
	spaceID int
	config  structs.Config
)

func main() {
	var windows []structs.Window
	var space structs.Space
	var wg sync.WaitGroup

	// don't start doing shit until we're actually on the correct space
	count := 0
	for {
	tryAgain:
		spaceCheck, _ := internal.RunCmdBytes("-m query --spaces --space")

		err := json.Unmarshal(spaceCheck, &space)
		if err != nil {
			count++
			if count > 3 {
				log.Fatal("max retries exceeded")
			}
			goto tryAgain
		}

		if space.Index == spaceID {
			break
		}
	}

	windowOut, err := internal.RunCmdBytes("-m query --windows --space " + fmt.Sprint(spaceID))
	if err != nil {
		log.Fatal("command fialed to run")
	}

	err = json.Unmarshal(windowOut, &windows)
	if err != nil {
		return
	}

	for _, v := range windows {
		if v.HasFullscreenZoom {
			for _, v := range config.Switch.IfFullscreen.Commands {
				wg.Add(1)
				internal.ExtraCmd(v, &wg)
			}
			return
		}
	}

	for _, v := range config.Switch.Else.Commands {
		wg.Add(1)
		internal.ExtraCmd(v, &wg)
	}
}

func init() {
	pflag.IntVarP(&spaceID, "space", "s", 0, "switch to space <index>")

	pflag.Parse()

	if spaceID == 0 {
		log.Fatal("Provide a space index with --space <index>")
	}

	// load in config
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}

	confFile, err := os.ReadFile(home + "/.config/yabai/tools/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(confFile, &config)
	if err != nil {
		log.Fatal(err)
	}
}
