package search_query

// regex_match
type regexMatch struct {
        baseComparison
}

func (rm regexMatch) Operator() string {
        return "regex_match"
}

func RegexMatch(val1,val2 string) SearchQuery {
        return regexMatch{baseComparison{val1,val2}}
}


// regex_not_match
type regexNotMatch struct {
        baseComparison
}

func (rnm regexNotMatch) Operator() string {
        return "regex_not_match"
}

func RegexNotMatch(val1,val2 string) SearchQuery {
        return regexNotMatch{baseComparison{val1,val2}}
}


