up:
	docker-compose -f docker-compose.yaml up --build

down:
	docker-compose -f docker-compose.yaml down

generate-mocks:
	mockgen -source=repository/interface.go -destination=repository/mocks/mocks.go -package mocks

test:
	go test -v -coverprofile=coverage.out ./... && go tool cover -html=coverage.out -o coverage.html

swagger:
	swag init