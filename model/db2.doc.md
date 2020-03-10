# DB2连接使用说明

## DB2产生的原因

为了适配其他数据库的明明不服和gorm的命名标准

## DB2的使用特点

* DB2是普通sql连接，需要自己写sql

* DB2得到的结果无法直接放入结构体，需要自己解出数据，然后放入结构体

## DB2的使用方法

```go
// 查询
query2, err := db.Query("select * from tmpdb.tmptab where id =?",1)
printResult(query2)

// 解数据过程
func printResult(query *sql.Rows) {
        column, _ := query.Columns()              //读出查询出的列字段名
        values := make([][]byte, len(column))     //values是每个列的值，这里获取到byte里
        scans := make([]interface{}, len(column)) //因为每次查询出来的列是不定长的，用len(column)定住当次查询的长度
        for i := range values {                   //让每一行数据都填充到[][]byte里面
                scans[i] = &values[i]
        }
        results := make(map[int]map[string]string) //最后得到的map
        i := 0
        for query.Next() { //循环，让游标往下移动
                if err := query.Scan(scans...); err != nil { //query.Scan查询出来的不定长值放到scans[i] = &values[i],也就是每行都放在values里
                        fmt.Println(err)
                        return
                }
                row := make(map[string]string) //每行数据
                for k, v := range values {     //每行数据是放在values里面，现在把它挪到row里
                        key := column[k]
                        row[key] = string(v)
                }
                results[i] = row //装入结果集中
                i++
        }
        for k, v := range results { //查询出来的数组
                fmt.Println(k, v)
        }
}
```
