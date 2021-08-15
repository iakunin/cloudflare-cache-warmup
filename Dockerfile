# build
FROM golang:1.16-alpine AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /out/main .

# prod
FROM alpine:3.14
RUN apk add --no-cache ca-certificates
COPY --from=build /out/main /
CMD ["/main"]
