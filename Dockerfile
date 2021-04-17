FROM golang:1.16 as build

WORKDIR /pkg 

COPY ./src .

RUN go build -o renderder

FROM debian:latest

COPY --from=build /pkg/renderder /bin/renderer

CMD ["/bin/renderer"]
