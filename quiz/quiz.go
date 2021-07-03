// Quiz package provides functions to run the quiz
package quiz

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Problem struct {
	Question string
	Answer   string
}

// ParseProblems converts 2D string array into slice of Problem
func ParseProblems(problems [][]string) []Problem {
	out := make([]Problem, len(problems))
	for i, problem := range problems {
		out[i] = Problem{
			Question: problem[0],
			Answer:   strings.TrimSpace(problem[1]),
		}
	}
	return out
}

// Result type
type Result struct {
	Total int
	Right int
}

// PlayQuiz starts the quiz and keep tracks of right questions.
// When quiz ends, it returns Result struct
func PlayQuiz(problems []Problem, ch <-chan time.Time) Result {
	rightQuestions := 0
	inputReader := bufio.NewReader(os.Stdin)

	for i, problem := range problems {
		fmt.Printf("Question %d: %s = ", i+1, problem.Question)
		answerChan := make(chan string)

		go func() {
			answer, err := inputReader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			// convert CRLF to LF
			answer = strings.Replace(answer, "\n", "", -1)
			answerChan <- answer
		}()
		select {
		case <-ch:
			return Result{
				Total: len(problems),
				Right: rightQuestions,
			}
		case answer := <-answerChan:
			if strings.Compare(answer, problem.Answer) == 0 {
				rightQuestions++
			}
		}
	}
	return Result{
		Total: len(problems),
		Right: rightQuestions,
	}
}
