# ORM - Gorm Wrapper for Kwiscale

[![Coverage Status](https://coveralls.io/repos/github/kwiscale/orm/badge.svg?branch=master)](https://coveralls.io/github/kwiscale/orm?branch=master)


[Gorm](https://github.com/jinzhu/gorm) is a complete ORM for Go. Kwiscale/Gorm is a simple wrapper to be able to implement database in framework.


# Usage

In your main package, simply import and initialize database connection.

```go
package main

import (
	"gopkg.in/kwiscale/framework.v1"
	"gopkg.in/kwiscale/orm.v1"
	_ "gopkg.in/kwiscale/orm.v1/dialects/sqlite"
)

// User is a simple orm.Model
type User struct {
	orm.Model
	Name string
	Age  int
}

func init() {
	orm.Initialize("sqlite", "./test.db")
	db := orm.Get()
	defer db.Close()
    // Create table, update if needed
	db.AutoMigrate(&User{})
}

```

Then, you may insert, delete, find, and so on. 


```go
user := User{
    Name: "Foo",
    Age: 42,
}

db := orm.Get()
defer db.Close()

db.Create(&user)

// find a user and set it on
// "res" structure
res:=User{}
db.First(&res, &User{Name:"Foo"})
```

To get more information on requests, see [Gorm documentation](http://jinzhu.me/gorm/)

