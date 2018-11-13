FROM golang:latest AS build
WORKDIR /application
COPY . .
ARG GOOS=linux
ARG CGO_ENABLED=0
RUN go build -a -o bin/httpd cmd/httpd/main.go


FROM scratch
ENV GIN_MODE=release
COPY --from=build /application/bin/httpd /
CMD ["/httpd"]