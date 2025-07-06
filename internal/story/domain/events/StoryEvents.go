package events

import "time"

// StoryPublished representa el evento cuando un cuento es publicado
type StoryPublished struct {
	StoryID     string    `json:"story_id"`
	Title       string    `json:"title"`
	Language    string    `json:"language"`
	PublishedAt time.Time `json:"published_at"`
}

// StoryCreated representa el evento cuando un cuento es creado
type StoryCreated struct {
	StoryID   string    `json:"story_id"`
	Title     string    `json:"title"`
	Language  string    `json:"language"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"created_at"`
}

// StoryArchived representa el evento cuando un cuento es archivado
type StoryArchived struct {
	StoryID    string    `json:"story_id"`
	Title      string    `json:"title"`
	ArchivedAt time.Time `json:"archived_at"`
}
