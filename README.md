# File Smurf

### A Tiny Library for Batch File Modification in Go
<img src="http://vignette1.wikia.nocookie.net/smurfs/images/9/92/Handy_Comic_Book.jpg/revision/latest?cb=20120920121205">

## What?
Filesmurf is super basic - it will start at a given location in the filesystem, walk the tree and apply a specified `action` to any files that meet the `match` condition.

## Why?
I needed a way to conditionally rename files across my machine. Specifically, clean up my photo library by moving all images (files that had either _.jpg_ or _.cr2_ file extensions) into directories based on what date they were taken.

## Usage
```
filesmurf.Run(rootPath string, match MatchFunc, action ActionFunc)
```

Filesmurf defines 2 function types: `MatchFunc` and `ActionFunc`.
#### MatchFunc
```
type MatchFunc func(string) bool
```
A MatchFunc takes a file path as a `string` and returns a `bool` indicating if the conditions outlined in the function were met

#### ActionFunc
```
type ActionFunc func(string)
```
An ActionFunc takes a file path as a `string` and should modify the file in some way.

See my project [photosmurf](https://www.github.com/tomquirk/photo-smurf-go) for a real example.

## Contributions
Please do, I'm a Go noob :joy:

## Todo
* Ability to specify multipe src paths (goroutines to walk each?)
* Ability to flag an `ActionFunc` to be executed as a goroutine
