FROM golang:1.17-alpine as base
RUN addgroup -S itscout -g 1001 && adduser -S itscout -G itscout -u 1001 -h /opt/itscout -s /bin/sh
USER itscout

FROM base as builder
WORKDIR /opt/itscout
COPY --chown=1001:1001 . .
RUN go mod tidy && go build -o itscout-server . 

FROM base as runtime
WORKDIR /opt/itscout
COPY  --from=builder  --chown=1001:1001 /opt/itscout/itscout-server  /opt/itscout/itscout-server
EXPOSE 8000
CMD [ "/opt/itscout/itscout-server" ] 