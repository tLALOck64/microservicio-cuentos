package application

import (
	"github.com/tLALOck64/microservicio-cuentos/internal/question/domain/entities"
	"github.com/tLALOck64/microservicio-cuentos/internal/question/domain/ports"
)

type GetByIdUseCase struct {
	QuestionRepository ports.QuestionRepository
}

func NewGetByIdUseCase(questionRepository ports.QuestionRepository) *GetByIdUseCase {
	return &GetByIdUseCase{QuestionRepository: questionRepository}
}

func (g *GetByIdUseCase) Run(id string) (*entities.Question, error) {
	question, err := g.QuestionRepository.GetById(id)
	if err != nil {
		return &entities.Question{}, err
	}
	return question, nil
}
