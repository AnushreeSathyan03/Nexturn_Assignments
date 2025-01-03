package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

// Product struct representing a product in inventory
type Product struct {
	ID       int
	Name     string
	UnitCost float64
	Stock    int
}

// Inventory slice to maintain all products
var inventory []Product

// AddProduct inserts a new product into the inventory
func AddProduct(id int, name string, cost float64, stock int) error {
	for _, product := range inventory {
		if product.ID == id {
			return errors.New("product ID already exists")
		}
	}

	inventory = append(inventory, Product{ID: id, Name: name, UnitCost: cost, Stock: stock})
	return nil
}

// UpdateStock adjusts the stock quantity for a product by ID
func UpdateStock(id int, updatedStock int) error {
	if updatedStock < 0 {
		return errors.New("stock cannot be negative")
	}
	for i, product := range inventory {
		if product.ID == id {
			inventory[i].Stock = updatedStock
			return nil
		}
	}
	return errors.New("product not found")
}

// SearchProduct retrieves a product by its ID or name
func SearchProduct(identifier string) (*Product, error) {
	for _, product := range inventory {
		if strings.EqualFold(product.Name, identifier) || fmt.Sprintf("%d", product.ID) == identifier {
			return &product, nil
		}
	}
	return nil, errors.New("product not found")
}

// ShowInventory lists all products in the inventory
func ShowInventory() {
	fmt.Printf("%-6s %-15s %-10s %-10s\n", "ID", "Name", "Unit Cost", "Stock")
	fmt.Println(strings.Repeat("-", 50))
	for _, product := range inventory {
		fmt.Printf("%-6d %-15s %-10.2f %-10d\n", product.ID, product.Name, product.UnitCost, product.Stock)
	}
}

// SortInventory arranges products based on a specific attribute (cost or stock)
func SortInventory(criteria string) {
	switch criteria {
	case "cost":
		sort.Slice(inventory, func(i, j int) bool {
			return inventory[i].UnitCost < inventory[j].UnitCost
		})
	case "stock":
		sort.Slice(inventory, func(i, j int) bool {
			return inventory[i].Stock < inventory[j].Stock
		})
	}
}

func main() {
	fmt.Println("Welcome to the Inventory Management System")

	// Preloaded products
	_ = AddProduct(201, "Eraser", 3.50, 150)
	_ = AddProduct(202, "Notebook", 45.00, 300)
	_ = AddProduct(203, "Ruler", 12.00, 90)
	_ = AddProduct(204, "Highlighter", 20.00, 60)

	for {
		fmt.Println("\nOptions:")
		fmt.Println("1. Add Product")
		fmt.Println("2. Update Stock")
		fmt.Println("3. Search Product")
		fmt.Println("4. Display Inventory")
		fmt.Println("5. Sort Inventory by Cost")
		fmt.Println("6. Sort Inventory by Stock")
		fmt.Println("7. Exit")
		fmt.Print("Select an option: ")

		var option int
		fmt.Scan(&option)

		switch option {
		case 1:
			var id, stock int
			var name string
			var cost float64
			fmt.Print("Enter Product ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter Product Name: ")
			fmt.Scan(&name)
			fmt.Print("Enter Product Unit Cost: ")
			fmt.Scan(&cost)
			fmt.Print("Enter Product Stock: ")
			fmt.Scan(&stock)
			if err := AddProduct(id, name, cost, stock); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Product added successfully.")
			}
		case 2:
			var id, stock int
			fmt.Print("Enter Product ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter Updated Stock: ")
			fmt.Scan(&stock)
			if err := UpdateStock(id, stock); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Stock updated successfully.")
			}
		case 3:
			var query string
			fmt.Print("Enter Product Name or ID: ")
			fmt.Scan(&query)
			if product, err := SearchProduct(query); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Found Product: %+v\n", *product)
			}
		case 4:
			ShowInventory()
		case 5:
			SortInventory("cost")
			fmt.Println("Inventory sorted by unit cost.")
			ShowInventory()
		case 6:
			SortInventory("stock")
			fmt.Println("Inventory sorted by stock.")
			ShowInventory()
		case 7:
			fmt.Println("Exiting the Inventory Management System. Goodbye!")
			return
		default:
			fmt.Println("Invalid selection. Please try again.")
		}
	}
}