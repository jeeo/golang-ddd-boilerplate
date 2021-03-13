FROM golang:buster as builder
WORKDIR /bin
COPY . .
RUN go mod download
RUN go get github.com/google/wire/cmd/wire
RUN CGO_ENABLED=1 cd ./cmd/some-app && wire
RUN CGO_ENABLED=0 GOOS=linux go build -o app github.com/Jeeo/golang-ddd-boilerplate/cmd/some-app

FROM alpine

COPY --from=builder /bin/app .
CMD [ "./app" ]  