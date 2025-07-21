package request

type CreateStoryRequest struct {
	ID          string                 `json:"id,omitempty"`
	Title       string                 `json:"title" binding:"required"`
	Description string                 `json:"description" binding:"required"`
	Language    string                 `json:"language" binding:"required,oneof=tseltal zapoteco maya"`
	Category    string                 `json:"category" binding:"required"`
	ContentJSON map[string]interface{} `json:"content_json,omitempty"`
	Status      string                 `json:"status,omitempty"`
}
