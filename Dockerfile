FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o scheduler.exe 

EXPOSE 8000
EXPOSE 8080

CMD ["./scheduler.exe"]