package application

import (
	"github.com/tLALOck64/microservicio-cuentos/internal/story/domain/entities"
	"github.com/tLALOck64/microservicio-cuentos/internal/story/domain/ports"
)

type GetByCategoryUseCase struct {
	StoryRepository ports.StoryRepository
}

func NewGetByCategoryUseCase(storyRepository ports.StoryRepository) *GetByCategoryUseCase {
	return &GetByCategoryUseCase{StoryRepository: storyRepository}
}

func (s *GetByCategoryUseCase) Run(category string) ([]*entities.Story, error) {
	stories, err := s.StoryRepository.GetByCategory(category)
	if err != nil {
		return []*entities.Story{}, err
	}
	return stories, nil
} 