# Golang setup and getting started
Following instructions are verfied setup of Golang on Ubuntu 16.04 version

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

## Setup Go development 
- Create directories for development
`mkdir -p $HOME/go/{bin,src,pkg}`
- Add path to environment permanently by adding in profile
```
echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.profile
echo "export GOPATH="\$HOME/go" >> ~/.profile
echo "export GOBIN="\$GOPATH/bin"" >> ~/.profile
source ~/.profile
```
Where, 

bin - directory to hold the go source executables

src - source directory in which all the source modules will be coded

pkg - directory to hold the dependency import packages


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
https://www.tecmint.com/install-go-in-linux/

https://golang.org/doc/install

https://github.com/golang/go/wiki/SettingGOPATH
