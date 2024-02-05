package main

import (
	"fmt"
	"math/rand"
	"os"
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

var questionSets [25][3]string

func main() {
	populateQuestionPapers()

	writeToFile()

}

func populateQuestionPapers() {
	// Populate question sets with random questions
	for outerIndex := 0; outerIndex < 25; outerIndex++ {
		for innerIndex := 0; innerIndex < 3; innerIndex++ {
			fmt.Println("inner index:", innerIndex)
			randQuestion := allQuestions[rand.Intn(len(allQuestions))]
			for _, q := range questionSets[outerIndex] {

				if q == randQuestion {
					// log.Fatalln("repeat found")
					fmt.Printf("repeat found. previous index: %d. current index %d", innerIndex, innerIndex-1)
					innerIndex = innerIndex - 1
					continue
				}

			}

			questionSets[outerIndex][innerIndex] = randQuestion

		}
	}

	// for i, q := range questionSets {
	// 	fmt.Printf("paper number %d", i+1)
	// 	for _, sq := range q {
	// 		// fmt.Println(sq)
	// 	}
	// }
}

func writeToFile() {
	file, err := os.Create("questions.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("error opening the file", err)
	} else {
		for outerIndex := 0; outerIndex < 25; outerIndex++ {
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

func checkForDoubles() {

}
