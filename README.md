

![turbo1](https://cloud.githubusercontent.com/assets/8342133/16713587/95b469bc-46ca-11e6-8fb3-e56c7ce7d19d.png)

# Turbo

Running Docker containers at the speed of light 

Powerful CLI for Docker users

````
$ go get github.com/ramitsurana/turbo
$ cd $GOPATH/src/github.com/ramitsurana/turbo
$ go run main.go
$ go build github.com/ramitsurana/turbo
$ ./turbo
Welcome to Turbo
Optimize your docker enviorment with ease

Usage:
  Turbo [command]

Available Commands:
  backup      backups all your docker stuff
  clean-all   cleans up all your docker images
  gc          <W.I.P.> cleans up all the stopped containers
  kickstart   Restarts all your containers in a jiff
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

## Commands
------------

### Backup

Backups all your stuff so that you can have a copy in case anything goes wrong.

````
$ turbo backup
````

### Clean-all

Cleans all your docker images 

````
$ turbo clean-all
````

### Garbage Container

Kills all of your stopped containers.

````
$ turbo gc
````

### Kickstart 

Restarts all of your containers.

````
$ turbo kickstart
````
### Version

Displays info about version of turbo.

````
$ turbo version
````
## License

Apache License 2.0
