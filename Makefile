.PHONY:

mig-up:
	migrate -path ./migrations -database 'postgresql://postgres:22222@localhost:5435/postgres?sslmode=disable' up

mig-down:
	migrate -path ./migrations -database 'postgresql://postgres:22222@localhost:5435/postgres?sslmode=disable' down

build: 
	go build -o ./api cmd/main.go

run: build mig-up
	./api

swag: 
	swag init -g main.go --dir cmd,internal/handlers
	
swag-rm:
	sudo rm -rf ./docs