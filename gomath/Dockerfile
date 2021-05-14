FROM golang:latest

WORKDIR /app

COPY . .

RUN ls -all
RUN go build -o math

CMD [ "./math" ]