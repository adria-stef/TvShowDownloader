# TvShowDownloader [:crown:](http://adria-stef.github.io/TvShowDownloader/)

[![Build Status](https://travis-ci.org/adria-stef/TvShowDownloader.svg?branch=master)](https://travis-ci.org/adria-stef/TvShowDownloader)
[![Coverage Status](https://coveralls.io/repos/adria-stef/TvShowDownloader/badge.svg?branch=master&service=github)](https://coveralls.io/github/adria-stef/TvShowDownloader?branch=master)
[![Go Report Card](http://goreportcard.com/badge/adria-stef/TvShowDownloader)](http://goreportcard.com/report/adria-stef/TvShowDownloader)
[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/adria-stef/TvShowDownloader)
[![MIT License](http://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/adria-stef/TvShowDownloader/blob/master/LICENSE)

:tv: TvShowDownloader is a tool you can use to automatically download all your favorite Tv Shows.

:confetti_ball: All you need to do is configure a path you want you files to appear in and make a list of all your shows.

## Prerequsites

1. You need an Internet connection.
1. You need to download and install the `ctorrent` tool.

## Get the Project

Run `go get github.com/adria-stef/TvShowDownloader`

## How to use?

### Configure the application
Populate the `files/list.yml` file. Add a downloading path and a list of all tv shows.

An example file would like this:
```bash
---
download_path: some/path/to/file
list:
  - New Girl
  - Stuck in the Middle
  - NCIS New Orleans
  - Colony
  - Portlandia
  - Impractical Jokers
```

### Run it!

Simply run the following command in the root of the repository:
```bash
go run main.go
```

If you want to specify the exact period of time that the checks are performed you can use the `minutes` parameter:
```bash
go run main.go minutes <number>
```
