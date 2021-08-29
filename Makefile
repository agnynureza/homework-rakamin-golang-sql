lint:
	go fmt ./...

test:
	go test -v ./...

build:
	go build -v .

run:
	go run main.go

coverage-out:
	go tool cover -html=coverage.out

coverage:
	go test ./... -coverprofile coverage.out -count=1 -v
	
generate-mock:
	mockgen -source=./repository/movies.go -destination=./mocks/repository/movies_mock.go -package=mocks
	mockgen -source=./services/movies.go -destination=./mocks/services/movies_mock.go -package=mocks