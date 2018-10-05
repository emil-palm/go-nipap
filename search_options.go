package nipap

import (
)

type SearchOptions struct {
	// How many levels of parents to return. Set to :data:`-1` to include all parents.
	ParentsDepth		int	`structs:"parents_depth"`
	// How many levels of children to return. Set to :data:`-1` to include all children.
	ChildrensDepth		int	`structs:"children_depth,omitempty"`
	// Include all parents, no matter what depth is specified.
	IncludeAllParents	bool	`structs:"include_all_parents,omitempty"`
	// Include all children, no matter what depth is specified.
	IncludeAllChildren	bool	`structs:"include_all_children,omitempty"`
	// The maximum number of prefixes to return (default :data:`50`).
	MaxResults		int	`structs:"max_result,omitempty"`
	// Offset the result list this many prefixes (default :data:`0`).
	Offset			int	`structs:"offset,omitempty"`
}

