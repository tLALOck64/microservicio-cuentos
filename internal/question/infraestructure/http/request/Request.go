package request

type CreateQuestionRequest struct {
	ID         string   `json:"id,omitempty"`
	StoryID    string   `json:"story_id" binding:"required"`
	Question   string   `json:"question" binding:"required"`
	Answer     string   `json:"answer" binding:"required"`
	Type       string   `json:"type" binding:"required,oneof=multiple_choice true_false open_ended fill_blank"`
	Difficulty string   `json:"difficulty" binding:"required,oneof=easy medium hard"`
	Points     int      `json:"points" binding:"required,min=0"`
	IsActive   bool     `json:"is_active,omitempty"`
	Options    []string `json:"options,omitempty"`
}

type UpdateQuestionRequest struct {
	Question   string   `json:"question" binding:"required"`
	Answer     string   `json:"answer" binding:"required"`
	Type       string   `json:"type" binding:"required,oneof=multiple_choice true_false open_ended fill_blank"`
	Difficulty string   `json:"difficulty" binding:"required,oneof=easy medium hard"`
	Points     int      `json:"points" binding:"required,min=0"`
	IsActive   bool     `json:"is_active"`
	Options    []string `json:"options,omitempty"`
}
