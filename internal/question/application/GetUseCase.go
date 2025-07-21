package application

import (
	"github.com/tLALOck64/microservicio-cuentos/internal/question/domain/entities"
	"github.com/tLALOck64/microservicio-cuentos/internal/question/domain/ports"
)

type GetUseCase struct {
	QuestionRepository ports.QuestionRepository
}

func NewGetUseCase(questionRepository ports.QuestionRepository) *GetUseCase {
	return &GetUseCase{QuestionRepository: questionRepository}
}

func (g *GetUseCase) Run() ([]*entities.Question, error) {
	questions, err := g.QuestionRepository.Get()
	if err != nil {
		return []*entities.Question{}, err
	}
	return questions, nil
}
