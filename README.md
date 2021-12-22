# Turbo 

Simple and Powerfull Utility for [Docker][1]

Introductory blog post - http://ramitsurana.github.io/turbo

 [![codebeat badge](https://codebeat.co/badges/e7fce2e3-69e8-451e-b9ba-de3d39b1da28)](https://codebeat.co/projects/github-com-ramitsurana-turbo)
 [![GoDoc](https://godoc.org/github.com/ramitsurana/turbo?status.svg)](https://godoc.org/github.com/ramitsurana/turbo)
 [![Build Status](https://travis-ci.org/ramitsurana/turbo.svg?branch=master)](https://travis-ci.org/ramitsurana/turbo)
 [![Build Status](https://drone.io/github.com/ramitsurana/turbo/status.png)](https://drone.io/github.com/ramitsurana/turbo/latest)
 [![Build Status](https://semaphoreci.com/api/v1/ramitsurana/turbo/branches/master/badge.svg)](https://semaphoreci.com/ramitsurana/turbo)

![turbo1](https://cloud.githubusercontent.com/assets/8342133/16713587/95b469bc-46ca-11e6-8fb3-e56c7ce7d19d.png)

The snail is &copy; under [Deamworks Animation](http://www.dreamworksanimation.com/) movie [Turbo](https://en.wikipedia.org/wiki/Turbo_(film)).It is used for non-profit & non-commercial representation purposes only.

## What is Turbo ?

Turbo is a simple and easy to use utility that can be used over a docker ready enviorment.Its main purpose is to simplify the use of docker,in a simple and useful manner.It is a useful combination of a variety of different tools used to manage docker containers.

## Features

* Supports multiple 3rd party tools
* Built to be used as Binary file
* Contains several advanced features for [docker][1]


## Getting Started

![turbo](https://cloud.githubusercontent.com/assets/8342133/16805119/72fd724c-492c-11e6-9da1-6151a70df1d4.gif)

### Requirements

* [Go](http://golang.org) v1.6
* [Docker][1] v1.11 and above
* Linux 

**Tested for the above specifications.Results may vary accordingly.**

(Note that the code uses the relatively new Go vendoring, so building requires Go 1.6 or later, or you must export GO15VENDOREXPERIMENT=1 when building with Go 1.5.) 

### 3rd party tools

Turbo uses some 3rd party tools for giving the best perfomance to the [docker][1] users.Some of these are

* Glances
* [Minikube][4]
* [Logspout][5]
* [Harbor][6]
* [Minikube][7]

*Note: These are third party installtions and works only when you use there specific commands.They are not dependencies.*

### Getting Started

Checkout [releases](https://github.com/ramitsurana/turbo/releases) for proper installation instructions.

````
$ curl -L https://github.com/ramitsurana/turbo/archive/v0.2.tar.gz | tar zx
$ cd turbo-0.2 
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
  harbor      installs and configures vmware harbor
  kickstart   restarts all your containers quickly
  kube        Installs and configure minikube for your system
  ldap        Uses to install and configure Openldap
  log         Uses logspout to collect your docker logs
  monitor     To monitor your containers
  proxy       Installs and configure squid3 proxy for your system
  refresh     Completely removes and re-installs docker
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

*Not in Aphabetical Order*

* [Backup](#backup) - Backups all your stuff so that you can have a copy in case anything goes wrong.
* [Clean](#clean) - Wipes of all your [docker][1] images from your system.
* [Destroy](#destroy) - Kills all of your exited containers.
* [Kickstart](#kickstart) - Restarts all your containers
* [Log](#log) - Installs and configures logspout on your system
* [Monitor](#monitor)
* [Refresh](#refresh) - Removes Docker ecosystem and Installs a new one
* [Replica](#replica)
* [Rkt](#rkt) - Installs and configures rkt on your system
* [Search](#search) - Searches Docker registry hub,quay.io for your images
* [Ship](#ship) -  Ships off all your Docker Stuff to remote <I.P>
* [Version](#version) - Displays Turbo version  
* [kube](#kube) - Installs and Configures Minikube on your system
* [Harbor](#harbor) - Installs and configures Harbor on your system
* [Proxy](#proxy) - Installs and configures squid proxy on your system
* [Ldap](#ldap) - Installs and configures Openldap on your system


**More to come**

## Contributing

Contributions can be made easily by making PR's and opening issues on the github repo.
Thank you to all the [contributors][3] for your awesome contributions !

## License

[Apache License 2.0](LICENSE)

[1]: http://docker.com
[2]: http://ramitsurana.github.io/turbo
[3]: https://github.com/ramitsurana/turbo/graphs/contributors
[4]: http://github.com/kubernetes/minikube
[5]: https://github.com/gliderlabs/logspout
[6]: https://github.com/vmware/harbor
[7]: https://github.com/kubernetes/minikube
