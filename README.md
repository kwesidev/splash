#Splash

[![Build Status](https://travis-ci.org/kwesidev/splash.svg?branch=master)](https://travis-ci.org/kwesidev/splash)

Downloads Random Pictures from  the Web and sets them as your Wallpaper every 30seconds by Default ,but you are free to change it.

##Tested on
* Ubuntu 14.04 Unity
* Ubuntu 16.04 Unity


##Dependencies
* Go >=1.5.x


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
