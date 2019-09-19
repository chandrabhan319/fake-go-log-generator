FROM golang as builder

WORKDIR /app/

ADD . .

RUN go get -d -v ./...

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o log

FROM scratch

COPY --from=builder /app/log .

ENTRYPOINT ["./log"]
