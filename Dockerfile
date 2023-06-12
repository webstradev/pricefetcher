FROM golang:1.20.0-alpine as builder

WORKDIR /app

COPY . . 

RUN C_GO_ENABLED=0 go build -o /app/bin/pricefetcher .

FROM alpine

COPY --from=builder /app/bin/pricefetcher /bin/pricefetcher

EXPOSE 3000 

CMD ["./bin/pricefetcher"]