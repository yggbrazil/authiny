package dtos

import "time"

type Police struct {
	ID            string
	ApplicationID string
	PermmionID    string
	SubjectID     string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
