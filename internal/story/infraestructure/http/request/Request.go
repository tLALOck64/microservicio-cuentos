package request

type CreateStoryRequest struct {
	ID          string                 `json:"id" binding:"required"`
	Title       string                 `json:"title" binding:"required"`
	Description string                 `json:"description" binding:"required"`
	Language    string                 `json:"language" binding:"required,oneof=tzeltal zapoteco maya"`
	Category    string                 `json:"category" binding:"required"`
	ContentJSON map[string]interface{} `json:"content_json,omitempty"`
	Status      string                 `json:"status,omitempty"`
}
