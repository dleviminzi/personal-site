package data

import (
	"html/template"
	"time"
)

// ExperienceItem gives skeleton for a work experience entry
type ExperienceItem struct {
	ID          int       `json:"id"`
	ItemType    string    `json:"item_type"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}

// Bio is background info on me
type Bio struct {
	ID        int    `json:"id"`
	Body      string `json:"body"`
	PhotoPath string `json:"photo_path"`
}

// Project has start date information, but no end date information
type Project struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	GithubLink  string    `json:"github_link"`
	Status      string    `json:"status"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}

// Note contains content regarding a specific subject
type Note struct {
	Title   string        `json:"title"`
	Topic   string        `json:"topic"`
	Content template.HTML `json:"content"`
}

// Photo provides metadata for a photo
type Photo struct {
	Title       string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Created     string `json:"created"`
	PhotoPath   string `json:"photo_path"`
}
