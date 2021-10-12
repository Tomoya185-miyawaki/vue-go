FROM golang:1.17

WORKDIR /api
COPY . .
ENV GO111MODULE=on

RUN go get github.com/pilu/fresh
CMD ["fresh"]
RUN go get