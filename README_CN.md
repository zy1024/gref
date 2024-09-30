# gref
gref 是一个用于在 Go 中进行反射操作的实用工具库。它提供了两个主要功能：生成更新字段和复制结构体字段。


### 安装
```
go get github.com/zy1024/gref
```

### 文档语言 : **[English](README.md), [中文](README_CN.md).**

## 功能

### 1. 生成更新字段
`GenerateUpdateFields` 函数根据提供的请求对象生成一个更新字段的映射。
它只包括非零值且可导出的字段（即以大写字母开头的字段），并将字段名称从驼峰命名法转换为蛇形命名法，将减少更新数据库数据时所需的代码量。
Notice: 该函数不会处理嵌套结构体,并且只处理结构体指针。
#### 示例
```go
// before dataBase { name: mark , age: 30 , address: ""  , email: "" , password: "1234567890" }
req := &User{ Age: 31, Address: "New York", Email: "mark@example.com" , password: "" }

updateFields, err := gref.GenerateUpdateFields(req)

fmt.Println(updateFields) // 输出: map[age:31]

gorm.Updates(updateFields)

// after dataBase { name: mark , age: 31 , address: "New York" , email: "alice@example.com" , password: "1234567890" }
```

### 2. 复制结构体字段
`CopyFields` 函数将字段从源结构体复制到目标结构体，常用于返回指定参数。它只复制dst为空且与src具有相同名称和类型的字段。
Notice: 该函数不会处理嵌套结构体,并且只处理结构体指针。

### 示例
```go
// before src  { name: mark , age: 31 , address: "New York" , email: "alice@example.com" , password: "1234567890"...other fields }
// before dst  { name: "" , age: 0 , address: "" , email: "" ... other fields }

err := gref.CopyFields(&src, &dst)

// after dst  { name: mark , age: 31 , address: "New York" , email: "alice@example.com" ... other fields }
```