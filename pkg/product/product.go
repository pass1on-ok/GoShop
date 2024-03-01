package product

type Product struct {
	ID              int     `json:"id"`
	Name            string  `json:"name"`
	Price           float64 `json:"price"`
	Description     string  `json:"description"`
	QuantityInStock int     `json:"quantity_in_stock"`
}

func GetAllProducts() []Product {
	products := []Product{
		{ID: 1, Name: "Product 1", Price: 10.99, Description: "Description of Product 1", QuantityInStock: 100},
		{ID: 2, Name: "Product 2", Price: 19.99, Description: "Description of Product 2", QuantityInStock: 50},
		{ID: 3, Name: "Product 3", Price: 29.99, Description: "Description of Product 3", QuantityInStock: 200},
	}
	return products
}
