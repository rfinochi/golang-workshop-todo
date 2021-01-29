FROM golang:latest
EXPOSE 82/tcp

COPY . /app
WORKDIR /app

RUN go test ./cmd/web -run TestCompleteAPIInMemory
RUN	go test ./cmd/web -run TestCompleteAPIInMemory
RUN go test ./cmd/web -run TestMain
RUN go test ./cmd/web -run TestSwagger
RUN go test ./cmd/web -run TestBadRequestError
RUN go test ./cmd/web -run TestNotFoundError
RUN go test ./cmd/web -run TestPageNotFoundError
RUN go test ./cmd/web -run TestRequestWithoutAPIToken
RUN go test ./cmd/web -run TestInternalServerError
RUN go test ./cmd/web -run TestValidations

ENV PORT 82 
ENV TODO_REPOSITORY_TYPE Mongo 
ENV TODO_MONGO_URI mongodb://db:27017
ENV TODO_API_TOKEN 85ba6be3-b2d5-4c15-aae5-d4878dfa203c 

ENTRYPOINT go run ./cmd/web
