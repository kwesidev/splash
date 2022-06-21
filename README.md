#Splash

![example workflow](https://github.com/kwesidev/splash/actions/workflows/go.yml/badge.svg)

Downloads Random Pictures from  https://unsplash.com/  and sets them as your Wallpaper every 30seconds , This works on GNOME .

## Tested on
* Ubuntu 18.04
* Ubuntu 20.04
* Ubuntu 22.04  


## Dependencies
* Go >=1.5.x


## Usage

```
  $ export GOPATH=~/yourgoworkspace
  $ go get -v github.com/kwesidev/splash
  $ cd $GOPATH/src/github.com/kwesidev/splash
  $ export GOBIN=~/yourgoworkspace/bin
  $ go install .
  $ cd $GOPATH/bin
  $ ./splash -interval 30s

```
