//  Author: Kevin Barry
//  Date:   15/11/2017
//  Description :problem set 3 regular expressions
//  Source : https://data-representation.github.io/problems/go-regexp.html

package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
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

	} else {
		response = "input :" + input + " \noutput :" + answers[rand.Intn(len(answers))]
	}

	//Check if input contains I am or prounouns
	//assign pattern to re
	re := regexp.MustCompile("(?i)(i[' a]*m) (.*)")
	if re.MatchString(input) {
		//capture subtring array of input after a match is found
		matched := re.FindStringSubmatch(input)[2]
		//pass matched string array to matchPronouns to reflect pronouns
		matched = matchPronouns(matched)
		response = "input :" + input + " \noutput : How do you know you are " + matched
		return response
	}
	return response
} //elizaResponse

func matchPronouns(inputStr string) string {
	// split inputStr into slice of strings
	splitStr := strings.Fields(inputStr)

	//map of reflected pronouns
	pronouns := map[string]string{
		"i":      "you",
		"was":    "were",
		"i'd":    "you would",
		"i've":   "you have",
		"i'll":   "you will",
		"my":     "your",
		"are":    "am",
		"you've": "I have",
		"you'll": "I will",
		"your":   "my",
		"yours":  "mine",
		"you":    "I",
		"me":     "you",
		"me.":    "you",
		"you're": "I’m",
	}

	//loop map word to check if splitStr is a pronoun
	// Found how to iterate over a map from https://gobyexample.com/range
	for index, word := range splitStr {
		if value, ok := pronouns[strings.ToLower(word)]; ok {
			splitStr[index] = value
		}
	}

	return strings.Join(splitStr, " ")
}
func main() {
	//slice of inputs

	userInput := []string{
		"People say I look like both my mother and father.",
		"Father was a teacher.",
		"I was my father’s favourite.",
		"I'm looking forward to the weekend.",
		"My grandfather was French!",
		"I am not happy with your responses.",
		"I am not sure that you understand the effect that your questions are having on me.",
		"I am supposed to just take what you’re saying at face value?",
	}

	rand.Seed(time.Now().UTC().UnixNano())

	word := elizaResponse(userInput[rand.Intn(len("userInput"))])
	fmt.Print("\n")
	fmt.Print(word)
	fmt.Print("\n")
}
