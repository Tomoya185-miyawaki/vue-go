FROM golang:1.17

WORKDIR /api
COPY . .
ENV GO111MODULE=on

RUN go get github.com/pilu/fresh \
  && go get -u github.com/go-sql-driver/mysql \
  && go get github.com/rubenv/sql-migrate/...
CMD ["fresh"]
RUN go get