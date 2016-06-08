#Splash
Downloads Random Pictures from  the Web and sets them as your Wallpaper every 30seconds by Default ,but you are free to change it.

##Tested on
* Ubuntu 14.04
* Ubuntu 16.04
* Fedora 24


##Prerequisite
* Go 1.5.x or Go 1.6.x

##Usage
```
  $ export GOPATH=~/yourgoworkspace
  $ go get -v github.com/kwesidev/splash
  $ cd $GOPATH/src/github.com/kwesidev/splash
  $ export GOBIN=~/yourgoworkspace/bin
  $ go install .
  $ cd $GOPATH/bin
  $ ./splash -interval 30s

```
