package books 

type Book struct {
  ID int `json:"id"`
  Title string `json:"title"` 
  Description string `json:"description"`
  AuthorID int64 `json:"author_id,omitempty"`
  Author *Author `json:"author,omitempty"`
}

type Author struct {
  ID int `json:"id"`
  Name string `json:"name"` 
  Bio string `json:"biography"`
}