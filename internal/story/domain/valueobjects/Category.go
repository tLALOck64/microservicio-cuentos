package valueobjects

import "fmt"

type Category string

const (
	Legends      Category = "legends"
	Fables       Category = "fables"
	Nature       Category = "nature"
	Daily        Category = "daily_life"
	Historical   Category = "historical"
	Mythological Category = "mythological"
	Educational  Category = "educational"
)

func NewCategory(value string) (Category, error) {
	category := Category(value)

	if !category.IsValid() {
		return "", fmt.Errorf("categoría no válida: %s", value)
	}

	return category, nil
}

func (c Category) IsValid() bool {
	validCategory := []Category{
		Legends, Fables, Nature, Daily, Historical,
		Mythological, Educational,
	}

	for _, valid := range validCategory {
		if c == valid {
			return true
		}
	}

	return false
}

func (c Category) String() string {
	return string(c)
}
