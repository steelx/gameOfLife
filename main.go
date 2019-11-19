package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func main() {

	//create world
	var world = NewWorld(10, 10)
	world.generateMap()

	for {
		cmd := exec.Command("clear")
		if runtime.GOOS == "windows" {
			cmd = exec.Command("cmd", "/c", "cls")
		}
		cmd.Stdout = os.Stdout
		cmd.Run()

		//print
		for y := 0; y < world.height; y++ {
			//creates new row
			fmt.Print("|")

			//columns
			for x := 0; x < world.width; x++ {
				fmt.Print(getChar(world.getCell(x, y)))
			}

			fmt.Println("|")
		}

		time.Sleep(1 * time.Second)
	}
}

func getChar(cell Cell) string {
	if cell.Alive {
		return "*"
	}
	return " "
}
