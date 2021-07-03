/**
* @Author: lik
* @Date: 2021/3/9 20:53
* @Version 1.0
 */
package model

import (
	"fmt"
	"gitee.com/open-product/dtcloud-api/app/util/json_params"
	"strconv"
	"strings"
)

func (ml *MemberUsers) CreateParamsField(cf map[string]interface{}, dat map[string]interface{}) bool {

	data := json_params.CloneTags(cf)
	delete(cf, "form-field-m2m")
	keys, values := json_params.GetKeys(cf)
	columns := strings.Join(keys, ",")
	table := strings.Replace(json_params.SliceValues(dat["model"].([]interface{})), ".", "_", -1)
	result := json_params.GetValues(values)

	s := fmt.Sprintf("INSERT INTO %s(%s) VALUES(%s) RETURNING ID", table, columns, result)
	rows, err := ml.Raw(s).Rows() // (*sql.Rows, error)
	if err != nil {
		return false
	}
	defer rows.Close()
	for rows.Next() {
		_ = rows.Scan(&ml.Id)
	}

	//fmt.Println(ml.Id)
	mid := ml.Id
	return ml.CreateMany(data, strconv.FormatInt(mid, 10))

}

type DataRaw struct {
	RelationTable string `json:"relation_table"`
	Column1       string `json:"column1"`
	Column2       string `json:"column2"`
}

func (ml *MemberUsers) CreateMany(data map[string]interface{}, mid string) bool {
	// ir_model_fields ir_model ir_model_relation

	for k := range data {

		m2m := data[k].(map[string]interface{})
		ids := json_params.SliceValues(m2m["ids"].([]interface{}))
		model := json_params.SliceValues(m2m["model"].([]interface{}))
		table := strings.Replace(json_params.SliceValues(m2m["model"].([]interface{})), ".", "_", -1)
		fields := m2m["m2m_fields"]
		var records DataRaw

		s := fmt.Sprintf("SELECT relation_table,column1,column2 FROM ir_model_fields where ttype = 'many2many' and name = '%s' and model = '%s'", fields, model)
		err := ml.Raw(s).Scan(&records).Error
		if err != nil {
			return false
		}

		tables := records.RelationTable
		columns := strings.Join([]string{records.Column1, records.Column2}, ",")

		if m2m["many2many"] == 6 {

			s1 := fmt.Sprintf("DELETE FROM %s where %s=%s", tables, records.Column1, mid)
			err := ml.Exec(s1).Error
			if err != nil {
				ml.DeleteMany(table, mid)
				return false
			}

			values := strings.Split(ids, ",")
			for v := range values {

				results := strings.Join([]string{mid, values[v]}, ",")
				s2 := fmt.Sprintf("INSERT INTO %s (%s) VALUES(%s)", tables, columns, results)
				err = ml.Exec(s2).Error

				if err != nil {
					ml.DeleteMany(table, mid)
					return false
				}

			}
			return true

		} else if m2m["many2many"] == 5 {

			s1 := fmt.Sprintf("DELETE FROM %s where %s=%s", tables, records.Column1, mid)
			err := ml.Exec(s1).Error
			if err != nil {
				ml.DeleteMany(table, mid)
				return false
			}

			return true

		} else if m2m["many2many"] == 4 {

			values := strings.Split(ids, ",")
			for v := range values {

				results := strings.Join([]string{mid, values[v]}, ",")
				s2 := fmt.Sprintf("INSERT INTO %s (%s) VALUES(%s)", tables, columns, results)
				err = ml.Exec(s2).Error

				if err != nil {
					ml.DeleteMany(table, mid)
					return false
				}

			}
			return true

		} else if m2m["many2many"] == 3 {

			fmt.Println(m2m)
		} else if m2m["many2many"] == 2 {

			fmt.Println(m2m)
		} else if m2m["many2many"] == 1 {

			fmt.Println(m2m)
		} else if m2m["many2many"] == 0 {

			fmt.Println(m2m)
		}
	}
	return true
}

func (ml *MemberUsers) DeleteMany(table, mid string) {
	ml.Table(table).Delete(mid)
}

func (ml *MemberUsers) PublicRead(fields string, ids string, page int64, limit int64, order string, table string) []map[string]interface{} {

	var wh = ""
	var pl = ""
	var od = ""
	var s = ""

	if table != "" {
		if fields == "" {
			fields = "*"
		}

		if page != -1 && limit != 0 {
			pegs := (page - 1) * limit
			pl += fmt.Sprintf("LIMIT %d OFFSET %d", limit, pegs)
		}
		if order != "" {
			od += fmt.Sprintf("order by %s", order)
		}

		if ids != "" {
			wh += fmt.Sprintf("WHERE id = any(array[%s])", ids)
			s = fmt.Sprintf("SELECT %s FROM %s %s %s %s", fields, table, wh, od, pl)

		} else {
			s = fmt.Sprintf("SELECT %s FROM %s %s %s %s", fields, table, od, pl, wh)
		}

		rows, _ := ml.Raw(s).Rows()

		columns, _ := rows.Columns()
		columnLength := len(columns)
		cache := make([]interface{}, columnLength) //临时存储每行数据
		for index, _ := range cache {              //为每一列初始化一个指针
			var a interface{}
			cache[index] = &a
		}
		var list []map[string]interface{} //返回的切片
		for rows.Next() {
			_ = rows.Scan(cache...)

			item := make(map[string]interface{})
			for i, data := range cache {
				item[columns[i]] = *data.(*interface{}) //取实际类型
			}
			list = append(list, item)
		}
		_ = rows.Close()
		return list
	}

	return nil

}
