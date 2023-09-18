package catalog

type Catalog struct {
	book Book
}

func SearchBook(catalogData Catalog, searchQuery string) Book {
	return nil
}

func AddBookItem(catalogData Catalog, bookItemInfo BookItem) bool {
	return false
}

func CheckOutBook(catalogData Catalog, bookItemId string) bool {
	return false
}

func ReturnBook(catalogData Catalog, searchQuery string) bool {
	return nil
}

func GetBookLendings(catalogData Catalog, userId string, memberId string) []library.BookLending {

}
