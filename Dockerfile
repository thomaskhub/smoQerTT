FROM golang:1.21-alpine
WORKDIR /app
CMD [ "go", "run",  "." ]
