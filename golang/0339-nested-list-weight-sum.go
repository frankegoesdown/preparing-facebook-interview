package main

type NestedInteger struct {
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (n NestedInteger) IsInteger() (res bool) { return }

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (n NestedInteger) GetInteger() (res int) { return }

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (n *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (n NestedInteger) GetList() (res []*NestedInteger) { return }

func main() {

}

func depthSum(nestedList []*NestedInteger) int {
	return helper(nestedList, 1)
}

func helper(list []*NestedInteger, d int) int {
	sum := 0
	for _, l := range list {
		if l.IsInteger() {
			sum += l.GetInteger() * d
		} else {
			sum += helper(l.GetList(), d+1)
		}
	}
	return sum

}
