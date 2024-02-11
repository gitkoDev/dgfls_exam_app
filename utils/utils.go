package utils

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/fumiama/go-docx"
)

// var testQ = `- Can you think of a really boring hobby?
// 	- Why do you think it's important to have a hobby?
// 	- Do you think having lots of hobbies is bad?
// 	- What hobbies do your family members have?
// 	- Do you prefer indoor or outdoor hobbies and why?
// 	- How often do you leave Dongguan?
// 	- What do you usually enjoy eating when you’re sad?
// 	- How often do you watch movies?
// 	-  What kind of food have you never eaten before?
// 	- What do you think is more difficult, learning a new language or a musical instrument, and why do you think so?
// 	- Which season (time of the year) do you think is the best? Why?
// 	- Is your hometown colder than Dongguan? How are they different?
// 	- What is the most beautiful city you have ever visited?
// 	- What is your dream job? Why?
// 	- What do you think is the most boring job? Why?
// 	- When you get a job, how much money would you like to earn?
// 	- What kind of job do you think is better: part-time or full time? Why?
// 	- What dishes can you cook? What is your best one?
// 	- Who do you think cooks the most delicious food?
// 	- What is the best way to cook eggs? Describe the steps
// 	- How often do you cook? Do you mostly cook breakfast, lunch or dinner?`

var originalQuestions string

var parsedQuestions []string

const (
	Papers    = "papers"
	Questions = "questions"
)

var questionSets [][]string
var papersNum int
var questionsNum int

// Step 0: explain to user how program works

func Onboarding() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("----------------")
	fmt.Println("This program will help you generate random exam papers with the questions you provide. Please ENTER to proceed")
	reader.ReadString('\n')
	fmt.Println("First, make sure to put your questions into 'put-your-questions-here' txt file in the folder. Each question should be put on a new line")
	reader.ReadString('\n')
}

// Step 1: Get user's questions

func ParseQuestions() {
	// The loop only ends when the user provides questions to the txt file. No questions = no moving forward
	for {
		file, err := os.Open("put-your-questions-here.txt")
		if err != nil {
			log.Fatalln("error opening file with user's questions:\n", err)
		}

		defer file.Close()

		content, err := io.ReadAll(file)
		if err != nil {
			log.Fatalln("error reading file with user's questions:\n", err)
		}

		originalQuestions = string(content)

		parsedQuestions = strings.Split(originalQuestions, "\n")

		if len(parsedQuestions) <= 1 {
			fmt.Println("Please provide more questions to the txt file in the folder. Press ENTER where you're ready")

			reader := bufio.NewReader(os.Stdin)
			reader.ReadString('\n')
			continue
		} else {
			fmt.Printf("%d questions detected\n", len(parsedQuestions))
			fmt.Println("----------------")
			return
		}

	}

}

// Step 2: get user input - how many papers and questions in each paper

func GetUserInput() {
	papersNum = parseUserInput(Papers)
	questionsNum = parseUserInput(Questions)
}

func parseUserInput(valueToParse string) int {
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

		if valueToParse == Questions && parsedNum > len(parsedQuestions) {
			fmt.Println("----------------")
			// The number of questions in each paper should not be greater than total number of questions provided by user
			fmt.Printf("You have only provided %d questions. Please make sure the number of questions in each paper is equal or less than that\n", len(parsedQuestions))
			continue
		}

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

// Step 3: populate question sets with data provided by user

func PopulateQuestionPapers() {
	// Outer index = papers, inner index = questions in each consequtive paper
	for outerIndex := 0; outerIndex < papersNum; outerIndex++ {
		questionSets = append(questionSets, []string{})
		for innerIndex := 0; innerIndex < questionsNum; innerIndex++ {
			doubleFound := false

			randQuestion := parsedQuestions[rand.Intn(len(parsedQuestions))]
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

// Step 4: write to file

func WriteToFile() {
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
