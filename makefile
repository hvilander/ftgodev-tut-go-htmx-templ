run: build
	@./bin/ftgodev-tut

install:
	@go install github.com/a-h/templ/cmd/templ@latest
	@go get ./...
	@go mod vendor
	@go mod tidy 
	@go mod download
	@npm install -D tailwindcss
	@npm install -D daisyui@latest

css:
	## action up without npx
	npx tailwindcss -i view/css/app.css -o public/styles.css

build: 
	@templ generate view
	@go build -o bin/ftgodev-tut main.go


up: ## DB migration up
	@go run cmd/migrate/main.go up

drop:
	@go rum cmd/drop/main.go up

down: # db mirgration down
	@go run cmd/mirgrate/main.go down

reset: 
	@go run cmd/migrate/reset/main.go

migration: ## Migrations against the db
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

gen:
	@go run cmd/generate/main.go

seed:
	@go run cmd/seed/main.go
