package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func main() {
	limitPtr := flag.Int("limit", 30, "an int")
	csvPtr := flag.String("csv", "problems.csv", "a string")

	flag.Parse()

	file, err := os.Open(*csvPtr)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvPtr))
	}
	timer := time.NewTimer(time.Duration(*limitPtr) * time.Second)

	r := csv.NewReader(file)
	problems, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	score := 0

quizloop:
	for i, p := range problems {
		if err != nil {
			fmt.Println("Failed to parse the provided CSV line.")
			continue
		}

		answerCh := make(chan string)

		go func() {
			var text string
			reader := bufio.NewReader(os.Stdin)
			fmt.Printf("Problem #%v: %v = ", i+1, p[0])
			text, _ = reader.ReadString('\n')
			answerCh <- text
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break quizloop
		case text := <-answerCh:
			text = strings.TrimSuffix(text, "\n")

			if text == p[1] {
				score++
			}
		}

	}
	fmt.Printf("You scored %v out of %v\n", score, len(problems))
}
