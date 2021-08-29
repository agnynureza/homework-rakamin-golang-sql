lint:
	go fmt ./...

test:
	go test -v ./...

build:
	go build -v .

run:
	go run main.go

generate-mock:
	mockgen -source=./modules/closing/service.go -destination=./modules/closing/mocks/service_mocks.go -package=mocks