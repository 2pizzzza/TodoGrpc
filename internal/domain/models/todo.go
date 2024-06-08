package models

import "time"

type Todo struct {
	ID          int
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
}

func (t Todo) Id() int {
	return t.ID
}

func (t Todo) Name() string {
	return t.Title
}

func (t Todo) Desc() string {
	return t.Description
}

func (t Todo) IsDone() bool {
	return t.Completed
}

func (t Todo) TimeCreat() time.Time {
	return t.CreatedAt
}
