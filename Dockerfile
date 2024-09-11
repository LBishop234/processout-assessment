FROM golang:1.22.2 as builder
WORKDIR /api
COPY api /api
RUN CGO_ENABLED=1 go build -o api

FROM golang:1.22.2
WORKDIR /api
COPY --from=builder /api/api /api
EXPOSE 8080
CMD ["/api/api"]