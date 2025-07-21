package events

import "time"

type QuestionCreated struct {
	QuestionID string
	StoryID    string
	Type       string
	Difficulty string
	CreatedAt  time.Time
}

type QuestionDeactivated struct {
	QuestionID    string
	StoryID       string
	DeactivatedAt time.Time
}

type QuestionActivated struct {
	QuestionID  string
	StoryID     string
	ActivatedAt time.Time
}
