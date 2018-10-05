package search_query

// Hidden compound query object
type compound struct {
	val1,val2 SearchQuery
}

func (c compound) Operator() string {
	return ""
}

func (c compound) Val1() interface{} {
	return Map(c.val1)
}

func (c compound) Val2() interface{} {
	return Map(c.val2)
}


// Hidden OR object which uses the compound
type or struct { compound }
func (o or) Operator() string { return "or" }


func Or(val1, val2 SearchQuery) (SearchQuery) {
	return or{compound{val1,val2}}
}

// Hidden AND object which uses the compound
type and struct { compound } 
func (a and) Operator() string { return "and" }

func And(val1,val2 SearchQuery) (SearchQuery) {
	return and{compound{val1,val2}}
}

