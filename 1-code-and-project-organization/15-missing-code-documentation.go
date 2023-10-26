package codeandprojectorganization

/*
Mistake 15: Missing code documentation

"First, every exported element must be documented.
Whether it is a structure, an interface, a function, or something else, if it’s exported, it must be documented.
The convention is to add comments, starting with the name of the exported element. For example,"
``
// Customer is a customer representation.
type Customer struct{}

// ID returns the customer identifier.
func (c Customer) ID() string { return "" }
``

Deprecated elements

It’s possible to deprecate an exported element using the // Deprecated: comment this way:
``
// ComputePath returns the fastest path between two points.
// Deprecated: This function uses a deprecated way to compute
// the fastest path. Use ComputeFastestPath instead.
func ComputePath() {}
``
Then, if a developer uses the ComputePath function, they should get a warning. (Most IDEs handle deprecated comments.)



*/
