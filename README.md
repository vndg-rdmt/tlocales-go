# tlocales

Localization package for Golang

## Install

```sh
go get -u github.com/vndg-rdmt/tlocales-go
```

### Overview

Packages exposes two main entities as an api - `Loader` and `Locales`.

Package constructor is used to create `Loader`, which uses a user-defined driver for loading with locales data,
so you can implement for yourself if needed, for example to load files from object storages, api and etc. Available drivers
can be found in the drivers directory. Currently out-of-box packages has a default fs driver to load files from filesystem.

> If modules traversing somehow does not works, just load drivers explicitly.

```sh
go get -u github.com/vndg-rdmt/tlocales-go/drivers
```

### Package

How to implement localization using this package

You create files or objects with named following pattern `name.language.extension`, here name - name of a dictionary, language - language code for this dictionary, extension - file extension for unmarshaller to parse its data.

```go
import tlocales github.com/vndg-rdmt/tlocales-go
```

```go

import driver_fs github.com/vndg-rdmt/tlocales-go/drivers/fs

loader := tlocales.New()
var err *tlocales.LocalesError = loader.Load("storage-key",
	tlocales.UseDriver(driver_fs.New("/etc/my-locales")),
	tlocales.RegisterUnmarshaller("json", json.Unmarshal),
)
```

You create new loader, which is able to load dictionaries more than just one time. Then call `Load` method, to load locales
by key, and then you pass `tlocales.Option`-s. Here it you define driver, which is a fs driver, register unmarshaller for parsing json encoded data, you can also specify and type and extension, just specify implement required interface.

Package uses custom errors types, so you have a generic error without maps for arguments.

What is a key here? It's a key for driver, it's fully used by driver and loader does not a clue what this key is all about, it's just passes it to your driver instance. For example for fs driver it's a directory path, where all locales files is stored.

For example, you have such directory

```
test/fs/
├── errors.ch.json
├── errors.en.json
├── something.ch.json
├── something.en.json
└── something.ru.json
```

load directory with locales

```go
dr.Load("test/fs",
	tlocales.UseDriver(New("./")),
	tlocales.RegisterUnmarshaller("json", json.Unmarshal),
)
```

Use locales

```go
somethingLocales, _ := dr.GetLocales("something")
errorsLocales, _ := dr.GetLocales("errors")

if msg, ok := somethingLocales.Say("en", "key1"); !ok {
    fmt.Printlng("key1 message for such language not defined")

} else {
    fmt.Println(msg) // value1_en
}
```

> Templating currently in dev, and shurly will be added

### Drivers

Implement interface from sdk packages. You can use drivers/fs driver as a reference.

```go
import tlocales_sdk "github.com/vndg-rdmt/tlocales-go/sdk"
```

### Docs

Will be added to release 1.0.0