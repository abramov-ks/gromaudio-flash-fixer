# Gromaudio usb dongle filenames fixer 

## Description

util for fixind filenames on USB dongle for GROMAUDIO USB3 dongle.
This removes dumn files from macos and renames audio into `0001.mp3` format

## Installation

compile by `go build`

## Usage
./gromfixer -volume <PATH> -action <ACTION>
PATH - path to mounted usb
ACTION - `clean` to clean from trash files or `fix` to clean from trash files and fix filenames
