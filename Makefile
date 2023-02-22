swag: install-swag
	swag init -g cmd/api/main.go

install-swag:
	go install github.com/swaggo/swag/cmd/swag@v1.8.10