FROM golang:1.16 as base

FROM base as dev

RUN apt update && apt upgrade && \
    apt install openssh

WORKDIR /opt/app/api

COPY go.mod go sum ./

RUN go mod download

COPY . /opt/app/api

RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./main"]