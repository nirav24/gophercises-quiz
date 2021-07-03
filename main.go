package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nirav24/gophercises-quiz/quiz"
)

func main() {

	filePath := flag.String("csv", "problems.csv", "A csv file containing problems in 'question,answer' format")
	timeLimit := flag.Int("time", 30, "A time limit for quiz in seconds")

	flag.Parse()

	file, err := os.Open(*filePath)
	if err != nil {
		log.Fatalf("Failed to Read csv file. Msg: %s", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	questions, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	problems := quiz.ParseProblems(questions)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	result := quiz.PlayQuiz(problems, timer.C)

	fmt.Printf("You scored %d out of %d.\n", result.Right, result.Total)
}
