FROM golang:1.17

WORKDIR /todolist_backend
COPY ./ /todolist_backend

RUN go get ./...


RUN go build -o /docker-gs-ping

EXPOSE 8000

CMD [ "go", "run", "main.go" ]