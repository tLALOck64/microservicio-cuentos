package application

import (
	"github.com/tLALOck64/microservicio-cuentos/internal/story/domain/entities"
	"github.com/tLALOck64/microservicio-cuentos/internal/story/domain/ports"
)

type GetUseCase struct {
	StoryRepository ports.StoryRepository
}

func NewGetUseCase(storyRepository ports.StoryRepository)*GetUseCase{
	return &GetUseCase{StoryRepository: storyRepository}
}

func (s *GetUseCase) Run()([]*entities.Story, error){
	storys, err := s.StoryRepository.Get()

	if err != nil {
		return []*entities.Story{}, err
	}
	
	return storys, nil
}