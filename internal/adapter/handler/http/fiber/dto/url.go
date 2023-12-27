package dto

type CreateURLDTO struct {
	OriginalURL string `json:"originalURL,omitempty"`
	ShortcutURL string `json:"shortcutURL,omitempty"`
}

type GeneralDTO struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}
