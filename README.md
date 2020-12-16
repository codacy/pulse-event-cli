# Pulse Events CLI

Pulse is a command-line interface to push events to <https://pulse.codacy.com/>.

Take a look <https://docs.pulse.codacy.com/>

## Github Action

You can use this Github-Action to send changes and deployment events to Pulse's service
directly from you CI.

The following workflow is an example where we use [git-version](https://github.com/codacy/git-version) to
generate new versions on each deployment and store that information in git tags.

```yaml
name: Pulse

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

      # Generate previous and next version from git tags
      - name: Git Version
        id: generate-version
        uses: codacy/git-version@2.4.0

      # Push git tag to repository
      - name: "Tag version"
        run: |
          git tag ${{ steps.generate-version.outputs.version }}
          git push --tags "https://codacy:${{ secrets.GITHUB_TOKEN }}@github.com/codacy/pulse-event-cli"

      # ...
      # Deployment steps
      #...

      # Push deployment and changes events to pulse
      - name: "Push data to pulse"
        uses: ./
        with:
          args: push git deployment \
            --api-key ${{ secrets.PULSE_ORG_PULSE_API_KEY }} \
            --system $GITHUB_REPOSITORY \
            --previous-deployment-ref ${{ steps.generate-version.outputs.previous-version }}\
            --identifier ${{ steps.generate-version.outputs.version }}\
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
