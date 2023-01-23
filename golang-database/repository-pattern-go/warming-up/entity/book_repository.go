package repository

type BookRepository interface {
	func InsertBook(ctx context.Cotext, book *entity.Book) (*entity.Book, error)
	
}