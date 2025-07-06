package application

import (
	"github.com/tLALOck64/microservicio-cuentos/internal/story/domain/entities"
	"github.com/tLALOck64/microservicio-cuentos/internal/story/domain/ports"
)

type CreateUseCase struct {
	StoryRepository ports.StoryRepository
}

func NewCreateUseCase(storyRepository ports.StoryRepository) *CreateUseCase{
	return &CreateUseCase{StoryRepository: storyRepository}
}

func (s *CreateUseCase) Run(story *entities.Story) (*entities.Story, error){
	newStory, err := s.StoryRepository.Create(story)

	if err != nil{
		return &entities.Story{}, err
	}

	return newStory, nil
}