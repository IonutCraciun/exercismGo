// Package listops
package listops

type IntList []int
type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

// Append returns a new slice that contains l1 + l2
func (l1 IntList) Append(l2 IntList) IntList {
    if len(l1) == 0 && len(l2) == 0{
        return l1
    }
    var results IntList
    for _, val := range l1 {
        results = append(results, val)
    }
    for _, val := range l2 {
        results = append(results, val)
    }
    return results
}

// Concat
func (l1 IntList) Concat(ll []IntList) IntList {
    results := l1
    for i, _ := range ll {
        results = results.Append(ll[i])
    }
    return results
}

// Filter
func (l1 IntList) Filter(predicate predFunc) IntList {
    if len(l1) == 0 {
        return l1
    }
    var results IntList
    for _, val := range l1 {
        if predicate(val) {
            results = append(results, val)
        }
    }
    return results
}

// Length
func (l1 IntList) Length() int {
    return len(l1)
}

// Map
func (l1 IntList) Map(predicate unaryFunc) IntList {
    for k, val := range l1 {
        l1[k] = predicate(val)
    }
    return l1
}

// Foldl
func (l1 IntList) Foldl(predicate binFunc, x int) int {
    if len(l1) == 0 {
        return x
    }
    results := x
    for _ , val := range l1 {
        if results == 0 {
            return results
        }
        results = predicate(results, val)
    }
    return results
}

// Foldr
func (l1 IntList) Foldr(predicate binFunc, x int) int {
    if len(l1) == 0 {
        return x
    }
    results := x
    for key , _ := range l1 {
        if results == 0 {
            return results
        }
        results = predicate(l1[len(l1) - key - 1], results)
    }
    return results
}

// Reverse
func (l1 IntList) Reverse() IntList {
    var results IntList

    if len(l1) == 0 {
        return l1
    }
    for k, _ := range l1 {
        results = append(results, l1[len(l1) - k - 1])
    }
    return results
}