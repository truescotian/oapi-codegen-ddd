FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN cd internal/petstore && CGO_ENABLED=0 GOOS=linux go build -o /oapi-codegen-example

EXPOSE 3000

CMD ["/oapi-codegen-example"]

