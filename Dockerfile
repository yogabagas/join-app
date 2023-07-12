FROM golang:alpine

ENV GOOS=linux

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go mod vendor

RUN go build -o binary -ldflags "-X cmd/bootstrap.Flags=$FLAGS"

EXPOSE 9090

ENTRYPOINT ["/app/binary"]

CMD ["api-serve"]