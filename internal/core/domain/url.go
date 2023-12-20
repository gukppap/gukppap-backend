package domain

// URL 단축 서비스라고 가정했을 떄
type URL struct {
	OriginalURL string `json:"orignalURL"`
	ShortcutURL string `json:"shortcutURL"`
}
