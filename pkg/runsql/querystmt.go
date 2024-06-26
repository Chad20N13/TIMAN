package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	test "timan/pkg/test"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	Ctx     context.Context
	MySQLDB *sql.DB
}

func Query(ctx context.Context, db *sql.DB, querySQL string) ([]string, []map[string]string, error) {
	var (
		cols []string
		res  []map[string]string
	)
	rows, err := db.QueryContext(ctx, querySQL)
	if err != nil {
		return cols, res, fmt.Errorf("general sql [%v] query failed: [%v]", querySQL, err.Error())
	}
	defer rows.Close()

	//不确定字段通用查询，自动获取字段名称
	cols, err = rows.Columns()
	if err != nil {
		return cols, res, fmt.Errorf("general sql [%v] query rows.Columns failed: [%v]", querySQL, err.Error())
	}

	values := make([][]byte, len(cols))
	scans := make([]interface{}, len(cols))
	for i := range values {
		scans[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scans...)
		if err != nil {
			return cols, res, fmt.Errorf("general sql [%v] query rows.Scan failed: [%v]", querySQL, err.Error())
		}

		row := make(map[string]string)
		for k, v := range values {
			// Oracle/Mysql 对于 'NULL' 统一字符 NULL 处理，查询出来转成 NULL,所以需要判断处理
			// 查询字段值 NULL
			// 如果字段值 = NULLABLE 则表示值是 NULL
			// 如果字段值 = "" 则表示值是空字符串
			// 如果字段值 = 'NULL' 则表示值是 NULL 字符串
			// 如果字段值 = 'null' 则表示值是 null 字符串
			if v == nil {
				row[cols[k]] = "NULLABLE"
			} else {
				// 处理空字符串以及其他值情况
				// 数据统一 string 格式显示
				row[cols[k]] = string(v)
			}
		}
		res = append(res, row)
	}

	if err = rows.Err(); err != nil {
		return cols, res, fmt.Errorf("general sql [%v] query rows.Next failed: [%v]", querySQL, err.Error())
	}
	return cols, res, nil
}

func (m *MySQL) GetMySQLTableColumn(schemaName, tableName string) ([]map[string]string, error) {
	var (
		res []map[string]string
		err error
	)

	_, res, err = Query(m.Ctx, m.MySQLDB, fmt.Sprintf(`SELECT COLUMN_NAME,
		DATA_TYPE,
		IFNULL(CHARACTER_MAXIMUM_LENGTH,0) DATA_LENGTH,
		IFNULL(NUMERIC_SCALE,0) DATA_SCALE,
		IFNULL(NUMERIC_PRECISION,0) DATA_PRECISION,
		IFNULL(DATETIME_PRECISION,0) DATETIME_PRECISION,
		IF(IS_NULLABLE = 'NO', 'N', 'Y') NULLABLE,
		IFNULL(COLUMN_DEFAULT,'NULLSTRING') DATA_DEFAULT,
		IFNULL(COLUMN_COMMENT,'') COMMENTS,
		IFNULL(CHARACTER_SET_NAME,'UNKNOWN') CHARACTER_SET_NAME,
		IFNULL(COLLATION_NAME,'UNKNOWN') COLLATION_NAME
 FROM information_schema.COLUMNS
 WHERE UPPER(TABLE_SCHEMA) = UPPER('%s')
   AND UPPER(TABLE_NAME) = UPPER('%s')
 ORDER BY ORDINAL_POSITION`, schemaName, tableName))

	if err != nil {
		return res, err
	}
	return res, nil
}

func main() {
	db, err := sql.Open("mysql", "root:hurricane@tcp(120.92.210.161:4000)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT VARIABLE_NAME,VARIABLE_VALUE FROM mysql.tidb")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("|-------------------------------------------|")
		fmt.Println("|", id, "|", ":", name, "|")

	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	//模拟一个数组
	// var a [3]string
	// a[0] = "array index 1"
	// a[1] = "array index 2"
	// a[2] = "array index 3"
	// fmt.Println(a)
	// //模拟一个map
	// m := map[string]string{
	// 	"tikv_gc_leader_uuid":  "62c0d0bc16c000a",
	// 	"tikv_gc_run_interval": "10m0s",
	// }
	// fmt.Println(m)
	test.Dddd()
}
