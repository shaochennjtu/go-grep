# go-grep
Implement go-grep (grep functionality in Golang)

Usage

go-grep [-e] [pattern] [file]


# utils

Utils is a go language development toolkit that mainly includes functions such as string, slice, map, time, type judgment, type conversion, and data compression.

# How to use utils?
Public

  import "github.com/shaochennjtu/utils"

  utils.Empty(arg interface{}) bool
  utils.Export(arg interface{}) string
  utils.Json(arg interface{}) string
  utils.Decode(s string) error
  utils.Unmarshal(s string, args ...interface{}) interface{}
  
  utils.Type(arg interface{}) string
  utils.IsInt(arg interface{}) bool
  utils.IsInt64(arg interface{}) bool
  utils.IsFloat64(arg interface{}) bool
  utils.IsString(arg interface{}) bool
  utils.IsTime(arg interface{}) bool
  utils.IsBool(arg interface{}) bool
  utils.IsSlice(arg interface{}) bool
