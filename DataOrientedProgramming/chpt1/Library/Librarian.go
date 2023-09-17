package library

/*A librarian.*/

type LibrarianAble interface {
	Login() bool
	BlockMember(member MemberAble) bool
	UnblockMember(member MemberAble) bool
	AddBookItem(bookItem BookItemAble) BookItemAble
	GetBooksLendingsOfMember(member MemberAble) []BookLendingAble
}

type Librarian struct {
	Id       string
	Email    string
	Password string
}
