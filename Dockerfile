FROM golang:latest

RUN apt install make
RUN mkdir /tmp/build
RUN mkdir /app
WORKDIR /tmp/build
COPY . .
RUN make
RUN go mod tidy 
RUN go build -o output/go-avatar
COPY output/* /app

WORKDIR /app

RUN rm -rf /tmp/build

EXPOSE 8050

CMD ["/app/go-avatar"]
