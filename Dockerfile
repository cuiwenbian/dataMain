FROM  golang:1.14.6-alpine AS builder
WORKDIR /alert
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux  go build -mod=vendor -o app .


FROM alpine:latest
#RUN apk --no-cache add ca-certificates
WORKDIR /app

COPY --from=builder /alert/conf.yaml .
COPY --from=builder /alert/app .

EXPOSE 8080

CMD ["./app"]