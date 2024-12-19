FROM golang:1.23-alpine as builder

WORKDIR /app

# Copy the go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o secret-satana .

FROM alpine:latest

# Install dependencies for SQLite (if needed)
# RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/secret-satana .

COPY templates templates

EXPOSE 8080

ENV SATANA_SELECTED=false
ENV RESELECT_SATANA=false

CMD ["./secret-satana"]
