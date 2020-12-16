FROM alpine:latest as alpine

RUN apk add -U --no-cache ca-certificates

FROM scratch
WORKDIR /workdir
COPY pulse-event-cli /
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/pulse-event-cli"]
