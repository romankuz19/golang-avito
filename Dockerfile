FROM golang:1.21.0

RUN go version
ENV GOPATH=/

COPY ./ ./


# build go app
RUN go mod download
RUN go build -o avito-proj ./cmd/main.go

CMD ["./avito-proj"]