FROM golang:alpine

WORKDIR .
COPY . .

RUN make

CMD ["app"]
EXPOSE 8080