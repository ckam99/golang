package books

type Book struct {
  ID int64 `json:"id"`
  Title string `json:"title"`
  Description string `json:"description"`
  AuthorID string `json:"author_id,omitempty"`
  Author *Author `json:"author,omitempty"`
}

type Author struct {
  ID int64 `json:"id"`
  Name string `json:"name"`
}

type FilterParam struct {
  Limit int64 `json:"limit"`
  Offset int64 `json:"offset"`
  OrderBy *string `nil=true > one_of=id,title,author_id`
}