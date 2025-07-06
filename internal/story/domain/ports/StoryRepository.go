package ports

import "github.com/tLALOck64/microservicio-cuentos/internal/story/domain/entities"

type StoryRepository interface {
	Create(story *entities.Story) (*entities.Story, error)
	Get() ([]*entities.Story, error)
	GetById(id string) (*entities.Story, error)
}