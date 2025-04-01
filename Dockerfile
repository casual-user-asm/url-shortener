FROM golang:1.23-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN go build -o shortener .
EXPOSE 8080
CMD [ "./shortener" ]