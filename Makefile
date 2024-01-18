include .env

dev:
	@echo Starting the server...
	@go run main.go
db:
	@echo Running sql files to create a database
	@psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) -d postgres \
		-f ./database/migrations/init.sql
	@psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) -d jobtask \
		-f ./database/migrations/create.sql
	@psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) -d jobtask \
		-f ./database/migrations/insert.sql
	@echo Done...
	
