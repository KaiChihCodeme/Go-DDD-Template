FROM golang:1.22

WORKDIR /app

COPY . .

RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest

ENV PATH="/go/bin:${PATH}"

RUN swag init -g cmd/main.go -o ./docs/swagger/docs

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]