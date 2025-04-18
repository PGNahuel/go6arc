package dto

type BasicRequestDTO struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
type BasicResponseDTO struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}
