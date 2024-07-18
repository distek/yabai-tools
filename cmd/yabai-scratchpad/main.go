package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/distek/yabai-tools/internal"
	"github.com/distek/yabai-tools/structs"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/pflag"
	"gopkg.in/yaml.v3"
)

var (
	name   string
	config structs.Config
)

func getScratch(name string) (structs.Scratchpad, error) {
	for _, v := range config.Scratchpads {
		if v.Name == name {
			return v, nil
		}
	}

	return structs.Scratchpad{}, fmt.Errorf("no scratch found by that name")
}

func find(windows []structs.Window, scratch structs.Scratchpad) (structs.Window, error) {
	for _, v := range windows {
		if v.App == scratch.App {
			// Match title as well if available
			if scratch.Title != "" {
				titleRx := regexp.MustCompile(scratch.Title)
				if titleRx.MatchString(v.Title) {
					return v, nil
				}
				continue
			}

			return v, nil
		}
	}

	return structs.Window{}, fmt.Errorf("window not found")
}

func main() {
	var windows []structs.Window

	windowOut, err := internal.RunCmdBytes("-m query --windows")
	if err != nil {
		log.Fatal("command failed to run:", err)
	}

	err = json.Unmarshal(windowOut, &windows)
	if err != nil {
		log.Fatal(err)
	}

	spaceOut, err := internal.RunCmdBytes("-m query --spaces --space")
	if err != nil {
		log.Fatal(err)
	}

	var currentSpace structs.Space

	err = json.Unmarshal(spaceOut, &currentSpace)
	if err != nil {
		log.Fatal(err)
	}

	scratch, err := getScratch(name)
	if err != nil {
		log.Fatal("Could not find scratch by that name")
	}

	window, err := find(windows, scratch)

	if err != nil {
		internal.BlindCmd(scratch.OpenCommand)

		return
	}

	switch {
	case window.Space != currentSpace.Index:
		internal.RunCmd(fmt.Sprintf("-m window %d --space %d", window.ID, currentSpace.Index))

		internal.RunCmd(fmt.Sprintf("-m window --focus %d", window.ID))

		// Re-adjusts the border position
		internal.RunCmd(fmt.Sprintf("-m window --move abs:%d:%d", int(window.Frame.X)+1, int(window.Frame.Y)+1))
	case window.Space == currentSpace.Index && !window.HasFocus:
		internal.RunCmd(fmt.Sprintf("-m window --focus %d", window.ID))
	default:
		internal.RunCmd(fmt.Sprintf("-m window %d --space %d", window.ID, config.ScratchpadWorkspace))
		internal.RunCmd(fmt.Sprintf("-m window --focus recent"))
	}
}

func init() {
	pflag.StringVarP(&name, "name", "n", "", "interact with scratchpad <name>")

	pflag.Parse()

	if name == "" {
		log.Fatal("Please provide a scratch name")
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
