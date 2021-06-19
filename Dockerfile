FROM golang:alpine

EXPOSE 8080

WORKDIR /app

# download and cache dependencies
COPY go.mod go.sum ./
RUN go mod download -x

# copy source and build
COPY . .
RUN go build -o main .

# set environment variables for app here

CMD ["./main"]
