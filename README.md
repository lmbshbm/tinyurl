# TinyURL

[![Build Status](https://travis-ci.org/adolphlwq/tinyurl.svg?branch=master)](https://travis-ci.org/adolphlwq/tinyurl)  [![Go Report Card](https://goreportcard.com/badge/github.com/adolphlwq/tinyurl)](https://goreportcard.com/report/github.com/adolphlwq/tinyurl)  [![GoDoc](https://godoc.org/github.com/adolphlwq/tinyurl?status.svg)](https://godoc.org/github.com/adolphlwq/tinyurl)

<p align="center">
  <a href="http://tinyurl.adolphlwq.xyz" target="_blank">
    <img src="tinyurl.png" width="700px">
    <br>
    Live Demo
  </a>
</p>

<p align="center">a url shorten web service written by Golang, Vue and Gin.</p>

## Requisites
- Golang(1.8+)
- [Govendor](https://github.com/kardianos/govendor)
- MySQL
- make

## Quick Start
1. clone project to **GOPATH**
```
git clone https://github.com/adolphlwq/tinyurl.git $GOPATH/src/github.com/adolphlwq/tinyurl
```
2. sync golang packages
```
govendor sync
```
3. build binary
```
make
```
4. run binary
```
./tinyurl -dbname tinyurl -user user -pass pass -dbport 2306
```

## TODOs
- [X] validate input url format
- [X] improve random generate string algorithm
    - [X] use math/rand.Read instead math/rand.Intn func
- [X] use logrus replace golang log lib
- [X] reserch [wrk](https://github.com/wg/wrk)
- [X] add test case
- [ ] dynamic adjust short path length (default is 4)
- [ ] count each url parse time (high concurrent situation)
- [ ] qrcode support
- [ ] list api
- [ ] retry logic with init db and request

## Reference
- [GitHub/Ourls](https://github.com/takashiki/Ourls)
- [GitHub/uriuni](https://github.com/dchest/uniuri)
