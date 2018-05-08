package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	/*DSN数据源名称
	  [username[:password]@][protocol[(address)]]/dbname[?param1=value1&paramN=valueN]
	  user@unix(/path/to/socket)/dbname
	  user:password@tcp(localhost:5555)/dbname?charset=utf8&autocommit=true
	  user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname?charset=utf8mb4,utf8
	  user:password@/dbname
	  无数据库: user:password@/
	*/
	fmt.Println("连接数据库")
	db, err := sql.Open("mysql", "test:test@tcp(127.0.0.1:3306)/?charset=utf8")
	checkErr(err)

	defer db.Close()

	fmt.Println("初始化数据库，初始化表，初始化数据")
	db.Exec("drop database if exists tmpdb")
	db.Exec("create database tmpdb")
	db.Exec("create table tmpdb.tmptab(c1 int not null, c2 varchar(20), c3 varchar(20), PRIMARY KEY (c1))")
	db.Exec("insert into tmpdb.tmptab values(101, '张三', '广州'), (102, '李四', '杭州'), (103, '王五', '北京'), (104, '赵六', '上海')")

	query0, err0 := db.Query("select * from tmpdb.tmptab")
	checkErr(err0)
	printResult(query0)

	fmt.Println("删除部分数据")
	db.Exec("delete from tmpdb.tmptab where c1=101")
	query1, err1 := db.Query("select * from tmpdb.tmptab")
	checkErr(err1)
	printResult(query1)

	fmt.Println("更新部分数据")
	db.Exec("update tmpdb.tmptab set c3='New York' where c1=102")
	query2, err2 := db.Query("select * from tmpdb.tmptab")
	checkErr(err2)
	printResult(query2)

	fmt.Println("清空数据")
	db.Exec("delete from tmpdb.tmptab")
	query3, err3 := db.Query("select * from tmpdb.tmptab")
	checkErr(err3)
	printResult(query3)

	fmt.Println("删除表，删除数据库，断开连接")
	db.Exec("drop table tmpdb.tmptab")
	db.Exec("drop database tmpdb")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func printResult(query *sql.Rows) {
	defer query.Close()

	column, _ := query.Columns()              // 查询出来的列字段名
	values := make([][]byte, len(column))     // 每个列的值
	scans := make([]interface{}, len(column)) // 当前查询的结果索引
	for i := range values {                   // 每一行数据填充到[][]byte中
		scans[i] = &values[i]
	}

	results := make(map[int]map[string]string) // 最后的完整结果
	rowNUM := 0
	for query.Next() { // 移动游标
		if err := query.Scan(scans...); err != nil {
			checkErr(err)
		}

		row := make(map[string]string) // 每行数据
		for k, v := range values {     // 每行数据从values挪到row中
			key := column[k]
			row[key] = string(v)
		}

		results[rowNUM] = row
		rowNUM++
	}

	for k, v := range results {
		fmt.Println(k, v)
	}
}
