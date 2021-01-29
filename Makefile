TODO_API_TOKEN=85ba6be3-b2d5-4c15-aae5-d4878dfa203c
export TODO_API_TOKEN

run:
	go run ./cmd/web

test:
	go test ./cmd/web -run TestCompleteAPIInMemory
	go test ./cmd/web -run TestMain
	go test ./cmd/web -run TestSwagger
	go test ./cmd/web -run TestBadRequestError
	go test ./cmd/web -run TestNotFoundError
	go test ./cmd/web -run TestPageNotFoundError
	go test ./cmd/web -run TestRequestWithoutAPIToken
	go test ./cmd/web -run TestInternalServerError
	go test ./cmd/web -run TestValidations

build:
	go build ./cmd/web

docker:
	docker-compose up