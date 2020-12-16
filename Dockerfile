FROM scratch
COPY pulse-event-cli /
ENTRYPOINT ["/pulse-event-cli"]
