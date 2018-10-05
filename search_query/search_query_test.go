package search_query

import (
	"testing"
)


func Test_Search_Query(t *testing.T) {
	eq1 := Equals("foo1","foo2")
	eq2 := Equals("bar1","bar2")
	and := And(eq1,eq2)
	or := Or(and,eq2)
	t.Fatalf("%+v\n", Map(or))
}
