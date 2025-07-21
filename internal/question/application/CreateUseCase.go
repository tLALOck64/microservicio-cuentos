package application

import (
	"github.com/tLALOck64/microservicio-cuentos/internal/question/domain/entities"
	"github.com/tLALOck64/microservicio-cuentos/internal/question/domain/ports"
)

type CreateUseCase struct {
	QuestionRepository ports.QuestionRepository
}

func NewCreateUseCase(questionRepository ports.QuestionRepository) *CreateUseCase {
	return &CreateUseCase{QuestionRepository: questionRepository}
}

func (c *CreateUseCase) Run(question *entities.Question) (*entities.Question, error) {
	newQuestion, err := c.QuestionRepository.Create(question)
	if err != nil {
		return &entities.Question{}, err
	}
	return newQuestion, nil
}
