package dto

type GetSaveDelete struct {
	Url string `json:"url" validate:"required,url"`
}

type Update struct {
	Shortcut string `json:"shortcut" validate:"required,custom_shortcut"`
	Want     string `json:"want" validate:"required,custom_shortcut"`
}

type General struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}
