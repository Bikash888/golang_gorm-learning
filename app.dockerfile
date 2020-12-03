FROM golang:alpine

RUN mkdir /locator

ADD . /locator

WORKDIR /locator

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT ["CompileDaemon", "--directory=/locator",  "--command=/locator/locator-api"]
