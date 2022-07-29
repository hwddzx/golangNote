package main

/*
1. 导入声明
import "crypto/rand"  // 默认模式: rand.Function
import R "crypto/rand" // 包重命名: R.Function
import . "crypto/rand" // 简便模式: Function
import _ "crypto/rand"  // 匿名导入: 仅让该包执行初始化函数

import (
	mrand "math/rand"  // 包重命名: mrand.Function
)
*/

/*
2. 匿名导入

	如果只是导入一个包而并不使用导入的包将会导致一个编译错误, 但是有时候我们只是想利用导入包而产生的副作用: 它会计算包级变量的初始化表达式和执行
导入包的init初始化函数. 这时候需要抑制"unused import"编译错误, 可以用下划线来重命名导入的包.

import _ "image/png"  // PNG解码器

数据库包database/sql就是采用了类似的技术, 让用户可以根据自己的需要选择导入必要的数据库驱动
例如:
import (
	"database/sql"
	_ "github.com/lib/pq"  				// 启用对Postgres的支持
	_ "github.com/go-sql-driver/mysql"	// 启用对MySQL的支持
)

db, err := sql.Open("postgres", dbname)  	// OK
db, err := sql.Open("mysql", dbname)  		// OK
db, err := sql.Open("sqlite3", dbname)  	// 返回错误: unknown driver "sqlite3"

*/
