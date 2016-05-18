# Levenshtein Distance in Golang

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/ferhatelmas/levenshtein)
[![Build Status](https://travis-ci.org/ferhatelmas/levenshtein.png?branch=master)](https://travis-ci.org/ferhatelmas/levenshtein)

> Calculate levenshtein distance in Golang.

## Install

By go tool: `go get github.com/ferhatelmas/levenshtein`

## Usage

This uses default calculator which has cost of 1 for additions, deletions and substitutions.

```go
import github.com/ferhatelmas/levenshtein

levenshtein.Dist("aaa", "ab") // 2
```

You can specify different weights to increment/deletion and substitutions.

```go
levenshtein.New(1, 1).Dist("aaa", "ab") // 2
levenshtein.New(1, 2).Dist("aaa", "ab") // 3
levenshtein.New(1, 3).Dist("aaa", "ab") // 3
levenshtein.New(1, 4).Dist("aaa", "ab") // 3
levenshtein.New(2, 2).Dist("aaa", "ab") // 4
levenshtein.New(3, 2).Dist("aaa", "ab") // 5
```

## LICENSE

MIT © [Ferhat Elmas](https://github.com/ferhatelmas)
