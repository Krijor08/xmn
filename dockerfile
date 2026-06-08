FROM golang:1.26.3-alpine3.23 AS go
WORKDIR /app
COPY . /app
CMD ["go", "run", "."]

