# gorm-kingbase-driver
A GORM driver adapter for KingbaseES (人大金仓数据库 GORM 驱动适配层)

本项目是一个基于 **GORM v2** 的 **人大金仓数据库（[KingbaseES](https://www.kingbase.com.cn/)**）适配驱动，旨在让开发者能够在使用 [GORM](https://gorm.io/index.html) 框架时，无缝地连接和操作人大金仓数据库。  
本项目在官方 [gokb](https://www.kingbase.com.cn/download.html#drive) 驱动基础上进行封装与扩展，使其具备更好的兼容性与迁移支持，适合在企业级 Go 项目中直接使用。

---

## ✨ 特性

- ✅ **基于 GORM v2 实现**，无缝集成 GORM ORM 框架  
- ✅ **兼容人大金仓官方 `gokb` 驱动**  
- ✅ **提供自定义 GORM 驱动类**（参考：[GORM 自定义驱动文档](https://gorm.io/zh_CN/docs/write_driver.html)）  
- ✅ **提供 Migrate 迁移类**，支持自动迁移、结构变更等操作（参考：[GORM Migration 文档](https://gorm.io/zh_CN/docs/migration.html)）  
- ✅ 支持基本的 CRUD 操作与事务管理  
- ✅ 支持事务与连接池管理  
- ✅ 支持原生 SQL 查询与 GORM 表达式
  
---

## 📦 快速开始
```bash
go get github.com/heqiang6/gorm-kingbase-driver
```


## 🚀 使用示例
```bash
package main

import (
    "fmt"
    "gorm.io/gorm"
	kingbase "github.com/heqiang6/gorm-kingbase-driver"
)

type User struct {
    ID   int64  `gorm:"primaryKey"`
    Name string `gorm:"size:100"`
}

func main() {
    dsn := "host=127.0.0.1 user=system password=123456 dbname=test port=54321 sslmode=disable"

    // 使用 Kingbase 驱动连接数据库
	db, err := gorm.Open(kingbase.Open(dsn), &gorm.Config{}) // 具体写法根据项目代码结构可能不同
    if err != nil {
        panic(fmt.Sprintf("failed to connect to Kingbase database: %v", err))
    }

    // 自动迁移
    db.AutoMigrate(&User{}, &Tag{}, &DataSource{})

    // 插入示例
    db.Create(&User{Name: "Kingbase User"})

    // 查询示例
    var user User
    db.First(&user)
    fmt.Println("User:", user)
}
```
