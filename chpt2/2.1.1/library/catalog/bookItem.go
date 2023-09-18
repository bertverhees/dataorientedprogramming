package catalog

/*A book can have multiple copies, and each copy is considered as a book item*/

type BookItem struct {
	Id    string
	LibId string
}
