package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Print("   Welcome to GeussTheGo\n\n")
	randomNumber := rand.Intn(100)
	var userNumber int
	var answer string
	fmt.Println("[Computer]: [1, 100] oralig'ida son o'yldim. Topishga harakat qiling.")
	fmt.Print("[ You say]: ")
	fmt.Scan(&userNumber)
	for {
		if userNumber > randomNumber {
			fmt.Println("[Computer]: Siz kiritgan son men o'ylagan sondan katta. Qayta urinib ko'ring.")
			fmt.Print("[ You say]: ")
			fmt.Scan(&userNumber)
			continue
		} else if userNumber < randomNumber {
			fmt.Println("[Computer]: Siz kiritgan son men o'ylagan sondan kichik. Qayta urinib ko'ring.")
			fmt.Print("[ You say]: ")
			fmt.Scan(&userNumber)
			continue
		} else {
			fmt.Println("[Computer]: Ajoyib!!! Siz men o'ylagan sonni topdingiz.")
			fmt.Println("[Computer]: Yana o'ynashini xohlaysizmi? [y/n]")
			fmt.Print("[ You say]: ")
			fmt.Scan(&answer)
			if answer == "y" {
				randomNumber = rand.Intn(100)
				fmt.Println("[Computer]: [1, 100] oralig'ida son o'yldim. Topishga harakat qiling.")
				fmt.Print("[ You say]: ")
				continue
			} else {
				fmt.Println("Bye!!")
				break
			}
		}
	}
}
