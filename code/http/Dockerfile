#build stage
FROM golang:alpine AS builder
RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk -U --no-cache add git
WORKDIR /go/src/app
COPY . .
RUN go build -o /go/bin/app -v code/http/main.go

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
COPY --from=builder /go/src/app/code/http/templates /templates
WORKDIR /
ENTRYPOINT /app