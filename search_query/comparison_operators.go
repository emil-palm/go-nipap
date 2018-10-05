package search_query

import (
	"fmt"
)

type baseComparison struct {
	val1,val2 interface{}
}

func (bc baseComparison) Operator() string {
	return ""
}

func (bc baseComparison) Val1() interface{} {
	return fmt.Sprintf("%v", bc.val1)
}
func (bc baseComparison) Val2() interface{} {
	return fmt.Sprintf("%v", bc.val2)
}

// Equals
type equals struct {
	baseComparison
}

func (c equals) Operator() string {
	return "equals"
}

func Equals(val1,val2 string) SearchQuery {
	return equals{baseComparison{val1,val2}}
}

// Not Equals
type notEquals struct {
	baseComparison
}

func (c notEquals) Operator() string {
	return "not_equals"
}

func NotEquals(val1,val2 string) SearchQuery {
	return notEquals{baseComparison{val1,val2}}
}

// less
type less struct {
        baseComparison
}

func (l less) Operator() string {
        return "less"
}

func Less(val1,val2 string) SearchQuery {
        return less{baseComparison{val1,val2}}
}

// less_or_equal
type lessOrEqual struct {
        baseComparison
}

func (loe lessOrEqual) Operator() string {
        return "less_or_equal"
}

func LessOrEqual(val1,val2 string) SearchQuery {
        return lessOrEqual{baseComparison{val1,val2}}
}


// greater
type greater struct {
        baseComparison
}

func (g greater) Operator() string {
        return "greater"
}

func Greater(val1,val2 string) SearchQuery {
        return greater{baseComparison{val1,val2}}
}


// greater_or_equal
type greaterOrEqual struct {
        baseComparison
}

func (goe greaterOrEqual) Operator() string {
        return "greater_or_equal"
}

func GreaterOrEqual(val1,val2 string) SearchQuery {
        return greaterOrEqual{baseComparison{val1,val2}}
}


// like
type like struct {
        baseComparison
}

func (l like) Operator() string {
        return "like"
}

func Like(val1,val2 string) SearchQuery {
        return like{baseComparison{val1,val2}}
}



// contains
type contains struct {
        baseComparison
}

func (c contains) Operator() string {
        return "contains"
}

func Contains(val1,val2 string) SearchQuery {
        return contains{baseComparison{val1,val2}}
}


// equals_any
type equalsAny struct {
        baseComparison
}

func (ea equalsAny) Operator() string {
        return "equals_any"
}

func EqualsAny(val1,val2 string) SearchQuery {
        return equalsAny{baseComparison{val1,val2}}
}


// contains_equals
type containsEquals struct {
        baseComparison
}

func (ce containsEquals) Operator() string {
        return "contains_equals"
}

func ContainsEquals(val1,val2 string) SearchQuery {
        return containsEquals{baseComparison{val1,val2}}
}


// contained_within
type containedWithin struct {
        baseComparison
}

func (cw containedWithin) Operator() string {
        return "contained_within"
}

func ContainedWithin(val1,val2 string) SearchQuery {
        return containedWithin{baseComparison{val1,val2}}
}


// contained_within_equals
type containedWithinEquals struct {
        baseComparison
}

func (cwe containedWithinEquals) Operator() string {
        return "contained_within_equals"
}

func ContainedWithinEquals(val1,val2 string) SearchQuery {
        return containedWithinEquals{baseComparison{val1,val2}}
}

