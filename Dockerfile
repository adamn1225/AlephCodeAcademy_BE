FROM golang:1.23

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

# Still download what's already resolved
RUN go mod download

COPY . ./

# THIS LINE FIXES THE BUILD ERROR
RUN go mod tidy
RUN apt-get update && apt-get install -y postgresql-client

RUN go build -o main .

CMD ["./main"]
