# Turbo 

Powerful CLI tool for [Docker][1]

[![codebeat badge](https://codebeat.co/badges/e7fce2e3-69e8-451e-b9ba-de3d39b1da28)](https://codebeat.co/projects/github-com-ramitsurana-turbo)
[![Build Status](https://travis-ci.org/ramitsurana/turbo.svg?branch=master)](https://travis-ci.org/ramitsurana/turbo)
[![Twitter URL](https://img.shields.io/twitter/url/http/shields.io.svg?style=social&maxAge=2592000?style=flat-square)](https://twitter.com/ramitsurana)
[![Twitter Follow](https://img.shields.io/twitter/follow/shields_io.svg?style=social&label=Follow&maxAge=2592000?style=flat-square)](https://twitter.com/ramitsurana)

![turbo1](https://cloud.githubusercontent.com/assets/8342133/16713587/95b469bc-46ca-11e6-8fb3-e56c7ce7d19d.png)

## Getting Started

### Prerequisites

* Go v1.6
* Docker v1.11 and above

### Steps to follow

````
$ go get -v github.com/ramitsurana/turbo
$ cd $GOPATH/src/github.com/ramitsurana/turbo
$ ./turbo
Turbo:
  Powerful CLI Tool for Docker

Usage:
  Turbo [command]

Available Commands:
  backup      backups all your docker stuff
  clean       cleans up all your docker images
  clear       wipes off all the stopped containers
  kickstart   <W.I.P.>Restarts all your containers quickly
  monitor     To monitor your containers
  replica     replicates your containers
  search      Search images from multiple registries
  ship        ships off all your docker images
  version     prints the current version number of turbo

Flags:
      --config string   config file (default is $HOME/.turbo.yaml)
  -h, --help            help for Turbo
  -t, --toggle          Help message for toggle

Use "Turbo [command] --help" for more information about a command.
````

Adding it to your path :

````
$ alias turbo="sudo '${GOPATH}/src/github.com/ramitsurana/turbo/turbo'"
````
## About Turbo

[Turbo][2] is a cool new way to use [docker][1] in a fun and easy way.It uses [docker][1] commands in the backend.Built using Golang.This project is dedicated to docker users who would like to use [docker][1] in a more effective and faster way.

## Commands

[Backup](#backup) | [Clean](#clean) | [Version](#version) | [Clean](#clean) | [Kickstart](kickstart)

### Backup

Backups all your stuff so that you can have a copy in case anything goes wrong.

````
$turbo backup
````

### Clean

Wipes of all your docker images from your system.

````
$turbo clean
````

### Clear

Kills all of your stopped containers.

````
$turbo clear
````

### Kickstart 

Restarts all of your containers.

````
$turbo kickstart
````
### Version

Displays info about version of turbo.

````
$turbo version
````
## Demo

![turbo](https://cloud.githubusercontent.com/assets/8342133/16805119/72fd724c-492c-11e6-9da1-6151a70df1d4.gif)

## Contributing

Contributions can be made easily by making PR's and opening issues for submitting your ideas.Big Thank you to all the contributors !

## License

[Apache License 2.0](LICENSE)

[1]: http://docker.com
[2]: http://ramitsurana.github.io/turbo
