<!-- mpl-go by NewForester:  programming notes on and examples in Go -->

# Go Refresher Notes

Anyone can remember things they use everyday but it is not everyday one reinstalls a programming language.
No matter how well one knows a programming language, some details fade if one spends any length of time using another.

Below is not a Go tutorial but quick refresher notes.
Hopefully just enough to bring back all you do know that has slipped just beyond the instant recall threshold.

Do not forget to update these notes with what you have learnt since they were written.

## Overview

[Go](https://en.wikipedia.org/wiki/Golang)
is modern programming language developed by Google that first appeared towards the end of 2009.

Go is a compiled language that follows the structured / imperative paradigm.
The language is strongly and statically typed but features type inference and structural typing.
Memory management is based on garbage collection.

Go features concurrency:  the implementation of true coroutines is perhaps its defining characteristic.

The current stable release is 1.10.1 (released March 28 2018).

Go, its compiler and tools are free and open source and available under a modified BSD licences.

Go seeks to improve on C/C++ with a feature set that emphasises:

  * small (enough to keep in one's head as in C)
  * syntactic conciseness (more common in dynamic languages);
  * fast compilation times;
  * statically linked, standalone binaries.

Go deliberately omits featuers common in other languages leading to criticisms.
The use of garbage collection introduces pauses that limit the language's
use for systems programming and for hard real-time applications.

Go is not a class-based object oriented language but provides 'interfaces' instead that are similar to 'prototypes' in Java or
abstract base classes in C++.

Go has a package system that includes remote management and online documentation but
it is not as complete as those of Rust or Node.js since there is no central repository.

Go comes with language tools that encourage a consistent approach to many tasks, including code formatting.

The official web-site is [https://golang.org/](golang.org).

## Installation, Upgrade, Removal

### Installation

Go is available from Debian repositories but is inevitably an out of date version.
Best to install the latest stable (binary) version from the Go project.

#### Officially

For *nix systems, download the appropriate tar-ball, then unpack:

```bash
    $ sudo tar -C /usr/local -xzf go1.10.1.linux-amd64.tar.gz
```

You are left to alter your PATH to include _/usr/local/go/bin_.

The tar-ball is the order of 110 Mb and the final install 350 Mb.

#### Alternatively

If you are installing language tools on another volume:

```bash
    $ mkdir /media/work/lang/go;
    $ sudo ln -s /media/work/lang/go/ /usr/local;

    $ tar -C /usr/local -xzf go1.10.1.linux-amd64.tar.gz
```

#### Binaries

You do not need to alter your PATH if _/usr/local/bin_ is already on it:

```bash
    $ sudo ln -s /usr/local/go/bin/* /usr/local/bin;
```

There are only three binaries.

### Upgrade

The instructions state to delete your current installation and install the upgrade as for a new installation.

### Removal

It a simple as deleteing the top directory:

```bash
    rm -r /usr/local/go
```

and/or:

```bash
    rm -r /media/work/lang/go
```

Do not forget to tidy your PATH.

---

*mpl-go* by NewForester.
Notes licensed under a Creative Commons Attribution-ShareAlike 4.0 International Licence.

<!-- EOF -->
