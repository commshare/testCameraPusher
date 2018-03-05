# gitea
it's forked gitea, only contains log module

## How to use?
```
/*
Simple Document For Log:
mode: "console" or "file"
log levels:
    0 - TRACE
    1 - DEBUG
    2 - INFO
    3 - WARN
    4 - ERROR
    5 - CRITICAL
    6 - FATAL
filename: log file path
*/

package main

import "github.com/djwackey/gitea/log"

func main() {
    mode := "console"

    config := `{"level":0,"filename":"test.log"}`
    log.NewLogger(0, mode, config)

    log.Info("Hello World")
}

## https://github.com/golang
- https://github.com/golang/sync