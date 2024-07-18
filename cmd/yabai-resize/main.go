package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"

	"github.com/spf13/pflag"
)

func getPos(pos []bool) string {
	//TODO: Deal with with extremes: i.e. 3x3 grid center window, windows with both north/south or east/west, etc
	// Maybe just pick a direction based on dir and pass dir to this function?
	voteDir := make(map[string]int)

	voteDir["top_right"] = 0
	voteDir["bottom_right"] = 0
	voteDir["top_left"] = 0
	voteDir["bottom_left"] = 0

	if pos[0] {
		voteDir["top_left"]++
		voteDir["top_right"]++
	}
	if pos[1] {
		voteDir["bottom_left"]++
		voteDir["bottom_right"]++
	}
	if pos[2] {
		voteDir["top_right"]++
		voteDir["bottom_right"]++
	}
	if pos[3] {
		voteDir["top_left"]++
		voteDir["bottom_left"]++
	}

	best := "top_right"
	bestInt := voteDir["top_right"]

	for v, i := range voteDir {
		if i > bestInt {
			bestInt = i
			best = v
		}
	}

	return best
}

func queryNeighbors() []bool {
	neighbors := []bool{false, false, false, false}
	dirs := []string{"north", "south", "east", "west"}

	for i := 0; i < len(dirs); i++ {
		neighborCmd := exec.Command("yabai", "-m", "query", "--windows", "--window", dirs[i])

		if err := neighborCmd.Start(); err != nil {
			fmt.Println(err)
		}

		if err := neighborCmd.Wait(); err == nil {
			neighbors[i] = true
		}
	}
	return neighbors
}

func amountString(p string, amt int) string {
	return fmt.Sprintf("%s%d", p, amt)
}

func main() {
	var amountX string
	var amountY string

	var resAmt int
	var dir string

	pflag.IntVarP(&resAmt, "amount", "a", 10, "amount to resize in px")
	pflag.StringVarP(&dir, "dir", "d", "", "direction to push")

	pflag.Parse()

	pos := queryNeighbors()
	corner := getPos(pos)

	// learn how to switch case, nerd
	switch dir {
	case "left":
		switch corner {
		case "top_left":
			amountX = amountString("", resAmt)
			amountY = strconv.Itoa(0)
		case "top_right":
			amountY = strconv.Itoa(0)
			amountX = amountString("-", resAmt)
		case "bottom_left":
			amountX = amountString("-", resAmt)
			amountY = strconv.Itoa(0)
		case "bottom_right":
			amountY = strconv.Itoa(0)
			amountX = amountString("-", resAmt)
		}
	case "right":
		switch corner {
		case "top_left":
			amountX = amountString("-", resAmt)
			amountY = strconv.Itoa(0)
		case "top_right":
			amountY = strconv.Itoa(0)
			amountX = amountString("", resAmt)
		case "bottom_left":
			amountX = amountString("", resAmt)
			amountY = strconv.Itoa(0)
		case "bottom_right":
			amountY = strconv.Itoa(0)
			amountX = amountString("", resAmt)
		}

	case "up":
		switch corner {
		case "top_left":
			amountY = amountString("", resAmt)
			amountX = strconv.Itoa(0)
		case "top_right":
			amountX = strconv.Itoa(0)
			amountY = amountString("", resAmt)
		case "bottom_left":
			amountY = amountString("-", resAmt)
			amountX = strconv.Itoa(0)
		case "bottom_right":
			amountX = strconv.Itoa(0)
			amountY = amountString("-", resAmt)
		}
	case "down":
		switch corner {
		case "top_left":
			amountY = amountString("-", resAmt)
			amountX = strconv.Itoa(0)
		case "top_right":
			amountX = strconv.Itoa(0)
			amountY = amountString("-", resAmt)
		case "bottom_left":
			amountY = amountString("", resAmt)
			amountX = strconv.Itoa(0)
		case "bottom_right":
			amountX = strconv.Itoa(0)
			amountY = amountString("", resAmt)
		}
	default:
		log.Println("Empty or invalid dir provided")
		pflag.Usage()
	}

	resizeVar := fmt.Sprintf("%s:%s:%s", corner, amountX, amountY)
	fmt.Println(resizeVar)

	resizeCmd := exec.Command("yabai", "-m", "window", "--resize", resizeVar)

	resizeCmd.Start()
}
