FROM alpine:latest AS os
WORKDIR /app
COPY . /app
CMD ["apk", "add", "bash"]

FROM golang:1.26.3-alpine3.23 AS go
WORKDIR /app
COPY --from=os /app /app
CMD ["go", "run", "."]

