# top95

Package top95 tests strings against a popular wordlist.

## Usage

```go
if top95.Include("hunter2") {
  log.Fatal("https://www.xkcd.com/936/")
}
```

## Benchmarks

Blazing fast!

```
$ go test -bench .
goos: linux
goarch: amd64
pkg: github.com/pnelson/top95
BenchmarkInclude-4   	 3000000	       420 ns/op
PASS
ok  	github.com/pnelson/top95	1.748s
```

## Licenses

Package top95 and the command line application are licensed under the terms
described in [LICENSE](https://github.com/pnelson/top95/blob/master/LICENSE).

Function `Include` depends on data generated from a subset of
[berzerk0/Probable-Wordlists](https://github.com/berzerk0/Probable-Wordlists)
and is licensed under the terms described in
[LICENSE](https://github.com/pnelson/top95/blob/master/LICENSE.top95).
