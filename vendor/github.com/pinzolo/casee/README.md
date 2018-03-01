# casee

[![Build Status](https://travis-ci.org/pinzolo/casee.png)](http://travis-ci.org/pinzolo/casee)
[![Coverage Status](https://coveralls.io/repos/github/pinzolo/casee/badge.svg?branch=master)](https://coveralls.io/github/pinzolo/casee?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/pinzolo/casee)](https://goreportcard.com/report/github.com/pinzolo/casee)
[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/pinzolo/casee)
[![license](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/pinzolo/casee/master/LICENSE)

Golang liibrary for case convertion of string.

## Usage

1. `go get -d github.com/pinzolo/casee`.
2. Add `github.com/pinzolo/casee` to `import` section in your go file.

## Functions

### Convert functions

* `ToSnakeCase`  
  Convert to snake_case style string.
* `ToChainCase`  
  Convert to chain-case style string.
* `ToCamelCase`  
  Convert to camelCase style string.
* `ToPascalCase`  
  Convert to PascalCase style string.

### Check functions

* `IsSnakeCase`  
  Check argument string is snake_case.
* `IsChainCase`  
  Check argument string is chain-case.
* `IsCamelCase`  
  Check argument string is camelCase.
* `IsPascalCase`  
  Check argument string is PascalCase.

If first character is digit, `IsCamelCase` and `IsPascalCase` always returns false.  
Because cannot judge camelCase or PascalCase.

