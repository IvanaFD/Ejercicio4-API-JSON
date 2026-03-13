FROM golang:1.25-alpine

WORKDIR /app

COPY . .

EXPOSE 24785

CMD ["go", "run", "cmd/server/main.go"]