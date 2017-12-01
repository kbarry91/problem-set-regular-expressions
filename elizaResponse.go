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
	//Check if input contains "i xxxx it"
	//assign pattern to re
	re = regexp.MustCompile("(?:i(.*)it)")
	if re.MatchString(input) {
		//capture subtring array of input after a match is found
		matched := re.FindStringSubmatch(input)[1]
		//pass matched string array to matchPronouns to reflect pronouns
		matched = matchPronouns(matched)
		response = "input :" + input + " \noutput : Why do you " + matched + " it? "
		return response
	}

	//Check if input contains "i xxxx you"
	//assign pattern to re
	re = regexp.MustCompile("(?:i(.*)you)")
	if re.MatchString(input) {
		//capture subtring array of input after a match is found
		matched := re.FindStringSubmatch(input)[1]
		//pass matched string array to matchPronouns to reflect pronouns
		matched = matchPronouns(matched)
		response = "input :" + input + " \noutput : Why do you  " + matched + " me?"
		return response
	}

	//Check if input contains a work reference
	//assign pattern to re
	re = regexp.MustCompile("(?:work|job|employ|working) (.*)")
	if re.MatchString(input) {
		//capture subtring array of input after a match is found
		matched := re.FindStringSubmatch(input)[1]
		//pass matched string array to matchPronouns to reflect pronouns
		matched = matchPronouns(matched)
		response = "input :" + input + " \noutput :And do you enjoy workin " + matched + "?"
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
	// **********************Automatically generated input *******************************
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
	//word := elizaResponse("")
	word := elizaResponse(userInput[rand.Intn(len("userInput"))])
	fmt.Print("\n")
	fmt.Print(word)
	fmt.Print("\n")
	/*
		// **********************Manually user entered input *******************************
		//uncomment the following section for userInput from console.You must comment out lines 136-151

		for reader := bufio.NewReader(os.Stdin); ; {
			fmt.Print("Please enter a query to ask eliza!")
			userInput, _ := reader.ReadString('\n')
			userInput = strings.Trim(userInput, "\r\n")
			word := elizaResponse(userInput)
			fmt.Print("\n")
			fmt.Print(word)
			fmt.Print("\n")
		}
	*/

}
