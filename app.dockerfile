FROM golang:alpine

RUN mkdir /learning

ADD . /learning

WORKDIR /learning

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT ["CompileDaemon", "--directory=/learning",  "--command=/learning/learning-api"]
