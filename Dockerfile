FROM golang as builder

ADD . /go/src/github.com/praveen4g0/url-shortner/

WORKDIR /go/src/github.com/praveen4g0/url-shortner/

RUN go mod tidy && \
    go mod vendor

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o url-shortner .

FROM scratch

ENV PORT 8080

COPY --from=builder /go/src/github.com/praveen4g0/url-shortner/url-shortner /app/

WORKDIR /app

CMD ["./url-shortner"]