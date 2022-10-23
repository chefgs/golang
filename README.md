[![Build Status](https://travis-ci.org/chefgs/golang.svg?branch=master)](https://travis-ci.org/chefgs/golang)

# Golang setup and getting started
Following instructions are verfied setup of Golang on Ubuntu 16.04 version

## Table of Contents
- [Install Golang](#install-golang)
- [Setup Go development environment](#setup-go-development-environment)
- [Verify the installation](#verify-the-installation)
- [Go to src path and start development](#go-to-src-path-and-start-development)
- [Bibliography](#bibliography)


## Install Golang
- Get the OS specific installer from [here](https://golang.org/dl/) 
- Run the command as `root` user or `sudo`
- Extract go executable in `/usr/local` directory
- For Ubuntu 16.04 version
```
cd /usr/local/
wget https://dl.google.com/go/go1.14.1.linux-amd64.tar.gz
tar -xvzf go1.14.1.linux-amd64.tar.gz 
rm go1.14.1.linux-amd64.tar.gz 
```
or
```
cd /tmp
wget https://dl.google.com/go/go1.14.1.linux-amd64.tar.gz
sudo tar -xvzf go1.14.1.linux-amd64.tar.gz -C /usr/local
```
-

## Setup Go development environment 
- Create directories for development

`mkdir -p $HOME/go/{bin,src,pkg}`

Where,

bin - directory to hold the go source executables

src - source directory in which all the source modules will be coded

pkg - directory to hold the dependency import packages

- Add path to environment permanently by adding in profile
```
export "PATH=$PATH:/usr/local/go/bin"
export "GOPATH=$HOME/go"
export "GOBIN=$GOPATH/bin"
```
- Execute .profile file to make effect of Env variable updates
```
source ~/.profile
```

## Verify the installation
```
go help
go env
```
## Go to src path and start development
```
cd $HOME/go/src
mkdir <new_go_module_name>
```
## Bibliography
[Install Go in Linux](https://www.tecmint.com/install-go-in-linux/)

[Golang Install Doc](https://golang.org/doc/install)

[Set Go Env Variable](https://github.com/golang/go/wiki/SettingGOPATH)

---

## If you're planning to contribute - Please read further 
- Read [Contribution Guidelines](./CONTRIBUTING.md) before contributing to this repository!

- Read GitHub [Code Of Conduct guidelines](./CODE_OF_CONDUCT.md)before contributing to this repository!

- Feel free to pick up any of the open [Issues](https://github.com/chefgs/golang/issues) or create new issue using the [template](https://github.com/chefgs/golang/issues/new/choose)!
