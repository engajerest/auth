FROM golang:1.15.6
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o server .
EXPOSE 3300
CMD ["/app/server"]