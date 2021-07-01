package entities

type ShortURL struct {
	ID      int64  `json:"id"`
	LongURL string `json:"long_url" db:"long_url"`
	Token   string `json:"token"`
	Visits  int64  `json:"visits"`
}
