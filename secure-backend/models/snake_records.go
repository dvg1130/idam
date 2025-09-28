package models

type Snake struct {
	SnakeId string `json:"sid"`
	Species string `json:"species"`
	Sex     string `json:"sex"`
	Age     int    `json:"age"`
	Genes   string `json:"genes"`
	Notes   string `json:"notes"`
}

type SnakesListItem struct {
	Sid     string `json:"sid"`
	Species string `json:"species"`
}
type UpdateSnake struct {
	Sid     string  `json:"sid"` // required
	Species *string `json:"species,omitempty"`
	Sex     *string `json:"sex,omitempty"`
	Age     *int    `json:"age,omitempty"`
	Genes   *string `json:"genes,omitempty"`
	Notes   *string `json:"notes,omitempty"`
}

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

type SnakeSid struct {
	Sid string `json:"sid"`
}

type HealthRecord struct {
	Sid       string `json:"sid"`
	CheckDate string `json:"check_date"`
	Weight    string `json:"weight"`
	Length    string `json:"length"`
	Topic     string `json:"topic"`
	Notes     string `json:"notes"`
}

type UpdateHealthRecord struct {
	Sid       string `json:"sid"`        // snake id (user-given)
	CheckDate string `json:"check_date"` // YYYY-MM-DD
	Weight    string `json:"weight,omitempty"`
	Length    string `json:"length,omitempty"`
	Topic     string `json:"topic,omitempty"`
	Notes     string `json:"notes,omitempty"`
}

type DeleteSnakeHealth struct {
	Sid       string `json:"sid"`
	CheckDate string `json:"check_date"`
}
