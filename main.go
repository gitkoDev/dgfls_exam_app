package main

import (
	"github.com/gitkoDev/test-paper-generator/utils"
)

func main() {
	utils.ParseQuestions()

	utils.GetUserInput()

	utils.PopulateQuestionPapers()

	utils.WriteToFile()

}
