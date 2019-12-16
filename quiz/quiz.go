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
	lines, err := r.ReadAll()
	problems := parseLines(lines)
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	score := 0

quizloop:
	for i, problem := range problems {
		if err != nil {
			fmt.Println("Failed to parse the provided CSV line.")
			continue
		}

		answerCh := make(chan string)

		go func() {
			var text string
			reader := bufio.NewReader(os.Stdin)
			fmt.Printf("Problem #%v: %v = ", i+1, problem.question)
			text, _ = reader.ReadString('\n')
			answerCh <- text
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break quizloop
		case text := <-answerCh:
			text = strings.TrimSuffix(text, "\n")

			if text == problem.answer {
				score++
			}
		}

	}
	fmt.Printf("You scored %v out of %v\n", score, len(problems))
}

type problem struct {
	question string
	answer   string
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return problems
}
