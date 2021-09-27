package entities

// ShortURL represents the short URL for this application
//
// swagger:model ShortURL
type ShortURL struct {
	// min: 1
	ID int64 `json:"id"`

	// unique: true
	LongURL string `json:"long_url" db:"long_url"`

	// unique: true
	Token string `json:"token"`

	// min: 0
	Visits int64 `json:"visits"`
}
