FROM golang:latest
EXPOSE 82/tcp

COPY . /app
WORKDIR /app

RUN go test -run TestCompleteApiInMemory

ENV PORT 82 
ENV REPOSITORY_TYPE Mongo 
ENV MONGO_URI mongodb://db:27017

ENTRYPOINT go run .