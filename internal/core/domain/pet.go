package domain

import "time"

type Pet struct {
	ID           uint64     `json:"id"`
	Name         string     `json:"name"`
	Age          uint64     `json:"age"`
	Weight       float64    `json:"weight"`
	Requirements string     `json:"requirements"`
	CreatedAt    time.Time  `json:"created_at,omitempty"`
	User         User       `json:"user"`
	Photos       []PetPhoto `json:"photos"`
}

type PetPhoto struct {
	ID       uint64 `json:"id,omitempty"`
	FileName string `json:"file_name,omitempty"`
	FilePath string `json:"file_path,omitempty"`
	PetID    uint64 `json:"pet_id,omitempty"`
}
