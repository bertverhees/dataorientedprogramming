package library

/*A book can have multiple copies, and each copy is considered as a book item*/

type BookItemAble interface {
	CheckOut(member MemberAble) BookLendingAble
}

type BookItem struct {
	Id    string
	LibId string
}
