package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	message1 := "Feliz Navidad"
	message2 := "And Happy New Year"
	message3 := "Familia Cetina Aldana"

	for {
		printCrismassTree(12, message1, message2, message3)
		time.Sleep(500 * time.Millisecond)
	}
}

func printCrismassTree(altura int, message1, message2, message3 string) {
	clearScreen()
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < altura; i++ {
		for j := 0; j < altura-i; j++ {
			fmt.Print(" ")
		}

		for j := 0; j < (2*i)+1; j++ {
			fmt.Print(getRandomColor() + "*" + resetColor())
		}
		fmt.Println()
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < altura-1; j++ {
			fmt.Print(" ")
		}
		fmt.Println("|||")
	}

	printCenteredMessage(message1, altura)
	printCenteredMessage(message2, altura)
	printCenteredMessage(message3, altura)
}

func resetColor() string {
	return "\033[0m"
}

func getRandomColor() string {
	colorCode := rand.Intn(6) + 31
	return fmt.Sprintf("\033[%dm", colorCode)
}

func printCenteredMessage(message string, width int) {
	padding := width - len(message)/2

	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}
	fmt.Println(message)
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
