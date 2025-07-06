package application

import (
	"github.com/tLALOck64/microservicio-cuentos/internal/story/domain/entities"
	"github.com/tLALOck64/microservicio-cuentos/internal/story/domain/ports"
)

type GetByIdUseCase struct {
	StoryRepository ports.StoryRepository
}

func NewGetByIdUseCase(storyRepository ports.StoryRepository) *GetByIdUseCase{
	return &GetByIdUseCase{StoryRepository: storyRepository}
}

func (s *GetByIdUseCase) Run(id string) (*entities.Story, error){
	story, err := s.StoryRepository.GetById(id)

	if err != nil{
		return &entities.Story{}, err
	}

	return story, nil
}