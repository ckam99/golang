package domain



type BookRepository interface {
   GetAll(ctx context.Context) ([]Book, error)
   Find(ctx context.Context ,id int) (Book, error)
   Store(ctx context.Context, book Book) error
}

type BookService interface {
  GetAll(ctx context.Context) ([]Book, error)
  Store(ctx context.Context, dto StoreBookDTO) (Book, error)
}