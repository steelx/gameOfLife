package main

import (
	"os"
	"os/exec"
	"runtime"
	"time"
)

func main() {

	//create world
	var world = NewWorld(50, 20)
	world.generateMap()

	for {
		clearTerminal()
		world.print()
		time.Sleep(1 * time.Second)
	}
}

func clearTerminal() {
	cmd := exec.Command("clear")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
