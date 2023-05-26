package migration

func CreateTableCustomer() string {
	return `CREATE TABLE IF NOT EXISTS customers (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`
}