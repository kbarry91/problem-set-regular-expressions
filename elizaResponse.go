//  Author: Kevin Barry
//  Date:   15/11/2017
//  Description :problem set 3 regular expressions
//  Source : https://data-representation.github.io/problems/go-regexp.html

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func elizaResponse(inputStr string) string {
	//return random string output
	input := inputStr
	rand.Seed(time.Now().UTC().UnixNano())
	answers := []string{
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
	}

	response := "input :" + input + " \noutput :" + answers[rand.Intn(len(answers))]
	return response
} //elizaResponse

func main() {
	//slice of inputs
	userInput := []string{
		"People say I look like both my mother and father.",
		"Father was a teacher.",
		"I was my father’s favourite.",
		"I'm looking forward to the weekend.",
		"My grandfather was French!",
	}
	rand.Seed(time.Now().UTC().UnixNano())

	word := elizaResponse(userInput[rand.Intn(len(userInput))])

	fmt.Print(word)
}
