VersionCheck
============

Version check uses the very convenient go-version library from Hashicorp. It
provides a simple API for registering go library versions, and checking them against
a constraint.

All the logic for checking lives in the go-version library, versioncheck is just
a way to more easily integrate version dependency error detection into your application.

How to use it
=============

Download the library

```go
go get "github.com/StabbyCutyou/versioncheck"
```

Import the library

```go
import "github.com/StabbyCutyou/versioncheck"
```

Wherever your app initially boots, you could integration VersionCheck possibly by
defining an initialization section, where in you register your known dependencies,
and the version you intend to use. You could use a function like so:

```go
func checkRegistry(){
  versioncheck.Register("LibA",libA.VersionString, "= 1.4")
  versioncheck.Register("LibB", libB.VersionString, "<~ 2.5.2")
  versioncheck.Register("LibC", libC.VersionString, ">= 3.9")
  errs := versioncheck.Check()
  for err in range errs {
    log.Error(err) // or Fatal?
  }
}
```

You can decide if errors are an issue or not.

But Go libraries don't really have a convention for versioning!
===============================================================

I know, isn't that super annoying? Maybe one will arise. But a simple string
based checker seemed the most flexible option of all.

LICENSE
=========
Apache v2 - See LICENSE
