FROM golang:alpine

ENV GOOS=linux

ENV GOARCH=amd64

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go mod vendor

RUN go build -o binary -ldflags "-X cmd/bootstrap.Flags=$FLAGS"

EXPOSE 8800

ENTRYPOINT ["/app/binary"]

CMD ["api-serve"]