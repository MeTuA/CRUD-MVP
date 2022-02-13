FROM golang:1.16-alpine AS builder
WORKDIR /source
COPY . /source
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -o crud-mvp ./cmd/.

FROM alpine:3.9
RUN mkdir /app
WORKDIR /app
COPY --from=builder /source/crud-mvp /usr/local/bin
ENTRYPOINT [ "crud-mvp" ]
