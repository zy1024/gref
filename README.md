# gref
gref is a utility library in Golang that utilizes reflection for data assignment. The main function is to assign values between two parameters, such as two different structs.

### Installation
```
go get github.com/zy1024/gref
```

### Document Language : **[English](README.md), [中文](README_CN.md).**

## Features

### 1. Copy
The `Copy(src,dst)` function requires two pointers, such as pointers to structs or slices. It will use reflection to assign the value of src to dst. The assignment of structs is based on matching field names, and only non-empty fields in dst will be assigned.
### Example
The sample code can be seen in the `TestCopy` method in `example_test.go`.

## Version
### 1.3.0
Official release, supporting assignment operations for structs, pointers, slices, and basic data types.

### Next Development Steps
1. Add support for other types, such as channels.