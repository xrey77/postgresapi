FROM golang:latest

LABEL maintainer: Reynald Marquez-Gragasin
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
ENV PORT 8080
RUN go build
EXPOSE 8090
CMD ["./postgresapi"]



# FROM golang:1.14.6-alpine3.12 as builder
# COPY go.mod go.sum /go/src/github.com/xrey77/postgresapi/
# WORKDIR /go/src/github.com/xrey77/postgresapi
# RUN go mod download
# COPY . /go/src/github.com/xrey77/postgresapi
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/postgresapi github.com/xrey77/postgresapi

# FROM alpine
# RUN apk add --no-cache ca-certificates && update-ca-certificates
# COPY --from=builder /go/src/github.com/xrey77/postgresapi/build/postgresapi /usr/bin/postgresapi
# EXPOSE 8080 8080
# ENTRYPOINT ["/usr/bin/postgresapi"]