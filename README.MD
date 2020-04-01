# Gromaudio usb dongle filenames fixer 

![Build](https://github.com/abramov-ks/gromaudio-flash-fixer/workflows/Build/badge.svg?branch=master)

## Description

util for fixind filenames on USB dongle for GROMAUDIO USB3 dongle.
This removes dumn files from macos and renames audio into `0001.mp3` format

## Installation

compile by `go build`

## Usage
`./gromfixer -volume path -action clean|fix`
- `PATH` - path to mounted usb
- `ACTION` - `clean` to clean from trash files or `fix` to clean from trash files and fix filenames
