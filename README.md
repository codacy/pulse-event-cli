# Pulse Events CLI

Command line interface to push events to the pulse service

## Push events

### Deployments

```sh
./event-cli push deployment \
    --credentials "<CREDENTIALS>" \
    --identifier 1.0.1 \
    --timestamp 1602852032 \
    d7c1baaa0975a0e3577dad1c4c2368d3dd4f33b5 6ccff2820ff356b609d8a000e082af866d144cc8
```

### Changes

```sh
./event-cli push change \
    --credentials "<CREDENTIALS>" \
    --identifier d7c1baaa0975a0e3577dad1c4c2368d3dd4f33b5 \
    --timestamp 1602852032
```

### Incidents

```sh
./event-cli push incident \
    --credentials "<CREDENTIALS>" \
    --identifier d7c1baaa0975a0e3577dad1c4c2368d3dd4f33b5 \
    --timestampCreated 1602852032 \
    --timestampResolved 1602852033
```

## Build

```sh
go build                                                                          
```

## Release

```sh
goreleaser --rm-dist
```
