package library

/*A base class for Librarian and Member.*/

type UserAble interface {
	Login() bool
}

type User struct {
	Id       string
	Email    string
	Password string
}
