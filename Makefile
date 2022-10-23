export REPO_GOSRC_PATH := $(shell pwd)

lint:
    golangci-lint run --enable-all

test:
echo `pwd`
ls -l | awk '{ print $9 }' | sed 1d | while read dir
do
cd $dir
	go get -d
	go test
cd ..
done


build:
echo `pwd`
ls -l | awk '{ print $9 }' | sed 1d | while read dir
do
cd $dir
	go get -d
	go build
cd ..
done

all: build test
 
# run:
 
clean:
    go clean
