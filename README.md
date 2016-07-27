# Turbo 

Simple and Powerfull Utility for [Docker][1]

[![codebeat badge](https://codebeat.co/badges/e7fce2e3-69e8-451e-b9ba-de3d39b1da28)](https://codebeat.co/projects/github-com-ramitsurana-turbo)
[![Build Status](https://travis-ci.org/ramitsurana/turbo.svg?branch=master)](https://travis-ci.org/ramitsurana/turbo)
[![Build Status](https://semaphoreci.com/api/v1/ramitsurana/turbo/branches/master/badge.svg)](https://semaphoreci.com/ramitsurana/turbo)

![turbo1](https://cloud.githubusercontent.com/assets/8342133/16713587/95b469bc-46ca-11e6-8fb3-e56c7ce7d19d.png)

## What is Turbo ?

Turbo is a simple and easy to use utility that can be used over a docker ready enviorment.Its main purpose is to simplify the use of docker,in a simple and useful manner.It is a useful combination of a variety of different tools used to manage docker containers.

## Features

* Supports multiple 3rd party tools
* Built to be used as Binary file
* Contains several advanced features for [docker][1]


## Getting Started

![turbo](https://cloud.githubusercontent.com/assets/8342133/16805119/72fd724c-492c-11e6-9da1-6151a70df1d4.gif)

### Requirements

* Go v1.6
* [Docker][1] v1.11 and above
* Linux 

**Tested for the above specifications.Results may vary accordingly.**

(Note that the code uses the relatively new Go vendoring, so building requires Go 1.6 or later, or you must export GO15VENDOREXPERIMENT=1 when building with Go 1.5.) 

### 3rd party tools

Turbo uses some 3rd party tools for giving the best perfomance to the [docker][1] users.Some of these are

* Glances
* [Minikube][4]

### Steps to follow

````
$ curl -Lo https://github.com/ramitsurana/turbo/archive/v0.1.tar.gz 
$ tar xvf turbov0.1.tar.gz
$ cd turbov0.1 
$ chmod +x turbo
$ sudo mv turbo /usr/local/bin
$ turbo
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

## Commands

* [Backup](#backup)
* [Clean](#clean)
* [Destroy](#destroy)
* [Kickstart](#kickstart)
* [Ship](#ship)
* [Version](#version)


**More to come**

### Backup

Backups all your stuff so that you can have a copy in case anything goes wrong.

````
$ turbo backup
````

### Clean

Wipes of all your [docker][1] images from your system.

````
$ turbo clean
````

### Destroy

Kills all of your exited containers.

````
$ turbo destroy
````

### Kickstart 

Restarts all of your containers.

````
$ turbo kickstart
````

### Ship

Transfers all your Docker images to a remote i.p.

````
$ turbo ship
````

### Version

Displays info about version of turbo.

````
$ turbo version
````

## Contributing

Contributions can be made easily by making PR's and opening issues on the github repo.Big Thank you to all the [contributors][3] !

## License

[Apache License 2.0](LICENSE)

[1]: http://docker.com
[2]: http://ramitsurana.github.io/turbo
[3]: https://github.com/ramitsurana/turbo/graphs/contributors
[4]: http://github.com/kubernetes/minikube
