FROM golang:1.22-alpine as builder
RUN apk --update add build-base curl

WORKDIR /src/todox
ADD go.mod .
RUN go mod download

ADD . .
RUN go run ./cmd/build

FROM alpine
RUN apk add --no-cache ca-certificates

WORKDIR /bin/

# Copying binaries
COPY --from=builder /src/todox/bin/app .

CMD /bin/app
