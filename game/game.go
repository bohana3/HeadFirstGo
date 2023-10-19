package game

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const maxRetries = 10

func Play() {

	var target int
	for {
		// Generate a random number between 0 and 99
		target = rand.Intn(101)
		if target != 0 {
			break
		}
	}
	retries := maxRetries
	for {

		fmt.Printf("Type a number: ")
		reader := bufio.NewReader(os.Stdin)

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)

		source, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		var status string
		if source > target {
			status = "Oops its too high"
		} else if source < target {
			status = "Oops its too low"
		} else {
			status = "Good job you guessed it"
		}

		fmt.Println(status)
		if source == target {
			break
		}

		retries--
		if retries == 0 {
			status = fmt.Sprintf("Sorry you did not guess the correct number. it was %d", source)
		} else {
			status = fmt.Sprintf("You have %d attempts left", retries)
		}
		fmt.Println(status)

		if retries == 0 {
			break
		}
	}
}
