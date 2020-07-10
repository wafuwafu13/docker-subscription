FROM golang:1.9

RUN mkdir /application
COPY main.go /application

CMD ["go", "run", "/application/main.go"]