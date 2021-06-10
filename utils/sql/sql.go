package sql

import (
	"strings"
)

func And(args ...interface{}) []interface{} {
	m := []interface{}{"AND", args}
	return m
}
func Or(args ...interface{}) []interface{} {
	m := []interface{}{"OR", args}
	return m
}
func OrList(args [][]interface{}) []interface{} {
	/* AND (xx=x1 or xx=x2)*/
	m := []interface{}{"AND", args}
	return m
}

func ParseFilter(filters [][]interface{}) (string, []interface{}) {
	args := make([]interface{}, 0)
	where := make([]string, 0)
	for _, filter := range filters {
		c := filter[0].(string)
		switch wheres := filter[1].(type) {
		case []interface{}:
			switch len(wheres) {
			case 1:
				switch paramReal := wheres[0].(type) {
				case string:
					where = append(where, c+" ("+paramReal+")")
				case map[string]interface{}:
					for k, v := range paramReal {
						where = append(where, " "+c+" "+k+" = ?")
						args = append(args, v)
					}
				}
			case 2:
				where = append(where, " "+c+" "+wheres[0].(string)+" = ?")
				args = append(args, wheres[1])
			case 3:
				name := wheres[0].(string)
				value := wheres[2]
				operator := wheres[1].(string)
				switch operator {
				case "like":
					where = append(where, " "+c+" "+name+"::text"+" LIKE ?")
				case "not like":
					where = append(where, " "+c+" "+name+"::text"+" NOT LIKE ?")
				case "in":
					where = append(where, " "+c+" "+name+" IN (?) ")
				case "not in":
					where = append(where, " "+c+" "+name+" NOT IN (?)")
				default:
					where = append(where, " "+c+" "+name+" "+operator+" ?")
				}
				args = append(args, value)
			}
		case [][]interface{}:
			// 嵌套的情况
			subWhere := make([]string, 0)
			for _, subFilter := range wheres {
				cc := subFilter[0].(string)
				wheres2 := subFilter[1].([]interface{})
				switch len(wheres2) {
				case 1:
					switch paramReal := wheres2[0].(type) {
					case string:
						subWhere = append(subWhere, cc+" ("+paramReal+")")
					case map[string]interface{}:
						for k, v := range paramReal {
							subWhere = append(subWhere, " "+cc+" "+k+" = ?")
							args = append(args, v)
						}
					}
				case 2:
					subWhere = append(subWhere, " "+cc+" "+wheres2[0].(string)+" = ?")
					args = append(args, wheres2[1])
				case 3:
					name := wheres2[0].(string)
					value := wheres2[2]
					operator := wheres2[1].(string)
					switch operator {
					case "like":
						subWhere = append(subWhere, " "+cc+" "+name+"::text"+" LIKE ?")
					case "not like":
						subWhere = append(subWhere, " "+cc+" "+name+"::text"+" NOT LIKE ?")
					case "in":
						subWhere = append(subWhere, " "+cc+" "+name+" IN (?) ")
					case "not in":
						subWhere = append(subWhere, " "+cc+" "+name+" NOT IN (?)")
					default:
						subWhere = append(subWhere, " "+cc+" "+name+" "+operator+" ?")
					}
					args = append(args, value)
				}
			}
			subQuery := strings.TrimLeft(strings.Trim(strings.Join(subWhere, " "), " "), "AND")
			subQuery = strings.TrimLeft(subQuery, "OR")
			where = append(where, c+" ("+subQuery+")")
		}
	}
	query := strings.TrimLeft(strings.Trim(strings.Join(where, " "), " "), "AND")
	query = strings.TrimLeft(query, "OR")
	return query, args
}

func ParsePage(page, limit int) (int, int) {
	if page == 0 {
		page = 1
	}
	if limit == 0 || limit > 100 {
		limit = 10
	}
	return (page - 1) * limit, limit
}
