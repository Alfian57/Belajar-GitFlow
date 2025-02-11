package main

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (p *Product) Create() {
	// Create product
}
