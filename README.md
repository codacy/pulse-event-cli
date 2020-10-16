# Pulse Events CLI

Command line interface to push events to the pulse service

## Usage

For a detailed list of all commands/flags use:

```
./event-cli --help
```

### Push single events

#### Deployments

```sh
./event-cli push deployment \
    --api-key "<API-KEY>" \
    --identifier 1.0.1 \
    --timestamp 1602852032 \
    d7c1baaa0975a0e3577dad1c4c2368d3dd4f33b5 6ccff2820ff356b609d8a000e082af866d144cc8
```

#### Changes

```sh
./event-cli push change \
    --api-key "<API-KEY>" \
    --identifier d7c1baaa0975a0e3577dad1c4c2368d3dd4f33b5 \
    --timestamp 1602852032
```

#### Incidents

```sh
./event-cli push incident \
    --api-key "<API-KEY>" \
    --identifier d7c1baaa0975a0e3577dad1c4c2368d3dd4f33b5 \
    --timestampCreated 1602852032 \
    --timestampResolved 1602852033
```

### Git helper

#### Deployments + changes

This will push the deployment and all commits from `<REF>` to the `HEAD` of the repo.

```sh
./event-cli push git deployment \
    --previous-deployment-ref="<REF>" \
    --api-key "<API-KEY>" \
    --identifier 1.0.1 \
    --timestamp `date +%s`
```

## Build

```sh
go build
```

## Release

```sh
goreleaser --rm-dist
```
