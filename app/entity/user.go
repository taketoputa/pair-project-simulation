package entity

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

type Staff struct {
	ID       int
	Name     string
	Email    string
	Position string
}

type Sale struct {
	ID        int
	ProductID int
	Quantity  int
	SaleDate  string
}
