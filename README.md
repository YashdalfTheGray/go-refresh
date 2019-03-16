# go-refresh

Refreshing golang knowledge by writing random golang

## building

Standard go project but has `go mod` support. So just need a `make` to get started.

## running

Running the whole binary just takes running `make run`.

## Cleaning

Just a simple `make clean` will clean all the built things.

## `make` Points of Interest

This section is a collection of `make` targets that are beyond the usual build, run, clean tasks.

### `image`

The `make image` target will use the base64 generated image that we get from the go binary (located in `bin/temp.base64`) and convert it into a `.png` file. We then use [`imgcat`](https://www.iterm2.com/documentation-images.html) to display it after we've made sure that `imgcat` actually exists on the system.
