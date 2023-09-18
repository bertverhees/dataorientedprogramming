package library

/*Contains a list of books.*/

type CatalogAble interface {
	Search(searchCriteria, queryString string) []Book
	AddBookItem(librarian LibrarianAble, bookItem BookItemAble) BookItemAble
}

type Catalog struct {
}
