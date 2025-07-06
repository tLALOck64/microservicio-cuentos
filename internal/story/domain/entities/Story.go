package entities

import "github.com/tLALOck64/microservicio-cuentos/internal/story/domain/valueobjects"

type Story struct {
	ID          string
	Title       string
	Description string
	Language    valueobjects.Language
	Category    valueobjects.Category
	ContentJSON map[string] interface{}
	Status valueobjects.Status
}

