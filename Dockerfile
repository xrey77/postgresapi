FROM golang:1.14.6-alpine3.12 as builder
COPY go.mod go.sum /go/src/gitlab.com/xrey77/postgresapi/
WORKDIR /go/src/gitlab.com/xrey77/postgresapi
RUN go mod download
COPY . /go/src/gitlab.com/xrey77/postgresapi
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/postgresapi gitlab.com/xrey77/postgresapi

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/gitlab.com/xrey77/postgresapi/build/postgresapi /usr/bin/postgresapi
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/postgresapi"]