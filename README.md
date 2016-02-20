# TvShowDownloader
Project for GO course

## Prerequsites

You need an Internet connection

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

Simply run the following command in the root of the repository :
```bash
go run main.go
```

## Copyright and license
Code released under [the MIT license](https://github.com/adria-stef/TvShowDownloader/blob/master/LICENSE).
