package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	randNum := rand.Intn(100) + 1
	success := false

	for guesses := 0; guesses < 10; guesses++ {
		fmt.Print("Provide number and press enter: ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		userNum, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if userNum < randNum {
			fmt.Println("Given number is too small.")
		} else if userNum > randNum {
			fmt.Println("Given number is too big.")
		} else {
			fmt.Println("You won!")
			success = true
			break
		}
	}

	if !success {
		fmt.Println("This time you lost :(. The number being searched for is:", randNum)
	}
}
