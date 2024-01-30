//go:build linux && cgo && !static_build

package main

// #cgo CFLAGS: -std=gnu11 -Wvla -Werror -fvisibility=hidden -Winit-self
// #cgo CFLAGS: -Wformat=2 -Wshadow -Wendif-labels -fasynchronous-unwind-tables
// #cgo CFLAGS: -pipe --param=ssp-buffer-size=4 -g -Wunused
// #cgo CFLAGS: -Werror=implicit-function-declaration
// #cgo CFLAGS: -Werror=return-type -Wendif-labels -Werror=overflow
// #cgo CFLAGS: -Wnested-externs -fexceptions
// #cgo CFLAGS: -I include
// #cgo pkg-config: lxc
// #cgo pkg-config: libcap
// #cgo LDFLAGS: -l:libintl.a
import "C"
