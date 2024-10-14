# gref
gref 是一个用于在 Go 中进行反射操作的实用工具库。它提供了两个主要功能：生成更新字段和复制结构体字段。


### 安装
```
go get github.com/zy1024/gref
```

### 文档语言 : **[English](README.md), [中文](README_CN.md).**

## 功能

### 1. 复制结构体字段
`CopyStructFields` 函数将字段从源结构体复制到目标结构体，常用于返回指定参数。它只复制dst为空且与src具有相同名称和类型的字段。
Notice: 该函数不会处理嵌套结构体,并且只处理结构体指针。

### 示例
示例代码可以查看 `example_test.go` 中的 `TestCopyStructFields` 方法。

### Version 1.1 
1.`CopyFields` 方法名修改为 `CopyStructFields` 同时新增对两个不同结构体下同一字段名，不同数据类型的赋值的支持。包括数组、结构体、结构体指针、结构体数组指针。

2.新增加 `CopyBasicValue` 方法用于不同基础数据类型间的赋值。