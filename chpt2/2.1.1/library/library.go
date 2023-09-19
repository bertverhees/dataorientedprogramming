package library

import (
	"errors"
)

type Library struct {
}

func GetBookLendings(libraryData LibraryData, userId string, memberId string) ([]BookLending, error) {
	if UserManagement.IsLibrarian()(libraryData.UserManagement, userId) {
		return Catalog.getBookLendings(libraryData.Catalog, memberId), nil
	} else {
		return nil, errors.New("Not allowed to get book lendings")
	}
}

func SearchBook(libraryData LibraryData, searchQuery string) Book {
	return nil
}

func AddBookItem(libraryData LibraryData, bookItemInfo BookItemInfo) bool {
	return false
}

func BlockMember(libraryData LibraryData, memberId string) bool {
	return false
}

func UnblockMember(libraryData LibraryData, memberId string) bool {
	return false
}

func login(libraryData LibraryData, loginInfo []string) bool {
	return false
}

func CheckOutBook(libraryData LibraryData, bookItemId string) bool {
	return false
}

func ReturnBook(libraryData LibraryData, userId, bookItemId string) bool {
	return false
}
