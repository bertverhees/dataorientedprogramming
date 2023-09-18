package library

/*When a book is lent, a book lending object is created*/

import "time"

type BookLendingAble interface {
	IsLate() bool
	ReturnBook() bool
}

type BookLending struct {
	Id          string
	LendingDate time.Time
	DueDate     time.Time
}
