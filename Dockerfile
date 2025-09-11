FROM alpine:latest AS alpine

# We need git to automatically get the changes of a deployment
RUN apk add -U --no-cache ca-certificates git

WORKDIR /workdir
COPY pulse-event-cli /bin/
ENTRYPOINT ["pulse-event-cli"]
