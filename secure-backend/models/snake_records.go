package models

type SnakeFeed struct {
	Sid      string `json:"sid"`       // snake id (user-given)
	FeedDate string `json:"feed_date"` // YYYY-MM-DD
	PreyType string `json:"prey_type"`
	PreySize string `json:"prey_size"`
	Notes    string `json:"notes"`
}

type SnakeFeedDeleteRequest struct {
	Sid  string `json:"sid"`       // snake ID or name
	Date string `json:"feed_date"` // feeding record date (match the format stored in DB)
}

type UpdateSnakeFeed struct {
	Sid      string  `json:"sid"`
	FeedDate string  `json:"feed_date"` // YYYY-MM-DD
	PreyType *string `json:"prey_type,omitempty"`
	PreySize *string `json:"prey_size,omitempty"`
	Notes    *string `json:"notes,omitempty"`
}
