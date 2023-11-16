# Use the golang:latest base image
FROM golang:1.21.4-alpine3.17 as build 
WORKDIR /app

COPY . .

RUN go build -o /go/bin/app

FROM gcr.io/distroless/static-debian11
COPY --from=build /go/bin/app /

CMD ["/app"]

