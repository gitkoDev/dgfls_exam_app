package main

import (
	"github.com/gitkoDev/dgfls_exam_app/utils"
)

func main() {
	utils.GetUserInput()

	utils.PopulateQuestionPapers()

	utils.WriteToFile()

}
