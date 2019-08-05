# goa
golang pinyin utils.

[![Build Status](https://travis-ci.org/lovego/pinyin.svg?branch=master)](https://travis-ci.org/lovego/pinyin)
[![Coverage Status](https://img.shields.io/coveralls/github/lovego/pinyin/master.svg)](https://coveralls.io/github/lovego/pinyin?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/lovego/pinyin?1)](https://goreportcard.com/report/github.com/lovego/pinyin)
[![GoDoc](https://godoc.org/github.com/lovego/pinyin?status.svg)](https://godoc.org/github.com/lovego/pinyin)

## usage
```go
  fmt.Println(Initials("长大"))
  fmt.Println(Initials("长城abc"))
  fmt.Println(Initials(" a长b大,c "))
  // Output:
  // ZD
  // ZCABC
  // AZBDC
```
