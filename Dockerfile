FROM alpine:latest AS alpine

# We need git to automatically get the changes of a deployment
RUN apk add -U --no-cache ca-certificates git

# See https://goreleaser.com/customization/dockers_v2/#how-it-works
ARG TARGETPLATFORM

WORKDIR /workdir
COPY $TARGETPLATFORM/pulse-event-cli /bin/
ENTRYPOINT ["pulse-event-cli"]
