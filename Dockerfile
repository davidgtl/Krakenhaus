FROM golang:alpine
ADD ./src /go/src/app
WORKDIR /go/src/app
ENV PORT=80
CMD ["go", "run", "main.go"]
