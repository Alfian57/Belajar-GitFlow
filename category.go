package main

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (c *Category) Create() {
	// Create category
}

func (c *Category) Update() {
	// Update category
}
