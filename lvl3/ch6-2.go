// For statements (Chapter 6)
// https://golang.org/ref/spec#For_statements
// https://gobyexample.com/for

// A "for" statement specifies repeated execution of a block. There are three forms: The iteration may be controlled by a single condition, a "for" clause, or a "range" clause.
// ForStmt = "for" [ Condition | ForClause | RangeClause ] Block .
// Condition = Expression .

// For statements with single condition:
// 	for a < b {
// 		a *= 2
// 	}

// For statements with for clause:
// ForClause = [ InitStmt ] ";" [ Condition ] ";" [ PostStmt ] .
// 	for i := 0; i < 10; i++ {
// 		f(i)
// 	}

// For statements with range clause
// RangeClause = [ ExpressionList "=" | IdentifierList ":=" ]
// "range" Expression .
// 	ex.1
// 	var a [10]string
// 	for i, s := range a {
// 		// type of i is int
// 		// type of s is string
// 		// s == a[i]
// 		g(i, s)
// 	}

// 	ex.2
// 	var testdata *struct {
// 		a *[7]int
// 	}
// 	for i, _ := range testdata.a {
// 		// testdata.a is never evaluated; len(testdata.a) is constant
// 		// i ranges from 0 to 6
// 		f(i)
// 	}

// Ellen Körbes For statements with for clause example

package main

import (
	"fmt"
)

func main() {
	for x := 0; x < 10; x++ {
		//	^IStmt  ^Cond   ^PStmt
		fmt.Println(x)
	}
}
