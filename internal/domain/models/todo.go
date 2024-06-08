package models

import "time"

type Model struct {
	ID          int `json:"id,omitempty"`
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time `json:"created-at,omitempty"`
}

func (t Model) Id() int {
	return t.ID
}

func (t Model) Name() string {
	return t.Title
}

func (t Model) Desc() string {
	return t.Description
}

func (t Model) IsDone() bool {
	return t.Completed
}

func (t Model) TimeCreat() time.Time {
	return t.CreatedAt
}
