FROM golang:1.9

RUN mkdir /application
COPY main.go /application

RUN go get github.com/AlecAivazis/survey

CMD ["go", "run", "/application/main.go"]