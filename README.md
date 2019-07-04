<h1 align="center">MemeGen</h1>
<h2 align="center">

  [![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)[![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com)[![forthebadge](https://forthebadge.com/images/badges/uses-badges.svg)](https://forthebadge.com)

  ![Go Version](https://img.shields.io/badge/Go%20Version-latest-brightgreen.svg)
  [![Go Report Card](https://goreportcard.com/badge/github.com/Depado/memegen)](https://goreportcard.com/report/github.com/Depado/memegen)
  [![Build Status](https://drone.depa.do/api/badges/Depado/memegen/status.svg)](https://drone.depa.do/Depado/memegen)
  [![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/Depado/memegen/blob/master/LICENSE)
  [![Say Thanks!](https://img.shields.io/badge/Say%20Thanks-!-1EAEDB.svg)](https://saythanks.io/to/Depado)
</h2>

<h2 align="center">Very Simple Meme Generator</h2>

## Introduction

Memegen is a very simple CLI program that can write text on a picture you
provide, either on top, bottom or both. It supports font loading, font sizing,
will word-wrap properly if the text is too long. Just a little project for fun
mostly.

## Install

Grab the latest binary on the [releases page](https://github.com/Depado/memegen/releases)
or install it from source:

```sh
$ go get -u github.com/Depado/memegen
```

Note: This project uses go modules

## Usage

```
Usage:
  memegen <picture> [options] [flags]
  memegen [command]

Available Commands:
  help        Help about any command
  version     Show build and version

Flags:
  -b, --bottom string      text to display on bottom
      --font.outline int   outline of the font (default 4)
  -f, --font.path string   font path to use (default "fonts/impact.ttf")
  -s, --font.size float    font size (default 70)
  -h, --help               help for memegen
  -o, --output string      output file (default "out.png")
  -t, --top string         text to display on top

Use "memegen [command] --help" for more information about a command.
```

In order for this to work, you'll need to download fonts or specify the path
of the `.ttf` font you wish to use (that is installed on your system). For 
example you can create the `fonts/` directory and 
[download the `impact` font](https://www.wfonts.com/font/impact) in it. 