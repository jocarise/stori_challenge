FROM golang:1.23

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build

RUN mkdir -p /app/uploads
RUN chown -R nobody:nogroup /app/uploads

CMD ["./newsletter-service"]