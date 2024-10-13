docker-run:
	docker compose up -d --build
migration-up:
	migrate --path db/migration --database "postgresql://go_finance:go_finance@localhost:5432/go_finance?sslmode=disable" -verbose up
migration-down:
	migrate --path db/migration --database "postgresql://go_finance:go_finance@localhost:5432/go_finance?sslmode=disable" -verbose down