# gref
Gref is a utility library for data assignment using reflection in Golang. Main function: Assignment between two parameters, such as two different structs.


### Installation
```
go get github.com/zy1024/gref
```

### Document Language : **[English](README.md), [中文](README_CN.md).**

## Features

### 1. Copy Struct Fields
The `CopyStructFields` function copies fields from the source struct to the destination struct, commonly used to return specified parameters. It only copies fields from src to dst if dst is empty and has the same name as src.
Notice: This function does not handle nested structs and only processes struct pointers.

### Example
The sample code can be seen in the `TestCopyStructFields` method in `example_test.go`.

## Version
### 1.1.0
1.The name of the `CopyFields` method is changed to `CopyStructFields`. In addition, the support for the assignment of the same field name and different data types in two different structures is added. Includes array, struct, struct pointer, struct array pointer.

2.The newly added `CopyBasicValue` method is used to assign values between different underlying data types.


## Next development (if there is a need, you can mention it in the issue)
1. Increase the assignment between different types of slices.