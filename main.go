package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var originalQuestions = `- Can you think of a really boring hobby?
	- Why do you think it's important to have a hobby?
	- Do you think having lots of hobbies is bad?
	- What hobbies do your family members have?
	- Do you prefer indoor or outdoor hobbies and why?
	- How often do you leave Dongguan?
	- What do you usually enjoy eating when you’re sad?
	- How often do you watch movies?
	-  What kind of food have you never eaten before?
	- What do you think is more difficult, learning a new language or a musical instrument, and why do you think so?
	- Which season (time of the year) do you think is the best? Why? 
	- Is your hometown colder than Dongguan? How are they different?
	- What is the most beautiful city you have ever visited?
	- What is your dream job? Why?
	- What do you think is the most boring job? Why?
	- When you get a job, how much money would you like to earn?
	- What kind of job do you think is better: part-time or full time? Why?
	- What dishes can you cook? What is your best one?
	- Who do you think cooks the most delicious food?
	- What is the best way to cook eggs? Describe the steps
	- How often do you cook? Do you mostly cook breakfast, lunch or dinner?`

var allQuestions = strings.Split(originalQuestions, "\n")

// var questionSets [25][3]string

var questionSets [][]string
var papersNum int
var questionsNum int

func main() {
	getUserInput()

	populateQuestionPapers()

	writeToFile()

}

func populateQuestionPapers() {
	// Populate question sets with random questions
	// Outer index = papers, inner index = questions in each consequtive paper
	for outerIndex := 0; outerIndex < papersNum; outerIndex++ {
		questionSets = append(questionSets, []string{})
		for innerIndex := 0; innerIndex < 3; innerIndex++ {
			randQuestion := allQuestions[rand.Intn(len(allQuestions))]
			// Check for doubles inside the paper before appending
			for _, q := range questionSets[outerIndex] {

				if q == randQuestion {
					// log.Fatalln("repeat found")
					fmt.Println(q, randQuestion)
					fmt.Printf("repeat found in paper %d. previous index: %d. current index %d", outerIndex+1, innerIndex, innerIndex-1)
					innerIndex = innerIndex - 1
					continue
				}

			}
			questionSets[outerIndex] = append(questionSets[outerIndex], randQuestion)
		}
	}

}

func writeToFile() {
	file, err := os.Create("questions.txt")
	if err != nil {
		fmt.Println("error opening the file", err)
	} else {
		defer file.Close()
		for outerIndex := 0; outerIndex < papersNum; outerIndex++ {
			headerString := fmt.Sprintf("----------Paper number %d----------\n", outerIndex+1)
			_, err := file.WriteString(headerString)
			if err != nil {
				fmt.Println(err, "here")
			}
			for innerIndex := 0; innerIndex < 3; innerIndex++ {
				questionString := fmt.Sprintf("%d. %v\n", innerIndex+1, questionSets[outerIndex][innerIndex])
				_, err := file.WriteString(questionString)

				if err != nil {
					fmt.Println(err, "HERE")
				}
			}
		}
	}
}

// func checkForDoubles() {

// }

func getQuestionsNum() {

}

func getUserInput() {
	// The loop keeps going until we get a valid input from the user. When we do, we return from the function
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("How many papers to create?")
		input, _ := reader.ReadString('\n')
		input = strings.Trim(input, "\n")

		parsedNum, err := strconv.Atoi(input)
		// We want to get a positive number, so even if it's an int, we need to make sure it's positive and not equal to zero
		if err != nil || parsedNum <= 0 {
			fmt.Println("----------------")
			fmt.Println("Please enter a valid positive number of papers")
			continue
		}

		fmt.Println("----------------")
		if parsedNum == 1 {
			fmt.Println("Generating 1 paper")
		} else {
			fmt.Printf("Generating %d papers\n", parsedNum)
		}
		papersNum = parsedNum
		return
	}

}
