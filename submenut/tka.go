package submenut

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
	"github.com/tonivaa/hellafly/hinnat"
)

func Toka() {
	fmt.Println("toka submenu")

	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		keyboard.Close()
	}()

	for {
		fmt.Println("Kerro jotai")

		char, key, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if key == keyboard.KeyBackspace2 {
			fmt.Println("Palataan..")
			break
		}

		i, _ := strconv.Atoi(string(char))

		_, ok := hinnat.Maalaus[i]
		if ok {
			fmt.Println("Tähän tulee", i, hinnat.Maalaus[i])
		} else {
			fmt.Println("Ei löydy:", i)
		}

	}
}
