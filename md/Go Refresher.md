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

---

*mpl-go* by NewForester.
Notes licensed under a Creative Commons Attribution-ShareAlike 4.0 International Licence.

<!-- EOF -->
