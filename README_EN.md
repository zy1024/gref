# gref

gref is a utility library for performing reflection operations in Go. It provides two main functionalities: generating update fields and copying struct fields.

### Installation
```
go get github.com/zy1024/gref
```

## Features

### 1. Generate Update Fields
The `GenerateUpdateFields` function generates a map of fields to update based on the provided request object. It only includes non-zero and exported fields (i.e., fields starting with an uppercase letter) and converts the field names from camelCase to snake_case, reducing the amount of code required to update database data.
Notice: This function does not handle nested structs and only processes struct pointers.

#### Example
```go
// before dataBase { name: mark , age: 30 , address: ""  , email: "" , password: "1234567890" }
req := &User{ Age: 31, Address: "New York", Email: "mark@example.com" , password: "" }

updateFields, err := gref.GenerateUpdateFields(req)

fmt.Println(updateFields) // 输出: map[age:31]

gorm.Updates(updateFields)

// after dataBase { name: mark , age: 31 , address: "New York" , email: "alice@example.com" , password: "1234567890" }
```

### 2. Copy Struct Fields
The `CopyFields` function copies fields from the source struct to the destination struct, commonly used to return specified parameters. It only copies fields from src to dst if dst is empty and has the same name and type as src.
Notice: This function does not handle nested structs and only processes struct pointers.

### Example
```go
// before src  { name: mark , age: 31 , address: "New York" , email: "alice@example.com" , password: "1234567890"...other fields }
// before dst  { name: "" , age: 0 , address: "" , email: "" ... other fields }

err := gref.CopyFields(&src, &dst)

// after dst  { name: mark , age: 31 , address: "New York" , email: "alice@example.com" ... other fields }
```