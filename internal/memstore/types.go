package memstore

type Taco struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ImageUrl    string  `json:"image_url"`
	Price       float32 `json:"price"`
}

// {"first_name":"aa","last_name":"","email":"","address":"","city":"","state":"","postal_code":"","country":"","save_info":false}
type Address struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Address    string `json:"address"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
	SaveInfo   bool   `json:"save_info"`
}
