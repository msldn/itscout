FROM golang:1.17-alpine as builder

RUN addgroup -S itscout && adduser -S itscout -G itscout -h /opt/itscout 
WORKDIR /opt/itscout

COPY gosrc gosrc
COPY go.mod ./

WORKDIR /opt/itscout/gosrc

RUN go mod tidy
RUN go build -o itscout-server . 

FROM golang:1.17-alpine as runtime
WORKDIR /opt/itscout
COPY  --from=builder   /opt/itscout/gosrc/itscout-server  /opt/itscout/itscout-server
RUN chmod +x /opt/itscout/itscout-server


EXPOSE 8000

CMD [ "/opt/itscout/itscout-server" ] 