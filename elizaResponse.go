//  Author: Kevin Barry
//  Date:   15/11/2017
//  Description :problem set 3 regular expressions
//  Source : https://data-representation.github.io/problems/go-regexp.html

package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

func elizaResponse(inputStr string) string {
	/*
		*	Problem 1
		*
		//return random string output
		input := inputStr
		rand.Seed(time.Now().UTC().UnixNano())
		answers := []string{
			"I’m not sure what you’re trying to say. Could you explain it to me?",
			"How does that make you feel?",
			"Why do you say that?",
		}

		response := "input :" + input + " \noutput :" + answers[rand.Intn(len(answers))]
	*/
	input := inputStr
	pattern := `.*father.*`
	pattern2 := []string{`.*i am(.*)`, `.*I AM(.*)`, `.*I'm(.*)`, `.*i'm(.*)`, `.*im(.*)`, `.*I am(.*)`}
	output := "why dont you tell me more about your father?"
	response := ""
	rand.Seed(time.Now().UTC().UnixNano())

	answers := []string{
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
	}

	//check if pattern is fouund in string (father/I am)
	if matched, _ := regexp.MatchString(pattern, input); matched {
		response = "input :" + input + "\noutput :" + output
		//} //else if re.MatchString(input) {
		//match := re.ReplaceAllString(input, "How do you know you are $1?")
		//response = "input :" + input + " \noutput : " + match
	} else {
		response = "input :" + input + " \noutput :" + answers[rand.Intn(len(answers))]
	}

	// loop through pattern2 array
	//if pattern is found extract substring
	//set response
	for _, checkPattern := range pattern2 {
		re := regexp.MustCompile(checkPattern)
		if re.MatchString(input) {
			match := re.ReplaceAllString(input, "How do you know you are $1?")
			response = "input :" + input + " \noutput : " + match
			return response

		} //if re.MatchString
	} //for pattern2

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
		"I am happy.",
		"I am not happy with your responses.",
		"I am not sure that you understand the effect that your questions are having on me.",
		"I am supposed to just take what you’re saying at face value?",
	}

	rand.Seed(time.Now().UTC().UnixNano())

	word := elizaResponse(userInput[rand.Intn(len("userInput"))])
	fmt.Print(word)
}
