FROM golang:1.22-alpine as builder
RUN apk --update add build-base curl

WORKDIR /src/todox
ADD go.mod .
RUN go mod download

ADD . .
RUN go build --ldflags '-linkmode=external -extldflags="-static"' -tags osusergo,netgo -buildvcs=false -o bin/app cmd/app/main.go

FROM alpine
RUN apk add --no-cache ca-certificates

WORKDIR /bin/

# Copying binaries
COPY --from=builder /src/todox/bin/app .

CMD /bin/app
