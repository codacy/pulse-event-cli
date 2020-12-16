# Pulse Events CLI

Command line interface to push events to the pulse service

## Usage

For a detailed list of all commands/flags use:

```
./pulse-event-cli --help
```

### Push single events

#### Deployments

```sh
./pulse-event-cli push deployment \
    --api-key "<API-KEY>" \
    --identifier 1.0.1 \
    --timestamp 1602852032 \
    d7c1baaa0975a0e3577dad1c4c2368d3dd4f33b5 6ccff2820ff356b609d8a000e082af866d144cc8
```

#### Changes

```sh
./pulse-event-cli push change \
    --api-key "<API-KEY>" \
    --identifier d7c1baaa0975a0e3577dad1c4c2368d3dd4f33b5 \
    --timestamp 1602852032
```

#### Incidents

```sh
./pulse-event-cli push incident \
    --api-key "<API-KEY>" \
    --identifier d7c1baaa0975a0e3577dad1c4c2368d3dd4f33b5 \
    --timestampCreated 1602852032 \
    --timestampResolved 1602852033
```

### Git helper

#### Deployments + changes

This will push the deployment and all commits from `<REF>` to the `HEAD` of the repo.

```sh
./pulse-event-cli push git deployment \
    --previous-deployment-ref="<REF>" \
    --api-key "<API-KEY>" \
    --identifier 1.0.1 \
    --timestamp `date +%s`
```

### Detected environments

Currently the CLI detects CI environments where data is being sent from.
We're collecting this information to improve the support on how we collect the metrics.

If your CI environment is not in this list,
contact <mailto:pulsesupport@codacy.com> to let us know about it.

Supported environments:

- appveyor
- aws-codebuild
- azure-pipelines
- bitrise
- buddy
- buildkite
- circleci
- codefresh
- codemagic
- codeship
- docker
- drone
- github-actions
- gitlab-ci
- gocd
- google-cloud-build
- greenhouse
- heroku-ci
- jenkins
- jfrog-pipelines
- magnum
- semaphore
- shippable
- solano
- teamcity
- travis
- wercker

## Build

```sh
go build
```

## Release

```sh
goreleaser --rm-dist
```
