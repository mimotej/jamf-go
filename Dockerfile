FROM golang:1.23.2 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o .

FROM scratch
COPY --from=builder /app/jamf-go /
EXPOSE 8080
ENTRYPOINT ["/jamf-go"]
