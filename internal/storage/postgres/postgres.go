package postgres

import (
	"WbTech0/internal/model"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.postgres.New"

	db, err := sql.Open("postgres", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s:%v", op, err)
	}
	stmt, err := db.Prepare(`SELECT EXISTS (SELECT FROM public.orders)`)
	if err != nil {
		return nil, fmt.Errorf("%s:%w", op, err)
	}
	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) GetOrderByTrack(track string) (model.Order, error) {
	const op = "storage.postgres.GetOrderByTrack"

	var order model.Order
	var orderId int64

	err := s.db.QueryRow("SELECT order_uid, track_number, entry,locale,internal_signature,customer_id,delivery_service,shardkey,sm_id,date_created,oof_shard,payments_id,delivery_id,id FROM orders WHERE order_uid=$1", track).Scan(&order.OrderUid, &order.TrackNumber, &order.Entry, &order.Locale, &order.InternalSignature, &order.CustomerId, &order.DeliveryService, &order.Shardkey, &order.SmId, &order.DateCreated, &order.OofShard, &orderId)

	switch {
	case err == sql.ErrNoRows:
		return model.Order{}, fmt.Errorf("%s:%w", op, err)
	case err != nil:
		return model.Order{}, fmt.Errorf("%s:%w", op, err)
	}

	return order, nil
}

func (s *Storage) GetOrderModels() []model.Order {
	const op = "storage.postgres.GetOrderModels"
	var orders []model.Order

	rows, _ := s.db.Query(`SELECT id,order_uid, track_number, entry,locale,internal_signature,customer_id,delivery_service,shardkey,sm_id,date_created,oof_shard,payments_id,delivery_id FROM orders`)
	defer rows.Close()
	for rows.Next() {
		var order model.Order
		var id int64
		var payment int64
		var deliveries int64
		if err := rows.Scan(&id, &order.OrderUid, &order.TrackNumber, &order.Entry, &order.Locale, &order.InternalSignature, &order.CustomerId, &order.DeliveryService, &order.Shardkey, &order.SmId, &order.DateCreated, &order.OofShard, &payment, &deliveries); err != nil {
			fmt.Errorf("%s:%w", op, err)
		} else {
			order.Deliver, _ = s.GetDeliverByID(deliveries)
			order.Payment, _ = s.GetPaymentByID(payment)
			order.Product, _ = s.GetProductByID(id)

			orders = append(orders, order)
		}
	}

	return orders
}
func (s *Storage) GetProductByID(orderId int64) ([]model.Product, error) {
	const op = "storage.postgres.GetPaymentByID"
	var products []model.Product
	rows, err := s.db.Query(`SELECT chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status FROM products WHERE order_id = $1`, orderId)
	if err != nil {
		return products, fmt.Errorf("%s:%w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ChrtId, &product.TrackNumber, &product.Price, &product.Rid, &product.Name, &product.Sale, &product.Size, &product.TotalPrice, &product.NmId, &product.Brand, &product.Status); err != nil {
			return products, fmt.Errorf("%s:%w", op, err)
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return products, fmt.Errorf("%s:%w", op, err)
	}
	return products, nil
}

func (s *Storage) GetPaymentByID(id int64) (model.Payment, error) {
	const op = "storage.postgres.GetPaymentByID"
	var payment model.Payment
	if err := s.db.QueryRow(`SELECT transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee FROM payments WHERE id = $1`, id).Scan(&payment.Transaction, &payment.RequestId, &payment.Currency, &payment.Provider, &payment.Amount, &payment.PaymentDt, &payment.Bank, &payment.DeliveryCost, &payment.GoodsTotal, &payment.CustomFee); err != nil {
		if err == sql.ErrNoRows {
			return payment, fmt.Errorf("%s:%w", op, err)
		}
		return payment, fmt.Errorf("%s:%w", op, err)
	}
	return payment, nil
}

func (s *Storage) GetDeliverByID(id int64) (model.Deliver, error) {
	const op = "storage.postgres.GetDeliverByID"
	var deliver model.Deliver
	if err := s.db.QueryRow(`SELECT name, phone, zip, city, address, region, email FROM deliveries WHERE id = $1`, id).Scan(&deliver.Name, &deliver.Phone, &deliver.Zip, &deliver.City, &deliver.Address, &deliver.Region, &deliver.Email); err != nil {
		if err == sql.ErrNoRows {
			return deliver, fmt.Errorf("%s:%w", op, err)
		}
		return deliver, fmt.Errorf("%s:%w", op, err)
	}
	return deliver, nil
}

func (s *Storage) InsertPayments(payment model.Payment) (int64, error) {
	const op = "storage.postgres.insertPayments"
	var id int64
	row := s.db.QueryRow(
		`INSERT INTO payments (transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
					VALUES ($1, $2, $3,$4,$5,$6,$7,$8,$9,$10 ) RETURNING id;`,
		payment.Transaction, payment.RequestId, payment.Currency, payment.Provider, payment.Amount, payment.PaymentDt, payment.Bank, payment.DeliveryCost, payment.GoodsTotal, payment.CustomFee,
	)

	if err := row.Err(); err != nil {
		return 0, fmt.Errorf("%s:%w", op, err)
	}

	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("%s:%w", op, err)
	}

	return id, nil
}

func (s *Storage) InsertDeliver(deliver model.Deliver) (int64, error) {
	const op = "storage.postgres.InsertDeliver"
	var id int64
	row := s.db.QueryRow(
		`INSERT INTO deliveries (name, phone, zip, city, address, region, email)
					VALUES ($1, $2, $3,$4,$5,$6,$7) RETURNING id;`,
		deliver.Name, deliver.Phone, deliver.Zip, deliver.City, deliver.Address, deliver.Region, deliver.Email,
	)

	if err := row.Err(); err != nil {
		return 0, fmt.Errorf("%s:%w", op, err)
	}

	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("%s:%w", op, err)
	}

	return id, nil
}

func (s *Storage) InsertProducts(orderId int64, product []model.Product) ([]int, error) {
	const op = "storage.postgres.InsertDeliver"
	var ids []int
	var id int
	var row *sql.Row
	for _, v := range product {
		row = s.db.QueryRow(
			`INSERT INTO products (chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status, order_id)
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id;`,
			v.ChrtId, v.TrackNumber, v.Price, v.Rid, v.Name, v.Sale, v.Size, v.TotalPrice, v.NmId, v.Brand, v.Status, orderId,
		)

		if err := row.Err(); err != nil {
			return []int{}, fmt.Errorf("%s:%w", op, err)
		}

		if err := row.Scan(&id); err != nil {
			return []int{}, fmt.Errorf("%s:%w", op, err)
		}
		ids = append(ids, id)
	}

	return []int{}, nil
}

func (s *Storage) InsertOrder(order model.Order) (int64, error) {
	const op = "storage.postgres.insertOrder"
	var id int64

	payment, err := s.InsertPayments(order.Payment)

	if err != nil {
		return 0, fmt.Errorf("%s:%w", op, err)
	}

	deliveries, err := s.InsertDeliver(order.Deliver)

	if err != nil {
		return 0, fmt.Errorf("%s:%w", op, err)
	}

	row := s.db.QueryRow(
		`INSERT INTO orders (order_uid, track_number, entry,locale,internal_signature,customer_id,delivery_service,shardkey,sm_id,date_created,oof_shard,payments_id,delivery_id)
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id;`,
		order.OrderUid, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature, order.CustomerId, order.DeliveryService, order.Shardkey, order.SmId, order.DateCreated, order.OofShard, payment, deliveries,
	)

	if err := row.Err(); err != nil {
		// TODO: delete other table that order
		return 0, fmt.Errorf("%s:%w", op, err)
	}

	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("%s:%w", op, err)
	}

	_, _ = s.InsertProducts(id, order.Product)

	return id, nil
}

func (s *Storage) GetDeliverByPhone(phone string) (model.Deliver, error) {
	const op = "storage.postgres.GetDeliverByName"

	var deliver dbDeliverWithPriority
	err := s.db.QueryRowContext(context.Background(), "SELECT id, name, phone, zip, city,address,region,email FROM deliveries WHERE phone=$1", phone).Scan(&deliver.Name, &deliver.Phone, &deliver.Zip, &deliver.City, &deliver.Address, &deliver.Region, &deliver.Email)

	switch {
	case err == sql.ErrNoRows:
		return model.Deliver{}, fmt.Errorf("%s:%w", op, err)
	case err != nil:
		return model.Deliver{}, fmt.Errorf("%s:%w", op, err)
	}

	return model.Deliver{
		Name:    deliver.Name,
		Phone:   deliver.Phone,
		Zip:     deliver.Zip,
		City:    deliver.City,
		Address: deliver.Address,
		Region:  deliver.Region,
		Email:   deliver.Email,
	}, nil
}

type dbDeliverWithPriority struct {
	Name    string `db:"name"`
	Phone   string `db:"phone"`
	Zip     string `db:"zip"`
	City    string `db:"city"`
	Address string `db:"address"`
	Region  string `db:"region"`
	Email   string `db:"email"`
}
