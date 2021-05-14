FROM golang:1.15.12

WORKDIR /app

COPY gomath/ .

RUN go build -o math

CMD [ "./math" ]