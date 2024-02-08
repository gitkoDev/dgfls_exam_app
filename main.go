package main

import (
	"github.com/gitkoDev/dgfls_exam_app/utils"
)

func main() {
	utils.ParseQuestions()

	utils.GetUserInput()

	utils.PopulateQuestionPapers()

	utils.WriteToFile()

}
