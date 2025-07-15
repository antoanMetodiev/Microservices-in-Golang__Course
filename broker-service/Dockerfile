FROM golang

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . /app

EXPOSE 80

CMD [ "go", "run", "./cmd/api"]