package model

type Todo struct {
	ID int64
	UserID int64
	Title  string
	Description *string
	IsCompleted *bool
}
