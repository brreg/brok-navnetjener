FROM golang:1.20 as builder

WORKDIR /usr/local/go/src/brok/navnetjener/

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY api ./api
COPY database ./database
COPY model ./model
COPY utils ./utils
COPY server.go .

RUN ls -R
RUN go build -v -o /navnetjener

FROM registry.access.redhat.com/ubi9-minimal

COPY --from=builder /navnetjener /navnetjener

CMD ["/navnetjener"]