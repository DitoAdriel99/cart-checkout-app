package product_repo

import (
	"database/sql"
	"fmt"
	"go-learn/entities"
	"go-learn/library/meta"
	"strings"
	"time"

	"github.com/google/uuid"
)

func (c *_ProductRepoImp) Create(pr *entities.Product) error {
	tx, err := c.conn.Begin()
	if err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	queryInsert := `INSERT INTO products (id, title, description, price, quantity, image, type, banner, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	_, err = tx.Exec(queryInsert, pr.ID, pr.Title, pr.Description, pr.Price, pr.Qty, pr.Image, pr.Type, pr.Banner, pr.CreatedAt, pr.UpdatedAt)
	if err != nil {
		err = fmt.Errorf("executing query: %w", err)
		return err
	}

	queryInfo := `INSERT INTO product_info (id, product_id, info, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`
	newId, _ := uuid.NewUUID()
	_, err = tx.Exec(queryInfo, newId, pr.ID, pr.Info, pr.CreatedAt, pr.UpdatedAt)
	if err != nil {
		err = fmt.Errorf("executing query: %w", err)
		return err
	}

	return nil
}

func (c *_ProductRepoImp) Detail(id uuid.UUID) (*entities.Product, error) {
	query := `
		SELECT 
			p.id, 
			p.title, 
			p.description, 
			p.price, 
			p.quantity,
			p.image, 
			p.type, 
			p.banner, 
			pi.info, 
			p.created_at, 
			p.updated_at 
		FROM 
			products p 
		JOIN 
			product_info pi 
		ON 
			p.id = pi.product_id 
		WHERE 
			p.id = $1`

	var object entities.Product

	err := c.conn.QueryRow(query, id).Scan(
		&object.ID,
		&object.Title,
		&object.Description,
		&object.Price,
		&object.Qty,
		&object.Image,
		&object.Type,
		&object.Banner,
		&object.Info,
		&object.CreatedAt,
		&object.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product not found")
		}
		err = fmt.Errorf("scanning activity objects: %w", err)
		return nil, err
	}

	return &object, nil
}

func (c *_ProductRepoImp) GetAll(m *meta.Metadata) ([]entities.Product, error) {
	q, err := meta.ParseMetaData(m, c)
	if err != nil {
		return nil, err
	}
	stmt := `SELECT p.id, p.title, p.description, p.price, p.quantity, p.image, p.type, p.banner, p.created_at, p.updated_at
	FROM products p
	`
	queries := QueryStatement(stmt)
	var (
		searchBy = q.SearchBy
		order    = q.OrderBy
	)
	if len(q.Search) > 2 {
		if len(q.SearchBy) != 0 {
			queries = queries.Where(searchBy, like, q.Search)
		}
	}
	if q.DateEnd.Valid && q.DateFrom.Valid {
		queries = queries.Where(order, between, q.DateFrom.Time.Local(), q.DateEnd.Time.Local())
	}

	query, _, args := queries.Order(order, direction(strings.ToUpper(q.OrderDirection))).
		Offset(q.Offset).
		Limit(q.Limit).Build()

	rows, err := c.conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	collections := make([]entities.Product, 0)
	for rows.Next() {
		var f entities.Product
		if err := rows.Scan(
			&f.ID,
			&f.Title,
			&f.Description,
			&f.Price,
			&f.Qty,
			&f.Image,
			&f.Type,
			&f.Banner,
			&f.CreatedAt,
			&f.UpdatedAt,
		); err != nil {
			return nil, err
		}

		collections = append(collections, f)
	}

	return collections, nil
}

func (c *_ProductRepoImp) Update(pr *entities.Product) error {
	tx, err := c.conn.Begin()
	if err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()
	query := `
		UPDATE products
		SET
			title = $2,
			description = $3,
			price = $4,
			quantity = $5,
			image = $6,
			type = $7,
			banner = $8,
			created_at = $9,
			updated_at = $10
		WHERE
			id = $1
	`

	_, err = tx.Exec(query, pr.ID, pr.Title, pr.Description, pr.Price, pr.Qty, pr.Image, pr.Type, pr.Banner, pr.CreatedAt, pr.UpdatedAt)
	if err != nil {
		err = fmt.Errorf("executing query: %w", err)
		return err
	}

	queryInfo := `
			UPDATE product_info 
			SET 
				info = $2, 
				updated_at = $3
			WHERE 
				product_id = $1`
	_, err = tx.Exec(queryInfo, pr.ID, pr.Info, pr.UpdatedAt)
	if err != nil {
		err = fmt.Errorf("executing query: %w", err)
		return err
	}

	return nil
}

func (c *_ProductRepoImp) Delete(id uuid.UUID) error {
	tx, err := c.conn.Begin()
	if err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()
	query := `DELETE FROM products WHERE id = $1`

	_, err = tx.Exec(query, id)
	if err != nil {
		err = fmt.Errorf("executing query update: %w", err)
		return err
	}

	query2 := `DELETE FROM product_info WHERE product_id = $1`

	_, err = tx.Exec(query2, id)
	if err != nil {
		err = fmt.Errorf("executing query update: %w", err)
		return err
	}
	return nil
}

func (c *_ProductRepoImp) AddToCart(payload entities.CartsPayload, email string) error {
	for _, v := range payload {
		checkQuery := `SELECT COUNT(*) FROM product_cart WHERE email = $1 AND product_id = $2 AND is_checkout is false`
		var count int
		err := c.conn.QueryRow(checkQuery, email, v.ProductsID).Scan(&count)
		if err != nil {
			err = fmt.Errorf("executing query: %w", err)
			return err
		}

		if count != 0 {
			return entities.ErrAlreadyInCart
		}

		queryInsert := `INSERT INTO product_cart (id, email, product_id, quantity,status, is_checkout, created_at, updated_at)
					VALUES($1,$2,$3,$4,$5,$6,$7,$8)`
		newId, _ := uuid.NewUUID()
		_, err = c.conn.Exec(queryInsert, newId, email, v.ProductsID, v.Qty, true, false, time.Now().Local(), time.Now().Local())
		if err != nil {
			err = fmt.Errorf("executing query insert: %w", err)
			return err
		}
	}
	return nil
}

func (c *_ProductRepoImp) GetCart(email string) ([]entities.Product, error) {
	query := `
		SELECT 
			pi.id, 
			pi.title, 
			pi.description, 
			pi.price, 
			pi.quantity,
			pi.image, 
			pi.type, 
			pi.banner,
			p.quantity,
			pi.created_at, 
			pi.updated_at 
		FROM 
			product_cart p 
		JOIN 
			products pi 
		ON 
			p.product_id = pi.id 
		WHERE 
			p.email = $1`

	rows, err := c.conn.Query(query, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	collections := make([]entities.Product, 0)
	for rows.Next() {
		var f entities.Product
		if err := rows.Scan(
			&f.ID,
			&f.Title,
			&f.Description,
			&f.Price,
			&f.Qty,
			&f.Image,
			&f.Type,
			&f.Banner,
			&f.QtyReq,
			&f.CreatedAt,
			&f.UpdatedAt,
		); err != nil {
			return nil, err
		}

		collections = append(collections, f)
	}
	return collections, nil
}
func (c *_ProductRepoImp) DeleteCart(email string, products_id []uuid.UUID) error {
	for _, v := range products_id {
		query := `DELETE FROM product_cart WHERE product_id = $1 and email = $2`
		_, err := c.conn.Exec(query, v, email)
		if err != nil {
			err = fmt.Errorf("executing query update: %w", err)
			return err
		}
	}
	return nil
}

func (c *_ProductRepoImp) Sortable(field string) bool {
	switch field {
	case "created_at", "updated_at":
		return true
	default:
		return false
	}
}
