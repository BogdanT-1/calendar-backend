# build stage
FROM golang:1.18-alpine3.15 AS builder

# destination directory of the files
WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# first . means the root directory of our app and Second is the path that is copied to
COPY . .
RUN go build -v -o main main.go
# main is the output binary files and main.go is entry point for the app

# run stage, use alpine to get a smaller image from the golang image
FROM alpine:3.15
WORKDIR /app
COPY --from=builder /app/main .

CMD ["app/main"]