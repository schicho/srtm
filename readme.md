# SRTM Library

## Introduction

The Shuttle Radar Topography Mission (SRTM) was a NASA mission to provide digital elevation data for the entire world.
The data is free and available for use by anyone.
A good place to find data is on this [website](http://viewfinderpanoramas.org/dem3.html).


This library provides simple functions to read and convert SRTM data.
Both the SRTM1 and SRTM3 data are supported.

## Usage

image.go provides simply functions to convert the SRTM data files to Go's [image](https://pkg.go.dev/image) implementation,
which can be processed further.

## Commands

The cmd subdirectory contains simple scripts to convert the SRTM data to gray scale images directly.

## Docs

The docs subdirectory contains NASA's SRTM documentation, which can currently be found on the Web Archive as the webpage does not exist anymore.