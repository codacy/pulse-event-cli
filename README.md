# Pulse Events CLI

Command line interface to push events to the pulse service

## Github Action

In order to send events about deployments and changes to pulse, we need to have a way to identify different deployments. A simple way to identify deployments is by using git-tags to store when a deployment has happened.

The following example uses Codacy's [git-version](https://github.com/codacy/git-version) Github-Action to
generate new versions on each deployment. The generated versions are then sent to Pulse together with
all the changes between the current and the previous deployment.

```yaml
name: Release

on:
  push:
    branches: ["master"]

jobs:
  release:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
        with:
          # Will fetch all history and tags required to generate version
          fetch-depth: 0

      - name: Git Version
        id: generate-version
        uses: codacy/git-version@2.4.0

      - name: "Tag version"
        run: |
          git tag ${{ steps.generate-version.outputs.version }}
          git push --tags "https://codacy:${{ secrets.GITHUB_TOKEN }}@github.com/codacy/pulse-event-cli"

      # ...
      # Deployment to live environment
      # ...

      # Send events to pulse
      - name: "Push data to pulse"
        uses: ./
        with:
          args: push git deployment \
            --api-key ${{ secrets.PULSE_API_TOKEN }} \
            --previous-deployment-ref ${{ steps.generate-version.previous-version }} \
            --identifier ${{ steps.generate-version.outputs.version }} \
            --timestamp "$(date +%s)"
```

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

## Build

```sh
go build
```

## Release

```sh
goreleaser --rm-dist
```
