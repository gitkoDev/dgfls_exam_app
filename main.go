package main

import (
	"os"

	"github.com/gitkoDev/test-paper-generator/utils"
)

func init() {
	_, err := os.Stat("put-your-questions-here.txt")
	if err != nil {
		os.Create("put-your-questions-here.txt")
	}
}

func main() {
	utils.Onboarding()

	utils.ParseQuestions()

	utils.GetUserInput()

	utils.PopulateQuestionPapers()

	utils.WriteToFile()

}
