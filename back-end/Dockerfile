FROM golang:1.22.6-alpine as build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o main .

CMD ["go", "run", "main.go"]

FROM golang:1.22.6-alpine as production
COPY --from=build /app /
EXPOSE 2727
CMD ["/main"]
