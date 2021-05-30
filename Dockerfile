FROM golang:1.16
WORKDIR /app
COPY . .

RUN go get -v ./...

COPY . .

EXPOSE 8000

COPY scripts/*.sh /scripts/

RUN chmod a+x /scripts/*.sh
