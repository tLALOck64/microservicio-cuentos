package valueobjects

import "fmt"

type Status string

const (
	Draft     Status = "draft"
	Published Status = "published"
	Archived  Status = "archived"
	Inactive  Status = "inactive"
)

func NewStatus(value string) (Status, error) {
	status := Status(value)

	if !status.IsValid() {
		return "", fmt.Errorf("estado no válido: %s. Estados válidos: draft, published, archived, inactive", value)
	}

	return status, nil
}

func (s Status) IsValid() bool {
	validStatuses := []Status{
		Draft, Published, Archived, Inactive,
	}

	for _, valid := range validStatuses {
		if s == valid {
			return true
		}
	}

	return false
}

func (s Status) String() string {
	return string(s)
}

func (s Status) IsDraft() bool {
	return s == Draft
}

func (s Status) IsPublished() bool {
	return s == Published
}

func (s Status) IsArchived() bool {
	return s == Archived
}

func (s Status) IsInactive() bool {
	return s == Inactive
}
