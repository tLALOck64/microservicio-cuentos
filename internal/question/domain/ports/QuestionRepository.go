package ports

import "github.com/tLALOck64/microservicio-cuentos/internal/question/domain/entities"

type QuestionRepository interface {
	Create(question *entities.Question) (*entities.Question, error)
	GetById(id string) (*entities.Question, error)
	Get() ([]*entities.Question, error)
	GetByStoryId(storyID string) ([]*entities.Question, error)
	Update(question *entities.Question) (*entities.Question, error)
	Delete(id string) error
}
