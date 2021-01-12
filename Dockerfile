FROM alpine:latest as alpine

RUN apk add -U --no-cache ca-certificates

# Build a statically linked git binary
ENV GIT_VERSION "v2.30.0"
WORKDIR /tmp/git
RUN apk add -U --no-cache git alpine-sdk autoconf automake zlib-dev zlib-static zlib
RUN \
  git clone --depth=1 git://github.com/git/git -b $GIT_VERSION /tmp/git && \
  make configure && \
  ./configure prefix=/tmp CFLAGS="${CFLAGS} -static" && \
  make install NO_TCLTK="YesPlease"

FROM scratch
WORKDIR /workdir
COPY pulse-event-cli /bin/
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=alpine /tmp/bin/git /bin/git
ENTRYPOINT ["pulse-event-cli"]
