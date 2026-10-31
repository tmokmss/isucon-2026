package isureal

import (
	"time"
)

type PostingWindow struct {
	ID             string
	Region         string
	LocalDate      time.Time
	StartsAt       time.Time
	OnTimeDeadline time.Time
	EndedAt        *time.Time
}
