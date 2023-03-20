FROM golang:latest

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build 

EXPOSE 8080

CMD [ "./url-shortener-go" ]