FROM golang:latest

RUN apt install make

WORKDIR /tmp/build

COPY . .
RUN make docker
COPY ./output /app

WORKDIR /app

RUN rm -rf /tmp/build

EXPOSE 8050

CMD ["/app/go-avatar"]
