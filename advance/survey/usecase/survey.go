package usecase

import (
	"context"

	"github.com/learn/simple-app/advance/survey/repository"
)

type surveyUsecase struct {
	UserRepo repository.UserRepository
}

func NewSurveyUsecase(userRepo repository.UserRepository) SurveyUsecase {
	return surveyUsecase{
		UserRepo: userRepo,
	}
}

func (s surveyUsecase) InsertSurvey(ctx context.Context, input repository.User) error {
	return s.UserRepo.InsertUserData(ctx, input)
}

func (s surveyUsecase) GetSurvey(ctx context.Context) (res []repository.User, err error) {
	return s.UserRepo.GetUserData(ctx)
}

type SurveyUsecase interface {
	InsertSurvey(ctx context.Context, input repository.User) error
	GetSurvey(ctx context.Context) (res []repository.User, err error)
}
