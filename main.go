package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	filePath := flag.String("path", "problems.csv", "Path to file containing problems")
	flag.Parse()

	file, err := os.Open(*filePath)

	if err != nil {
		log.Fatalf("Failed to Read csv file. Msg: %s", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	questionset, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	totalQuestions := len(questionset)
	rightQuestions := 0

	inputReader := bufio.NewReader(os.Stdin)

	for _, question := range questionset {
		fmt.Printf("%s = ", question[0])
		answer, err := inputReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		// convert CRLF to LF
		answer = strings.Replace(answer, "\n", "", -1)
		if strings.Compare(answer, question[1]) == 0 {
			rightQuestions++
		}
	}

	fmt.Printf("\nTotal Question %d, Right Questions: %d", totalQuestions, rightQuestions)

}
