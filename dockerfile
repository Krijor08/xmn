FROM golang:1.26.3-alpine3.23
WORKDIR /app
COPY . /app
CMD ["go", "run", "."]
