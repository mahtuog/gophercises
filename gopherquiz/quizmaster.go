package gopherquiz

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func checker(userAnswer, realAnswer string) int {
	if realAnswer == userAnswer {
		return 1
	}
	return 0
}

func readInput(c chan string) {
	var userAnswer string
	if _, err := fmt.Scanf("%s", &userAnswer); err != nil {
		fmt.Println("Error reading input due to ", err)
	}
	c <- userAnswer
}

var proceed int

// Quizzer main
func Quizzer(timeLimit int, quizFile string) {
	limit := time.Duration(timeLimit) * time.Second
	file, err := ioutil.ReadFile(quizFile)
	if err != nil {
		fmt.Printf("error opening file %v :", err.Error())
	}

	c := make(chan string)
	contents := strings.Split(string(file), "\n")
	score := 0
	fmt.Println("Number of questions : ", len(contents))

	for i := 0; i < len(contents); i++ {
		proceed = 0

		question := strings.Split(contents[i], "=")[0]
		realAnswer := strings.TrimSpace(strings.Split(contents[i], "=")[1])

		fmt.Printf("%v : ", question)

		go readInput(c)
		select {
		case userAnswer := <-c:
			score = score + checker(realAnswer, userAnswer)
		case <-time.After(limit):
			fmt.Println("\nOut of time after ... " + limit.String() + " seconds")
			proceed = 1
		}
		if proceed != 0 {
			break
		}
	}
	close(c)
	fmt.Printf("Your Score : %v/%v \n", score, len(contents))
}
