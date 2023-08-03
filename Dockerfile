FROM golang:1.20

WORKDIR /usr/local/go/src/brok/navnetjener/

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY api ./api
COPY database ./database
COPY model ./model
COPY server.go .

RUN ls
RUN go build -v -o /navnetjener

ENV DOCKER=true

CMD ["/navnetjener"]