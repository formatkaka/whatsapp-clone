FROM golang:1.14.2
WORKDIR /home/mcstk/trell/go-starter/
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-starter .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /home/mcstk/trell/go-starter/go-starter .
ENV APP_ENV=production
EXPOSE 80
CMD ["./go-starter"]
