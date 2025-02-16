FROM golang:1.23.2

WORKDIR /app

EXPOSE ${PORT}

CMD ["go", "run", "cmd/name/main.go"]