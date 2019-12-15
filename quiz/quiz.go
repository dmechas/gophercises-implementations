package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	limitPtr := flag.Int("limit", 30, "an int")
	csvPtr := flag.String("csv", "problems.csv", "a string")

	flag.Parse()

	file, err := os.Open(*csvPtr)
	check(err)

	r := csv.NewReader(file)
	score := 0
	nProblems := 0
	for {
		problem, err := r.Read()
		if err == io.EOF {
			break
		}
		nProblems++

		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Problem #%v: %v =", nProblems, problem[0])
		text, _ := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")

		if text == problem[1] {
			score++
		}
	}
	fmt.Printf("You scored %v out of %v\n", score, nProblems)
}
