package migration

func CreateTableOrder() string {
	return `CREATE TABLE IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY,
		customer_id INT NOT NULL,
		product VARCHAR(100) NOT NULL,
		quantity INT NOT NULL,
		price INT NOT NULL,
		created_at TIMESTAMPTZ DEFAULT NOW(),
		updated_at TIMESTAMPTZ DEFAULT NOW(),
		FOREIGN KEY (customer_id) REFERENCES customers (id) ON DELETE CASCADE
	)`
}