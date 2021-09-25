package entitys

import "time"

type Police struct {
	ID            string    `json:"id"`
	ApplicationID string    `json:"application_id"`
	PermmionID    string    `json:"permmion_id"`
	SubjectID     string    `json:"subject_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
