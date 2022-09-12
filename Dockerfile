FROM golang

ADD . /go/src/github.com/hunterlong/tokenexporter
WORKDIR /go/src/github.com/hunterlong/tokenexporter
RUN go mod tidy
RUN go install

ENV GETH https://mainnet.infura.io/Nsy8W84s3hqW4eE49ljZ
ENV PORT 9021
ENV ADDRESSES ${ADDRESSES}
ENV TOKEN_LIST ${TOKEN_LIST}
ENV PORT ${PORT}

RUN mkdir /app
WORKDIR /app
#ADD addresses.txt /app
#ADD tokens-list.json /app

EXPOSE ${PORT}

CMD ["/bin/bash", "-c", "echo $ADDRESSES > /app/addresses.txt && echo $TOKEN_LIST > /app/token-list.json && /go/bin/tokenexporter"]
