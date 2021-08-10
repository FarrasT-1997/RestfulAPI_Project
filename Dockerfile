From golang:1.16-alpine

WORKDIR /app

COPY ./go.mod ./
COPY ./go.sum ./

RUN go mod download

COPY ./main.go ./
COPY ./config ./config
COPY ./constant ./constant
COPY ./controller ./controller
COPY ./lib ./lib
COPY ./middlewares ./middlewares
COPY ./models ./models
COPY ./routes ./routes

RUN go build -o program

CMD ./program