package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
	"github.com/tonivaa/hellafly/submenut"
)

func menu() {

	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	submenut := map[int]func(){
		1: submenut.Eka,
		2: submenut.Toka,
		3: submenut.Kolmas,
	}

	for {
		fmt.Println("Valitse alakategoria 1-3 atm. q poistuu jne")

		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}
		if char == 'q' || char == 'Q' {
			break
		}

		i, _ := strconv.Atoi(string(char))

		_, ok := submenut[i]
		if ok {
			submenut[i]()
		}
	}

	fmt.Println("Suljetaan,,,,")

}

func main() {

	menu()

	clearScreen()

	a, b := 5, 2
	calculator(a, b, sub)
	calculator(a, b, add)
}

func clearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		// windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// linux or mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

type operation func(int, int) int

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func calculator(a, b int, op operation) int {
	return op(a, b)
}
