# log

Alternative log writer for Go.

## Summary

* Writter in pure Go
* Require Go version > 1.8
* Async write log to files
* Can write log to stdin and stderr
* Support formatted records
* Automatical log rotate
* Support log levels
* MIT license

## Install

```
go get github.com/belfinor/go
```

## Usage

### Open/Close log

```go
package main

import (
  "github.com/belfinor/log"
)

func main() {
  log.Open("name=app path=. period=day level=info save=15")
  defer log.Close() // or log.Finish("app finish")

  log.Info("log record")
}
```

*period* can take one of the following values: *hour*, *day* or *month*. The *save* parameter specifies the number of log files to save. Logs are saved in the *path* directory and has names (app-YYYYMMDD.log if period is day, app-YYYYMM.log if period is month and app-YYYYMMDDHH.log if period is hour).

When a new period comes, the old log file will be closed and a new one will be created.

*log.Finish* closes log writer and save all messages to log file.

You can use multiple log writers in your application. Add params *global=0* and *log.Open* return custom log object.

```go
package main

import (
  "github.com/belfinor/log"
)

func main() {
  log1, err1 := log.Open("name=app1 path=. period=day level=info save=15 global=0")
  if err1 != nil {
    panic(err1)
  }
  defer log1.Close() // or log1.Finish("app finish")

  log2, err2 := log.Open("name=app2 path=. period=day level=info save=15 global=0")
  if err2 != nil {
    panic(err2)
  }
  defer log2.Close() // or log2.Finish("app finish")

  log1.Infof("log record %d %s", 1, "log1")
  log2.Errorf("log record %d %s", 1, "log2")
}
```

Add parameter *stderr=1* or *stdout=1* if you want print log message to *stderr* or *stdout*.

### Log levels

Log object support log levels (trace/debug/info/warn/error/fatal/none). Log writes nothing if selected level is *none*.
All *trace* and *debug* messages will be skipped if you select the logging level *info*.

You can set log level in *Open* (like *level=debug*) call or change it in runtime via *SetLevel*. *GetLevel* return current level.

```go
func main(){
  ...

  log.SetLevel("debug")
  fmt.Println(log.GetLevel()) // debug

  ...
}
```

On each level we have calls Trace/Debug/Info/Warn/Error/Fatal.

Moreover, *Fatal* writes message to log file and kills application.

### Formatted messages

Log object has functions Tracef/Debugf/Infof/Warnf/Errorf/Fatalf like Trace/Debug/Info/Warn/Error/Fatal, but they take format string and params list and print message to log file. The format options are same of fmt package.

```go
log.Infof("Hello, %s!", "Mike")
```
