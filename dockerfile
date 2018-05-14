FROM golang:latest

WORKDIR .
COPY . .

RUN make

CMD ["app"]
EXPOSE 8080