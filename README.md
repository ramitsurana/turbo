# Turbo 

Simple and Powerfull Utility for [Docker][1]

[![codebeat badge](https://codebeat.co/badges/e7fce2e3-69e8-451e-b9ba-de3d39b1da28)](https://codebeat.co/projects/github-com-ramitsurana-turbo)
[![Build Status](https://travis-ci.org/ramitsurana/turbo.svg?branch=master)](https://travis-ci.org/ramitsurana/turbo)
[![Build Status](https://semaphoreci.com/api/v1/ramitsurana/turbo/branches/master/badge.svg)](https://semaphoreci.com/ramitsurana/turbo)

![turbo1](https://cloud.githubusercontent.com/assets/8342133/16713587/95b469bc-46ca-11e6-8fb3-e56c7ce7d19d.png)

## Getting Started

![turbo](https://cloud.githubusercontent.com/assets/8342133/16805119/72fd724c-492c-11e6-9da1-6151a70df1d4.gif)

### Prerequisites

* Go v1.6
* Docker v1.11 and above

(Note that the code uses the relatively new Go vendoring, so building requires Go 1.6 or later, or you must export GO15VENDOREXPERIMENT=1 when building with Go 1.5.) 

### Principles:

We welcome every idea that the contributors put forward to.But to analyze the best amongst them, we consider two basic & important fundamental principles

* KISS(Keep it simple stupid)

The idea of the project [Turbo][2] is to make the lives of docker users more easy and time saving than before.If you have got an idea that you think is useful and simple to implement for [docker][1] users.Please submit us your idea by creating an issue !!

* Don't Reinvent the Wheel

Docker is an amazing effort towards making the containers easy to adapt.It has a variety of amazing and awesome features for the open source community.By building [Turbo][2],let's try to not build it into a subsitute of features that [Docker][1] already has !!


### Steps to follow

````
$ go get -v github.com/ramitsurana/turbo
$ cd $GOPATH/src/github.com/ramitsurana/turbo
$ ./turbo
````
Turbo:
  Simple and Powerfull utility for Docker

Usage:
  Turbo [command]

Available Commands:
  backup      backups all your docker stuff
  clean       Cleans up all your docker images
  destroy     Erases off all the exited containers
  kickstart   restarts all your containers quickly
  monitor     To monitor your containers
  replica     To create Replicas of your containers
  rkt         Installs and configures rkt
  search      Search images from multiple registries
  ship        Transfer your docker images over a remote i.p.
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

## Contributing

Contributions can be made easily by making PR's and opening issues on the github repo.Big Thank you to all the contributors !

## License

[Apache License 2.0](LICENSE)

[1]: http://docker.com
[2]: http://ramitsurana.github.io/turbo
