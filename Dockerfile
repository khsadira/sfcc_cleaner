#builder container
FROM golang:latest as builder
LABEL maintainer="Khan Sadirac <khan.sadirac42@gmail.com"
WORKDIR /app
ENV GO111Module=on
RUN	go mod init "github.com/khsadira/cleaner"
RUN	go mod vendor
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o sfcc_clean .

#main container
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app .
EXPOSE 9250
CMD ["./sfcc_clean"]
