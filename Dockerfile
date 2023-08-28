FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o avito ./cmd/app/main.go

CMD [ "./avito" ]