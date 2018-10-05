package search_query

import (
	//"github.com/fatih/structs"
)

type SearchQuery interface {
	Val1()		interface{}
	Val2()		interface{}
	Operator()	string
}

func Map(sq SearchQuery) interface{} {
	return map[string]interface{}{
		"operator":sq.Operator(),
		"val1":sq.Val1(),
		"val2":sq.Val2(),
	}
}

