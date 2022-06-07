package delivery

import (
	"context"
	"fmt"
	"log"

	"github.com/learn/simple-app/advance/survey/repository"
	"github.com/learn/simple-app/advance/survey/usecase"
)

type SurveyHandler struct {
	SurveyUsecase usecase.SurveyUsecase
}

func NewSurveyHandler(surveyUsecase usecase.SurveyUsecase) SurveyHandler {
	return SurveyHandler{SurveyUsecase: surveyUsecase}
}

func (sh SurveyHandler) StartSurvey() {
	var userData repository.User

	fmt.Println("What is your name? ")
	fmt.Scanln(&userData.Name)
	fmt.Println("How old are you?")
	fmt.Scanln(&userData.Age)
	fmt.Println("What is your hair color?")
	fmt.Scanln(&userData.HairColor)
	fmt.Println()

	err := sh.SurveyUsecase.InsertSurvey(context.Background(), userData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("=== Survey Report ===")

	resultSurvey, err := sh.SurveyUsecase.GetSurvey(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range resultSurvey {
		fmt.Println("User : ", v.UserID)
		fmt.Println("Question 1: What is your name?")
		fmt.Println("Answer 1 : ", v.Name)
		fmt.Println("Question 2: How old are you?")
		fmt.Println("Answer 2: ", v.Age)
		fmt.Println("Question 3: What is your hair color?")
		fmt.Println("Answer 3: ", v.HairColor)
		fmt.Println()
	}
}
