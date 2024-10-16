# gref
gref 是一个在 Golang 中利用反射进行数据赋值的实用工具库。主要功能：两个参数之间赋值，例如两个不同的结构体。


## 安装
```
go get github.com/zy1024/gref
```

## 文档语言 : **[English](README.md), [中文](README_CN.md).**

# 功能

## 1. Copy
`Copy(src,dst)` 函数需传入两个指针，例如两个指向结构体的指针、指向切片的指针。将通过反射将 src 的值赋予 dst。其中结构体的赋值是根据字段名相同来进行判断，且仅赋值 dst 的非空字段。

### 支持类型
支持基础数据类型、结构体、指针、切片。

### 运行效率
对比传统的json.Marshal和json.Unmarshal方式效率更高。相同的数据下，效率约为 3 - 4 倍。

### 示例
示例代码可以查看 `example_test.go` 中的 `TestCopy` 方法。

## 版本
### 1.3.0 
正式版，已支持 结构体、指针、切片以及基础数据类型的赋值操作。

## 下一步开发(有需求可以提到issue中)
1.增加对其余类型的支持，例如 chan、map.
