package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/fumiama/go-docx"
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

const (
	Papers    = "papers"
	Questions = "questions"
)

var allQuestions = strings.Split(originalQuestions, "\n")

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
		for innerIndex := 0; innerIndex < questionsNum; innerIndex++ {
			doubleFound := false

			randQuestion := allQuestions[rand.Intn(len(allQuestions))]
			// Check for doubles inside the paper before appending
			for _, q := range questionSets[outerIndex] {
				if q == randQuestion {
					doubleFound = true
					innerIndex = innerIndex - 1
				}
			}
			if doubleFound {
				continue
			}
			questionSets[outerIndex] = append(questionSets[outerIndex], randQuestion)
		}
	}

}

func writeToFile() {
	file, err := os.Create("questions.docx")
	if err != nil {
		log.Fatalln(err)
	}

	cont := docx.NewA4()

	defer file.Close()
	for outerIndex := 0; outerIndex < papersNum; outerIndex++ {
		headerString := fmt.Sprintf("----------Paper number %d----------\n", outerIndex+1)
		parag := cont.AddParagraph().Justification("center")
		parag.AddText(headerString).Size("40").Bold().Font("Arial", "Arial", "Arial")

		for innerIndex := 0; innerIndex < questionsNum; innerIndex++ {
			questionString := fmt.Sprintf("%d. %v\n", innerIndex+1, questionSets[outerIndex][innerIndex])
			parag := cont.AddParagraph()
			parag.AddText(questionString).Size("30").Font("Arial", "Arial", "Arial")

		}
		cont.AddParagraph()
		cont.AddParagraph()
	}

	_, err = cont.WriteTo(file)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Done! %d papers with %d questions in each one are ready. Get the results in the txt file\n", papersNum, questionsNum)

}

func checkForDoubles() {

}

func getUserInput() {
	papersNum = parseUserInput(Papers)
	questionsNum = parseUserInput(Questions)
}

func parseUserInput(valueToParse string) int {
	// The loop keeps going until we get a valid input from the user. When we do, we return from the function
	for {
		reader := bufio.NewReader(os.Stdin)
		// Different message depending on whether we want the user to input amount of papers or questions in each paper
		if valueToParse == Papers {
			fmt.Println("How many papers to create?")
		} else if valueToParse == Questions {
			fmt.Println("How many questions in each paper?")
		}
		input, _ := reader.ReadString('\n')
		input = strings.Trim(input, "\n")

		parsedNum, err := strconv.Atoi(input)
		// We want to get a positive number, so even if it's an int, we need to make sure it's positive and not equal to zero
		if err != nil || parsedNum <= 0 {
			fmt.Println("----------------")
			// Different message depending on whether we want the user to input amount of papers or questions in each paper
			fmt.Printf("Please enter a valid positive number of %s\n", valueToParse)
			continue
		}

		fmt.Println("----------------")
		return parsedNum
	}

}
