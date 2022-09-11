FROM golang

ADD . /go/src/github.com/hunterlong/tokenexporter
WORKDIR /go/src/github.com/hunterlong/tokenexporter
RUN go mod tidy
RUN go install

ENV GETH https://mainnet.infura.io/Nsy8W84s3hqW4eE49ljZ
ENV PORT 9021

RUN mkdir /app
WORKDIR /app
ADD addresses.txt /app
ADD tokens-list.json /app

EXPOSE 9021

ENTRYPOINT /go/bin/tokenexporter
