package library

/*A member of the library.*/

type MemberAble interface {
	Login() bool
	IsBlocked() bool
	Block() bool
	Unblock() bool
	ReturnBook(bookLending BookLendingAble) bool
	CheckOut(bookItem BookItem) BookLendingAble
}

type Member struct {
	Id       string
	Email    string
	Password string
}
