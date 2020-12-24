package model

import "time"

type SourceLost struct {
	ID        int64
	SourceID  int64
	BindISBN  string
	ErrIp     string
	Err       string
	CreatedAt *time.Time
}

func (m *SourceLost) TableName() string {
	return "source_lost"
}
