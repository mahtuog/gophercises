package main

import (
	"flag"
	"gophercises/gopherquiz"
)

func main() {

	timePtr := flag.Int("limit", 2, "time limit for each question")
	topicPtr := flag.String("file", "numeric.txt", "quiz file path")

	flag.Parse()

	gopherquiz.Quizzer(*timePtr, *topicPtr)
}
