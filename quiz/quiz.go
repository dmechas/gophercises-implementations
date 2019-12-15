package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
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
	score := 0
	nProblems := 0

quizloop:
	for {
		problem, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Failed to parse the provided CSV line.")
			continue
		}
		nProblems++

		answerCh := make(chan string)

		go func() {
			var text string
			reader := bufio.NewReader(os.Stdin)
			fmt.Printf("Problem #%v: %v = ", nProblems, problem[0])
			text, _ = reader.ReadString('\n')
			answerCh <- text
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break quizloop
		case text := <-answerCh:
			text = strings.TrimSuffix(text, "\n")

			if text == problem[1] {
				score++
			}
		}

	}
	fmt.Printf("You scored %v out of %v\n", score, nProblems)
}
