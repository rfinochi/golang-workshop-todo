FROM golang:latest
EXPOSE 82/tcp

COPY . /app
WORKDIR /app

RUN go test ./cmd/web -run TestCompleteAPIInMemory

ENV PORT 82 
ENV TODO_REPOSITORY_TYPE Mongo 
ENV TODO_MONGO_URI mongodb://db:27017

ENTRYPOINT go run ./cmd/web
