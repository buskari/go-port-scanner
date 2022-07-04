FROM golang

WORKDIR /usr/port-scanner
RUN go install github.com/cosmtrek/air@latest

COPY ./ ./

CMD air

