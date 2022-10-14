FROM golang:1.16-alpine
WORKDIR /app
COPY ./server ./
RUN go build -o /server/main
EXPOSE 8080
CMD ["/server/main"]