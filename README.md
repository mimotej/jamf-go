<h2 align="center"> GO-jamf </h2>

<h3 align="center"> Simple REST application written in Golang which is used as an example application for CI/CD stack </h3>


<p align="center">
  <a href="https://github.com/mimotej/jamf-go/actions/workflows/release.yaml">
    <img alt="Build Status" src="https://img.shields.io/github/actions/workflow/status/mimotej/jamf-go/release.yaml">
  </a>
  <a href="https://github.com/mimotej/jamf-go">
    <img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/mimotej/jamf-go">
  </a>
<a href="https://github.com/mimotej/jamf-go/tags">
    <img alt="Latest GitHub tag" src="https://img.shields.io/github/v/tag/mimotej/jamf-go">
  </a>
<a href="https://github.com/mimotej/jamf-go/blob/master/LICENSE">
    <img alt="Latest GitHub tag" src="https://img.shields.io/github/license/mimotej/jamf-go">
  </a>
</p>

---


## :desktop_computer: How to run locally

Just run in the root of the application:

```shell
go run .
```

And application will start on port 8080 - http://localhost:8080 .

## :wrench: How to build

Run in the root:

```shell
go build -o .
```

Resulting binary is called `jamf-go`.

## :whale: Docker setup

This application uses docker for deployment you can build application running following command:

```shell
docker build -t go-jamf .
```

Then to run it use:

```shell
docker run -p 8080:8080 go-jamf:latest
```

Application is also being build as a part of CI pipeline in pull requests. Resulting images then can be found in [Dockerhub](https://hub.docker.com/repository/docker/mimotej/go-jamf/general). To deploy them to dev environment use `/deploy` in the PR, which will then open another pull request with this image in https://github.com/mimotej/jamf-manifests . After merging of this PR you will get new version of application in your dev environment. To then release new version just open New release issue and let workflow do its job :) .

## :keyboard: Development setup

To make development easier, this application uses [pre-commit](https://pre-commit.com/), which can be installed by running:

```shell
pip install pre-commit
pre-commit install
```
If you have any issues installing pre-commit feel free to consult [manual](https://pre-commit.com/index.html#install).


## Important sources

This application release process was inspired by release process of [peribolos as a service](https://github.com/operate-first/peribolos-as-a-service). Main inspiration was to utilize semantic release and to communicate current state of pipeline to user using issue comments, which makes it interactive and easier to understand.