package main

import (
	"os"
	"os/exec"
	"runtime"
	"time"
)

func main() {

	//create world
	var world = NewWorld(40, 10)
	world.GenerateGrid(30)

	for {
		clearTerminal()
		world.Print()
		world.Next()
		time.Sleep(500 * time.Millisecond)
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
