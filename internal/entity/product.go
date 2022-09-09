package entity

type Brand struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type Category struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Parent   *Category
	ParentID int `db:"parent_id default:null"`
}

type Product struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Price      float64 `db:"price"`
	Category   *Category
	CategoryID int `db:"category_id"`
	Brand      *Brand
	BrandID    int `db:"brand_id"`
}

type Store struct {
	ID      int    `db:"id"`
	Name    string `db:"name"`
	Address string `db:"address"`
}
