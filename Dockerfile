FROM golang:1.20 as builder

RUN go install github.com/ServiceWeaver/weaver/cmd/weaver@latest

COPY . /app

WORKDIR /app

RUN go mod tidy
RUN weaver generate .
RUN go build .

FROM golang:1.20

RUN go install github.com/ServiceWeaver/weaver/cmd/weaver@latest

COPY --from=builder /app/simple-auth-sw /app/simple-auth-sw

WORKDIR /app