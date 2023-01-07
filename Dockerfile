FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /cart_service

EXPOSE 8081

CMD [ "/cart_service" ]