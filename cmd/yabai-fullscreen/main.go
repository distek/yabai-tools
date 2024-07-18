package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/distek/yabai-tools/internal"
	"github.com/distek/yabai-tools/structs"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/pflag"
	yaml "gopkg.in/yaml.v3"
)

var (
	fullscreenBool bool
	config         structs.Config
)

func fullscreen(w structs.Window) {
	var space structs.Space

	var wg sync.WaitGroup

	wg.Add(1)
	internal.StartCmd(&wg, "-m window --toggle zoom-fullscreen")

	spacesSlice, _ := internal.RunCmdBytes("-m query --spaces --space")

	err := json.Unmarshal(spacesSlice, &space)
	if err != nil {
		return
	}

	currentPadding, _ := internal.RunCmd("-m config --space " + fmt.Sprint(space.ID) + " top_padding")

	if len(space.Windows) > 0 {
		if w.IsFloating {
			return
		}

		if w.HasFullscreenZoom {
			padInt, _ := strconv.ParseInt(currentPadding[0], 10, 64)
			if padInt > 0 {
				wg.Add(1)
				go internal.StartCmd(&wg, "-m space --toggle padding")
				for _, v := range config.Fullscreen.Restore.Commands {
					wg.Add(1)
					go internal.ExtraCmd(v, &wg)
				}
			}
		} else {
			if currentPadding[0] != "0" {
				wg.Add(1)
				go internal.StartCmd(&wg, "-m space --toggle padding")
				for _, v := range config.Fullscreen.Full.Commands {
					wg.Add(1)
					go internal.ExtraCmd(v, &wg)
				}
			}
		}
	}
	wg.Wait()
}

func main() {
	var window structs.Window
	windowInfo, _ := internal.RunCmdBytes("-m query --windows --window")

	err := json.Unmarshal(windowInfo, &window)
	if err != nil {
		log.Fatal(err)
	}

	fullscreen(window)
}

func init() {
	pflag.BoolVarP(&fullscreenBool, "fullscreen", "f", false, "toggle fullscreen")

	pflag.Parse()

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
