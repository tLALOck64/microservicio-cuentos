package application

import (
	"github.com/tLALOck64/microservicio-cuentos/internal/question/domain/entities"
	"github.com/tLALOck64/microservicio-cuentos/internal/question/domain/ports"
)

type GetByStoryIdUseCase struct {
	QuestionRepository ports.QuestionRepository
}

func NewGetByStoryIdUseCase(questionRepository ports.QuestionRepository) *GetByStoryIdUseCase {
	return &GetByStoryIdUseCase{QuestionRepository: questionRepository}
}

func (g *GetByStoryIdUseCase) Run(storyID string) ([]*entities.Question, error) {
	questions, err := g.QuestionRepository.GetByStoryId(storyID)
	if err != nil {
		return []*entities.Question{}, err
	}
	return questions, nil
}
