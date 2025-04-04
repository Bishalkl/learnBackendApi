package product

import (
	"database/sql"
	"fmt"

	"github.com/bishalkl/learnBackendApi/types"
)

// creating type
type Store struct {
	db *sql.DB
}

// creating new instance
func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// function for get All Product
func (s *Store) GetProducts() ([]types.Product, error) {
	// SQL query to fetch all product details
	query := "SELECT id , name , description, image, price, quantity, createdAt FROM products"

	// Execute the query
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err //Return an error if the query fails
	}
	defer rows.Close() // Ensure rows are closed after function execution

	var products []types.Product

	// Iterate through the result set
	for rows.Next() {
		var product types.Product
		// Scan row data int the product struct
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Image,
			&product.Price,
			&product.Quantity,
			&product.CreatedAt,
		)
		if err != nil {
			return nil, err //Return an error if scanning fails
		}

		products = append(products, product)
	}

	// check for iteration errors
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

// function for get Product by id
func (s *Store) GetProductById(id int) (*types.Product, error) {
	// SQL query to fetch a product by ID
	query := "SELECT id ,name ,description, image, price, quantity, createdAt FROM products WHERE id = ?"

	// Execute the query with given ID
	rows := s.db.QueryRow(query, id)

	var product types.Product

	// Scan the  result into product struct
	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Image,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil //Return an empty product if not found
		}
		return nil, err //Return an error for other issues
	}
	return &product, nil //Return the found product
}

// CreateProduct inserts a new product into the database
func (s *Store) CreateProduct(product *types.Product) error {
	// SQL query to insert a new product
	query := "INSERT INTO products (name, description, image, price, quantity, createdAt) VALUES (?, ?, ?, ?, ?, ?)"

	// Execute the query
	_, err := s.db.Query(query, product.Name, product.Description, product.Image, product.Price, product.Quantity, product.CreatedAt)
	if err != nil {
		return fmt.Errorf("faild to insert product: %v", err)
	}

	return nil
}
