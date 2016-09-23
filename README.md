# pidfile

This package reads and writes pid's from and to a file. It is primarily used to help with creating systemd services.

## Installation

Use the `go` tool:

`$ go get github.com/neoplatonist/pidfile`

## Usage

To read a file use:

`pidfile.ReadPid("path/to/file")`

To write to a file use:

`pidfile.WritePid(os.Getpid(), "path/to/file")`
