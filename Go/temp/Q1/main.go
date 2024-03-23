package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Runnning Implementation 1")
	tc1 := NewTokenCounterImpl1()
	startTime := time.Now()
	Test(tc1)
	fmt.Println("Execute time =", time.Since(startTime))
	fmt.Println()

	fmt.Println("Runnning Implementation 2")
	startTime = time.Now()
	tc2 := NewTokenCounterImpl2()
	Test(tc2)
	fmt.Println("Execute time =", time.Since(startTime))
}

func Test(tc TokenCounter) {
	fmt.Println("Original case")
	tc.Ingest("oursky:uk:dev")
	tc.Ingest("oursky:hk:design")
	tc.Ingest("oursky:hk:pm")
	tc.Ingest("oursky:hk:dev")
	tc.Ingest("skymaker")
	fmt.Println("Appearance('oursky'):", tc.Appearance("oursky"))
	fmt.Println("Appearance('oursky:hk'):", tc.Appearance("oursky:hk"))

	tc.Ingest("skymaker:london:ealing:dev")
	tc.Ingest("skymaker:london:design")
	tc.Ingest("skymaker:london:design")
	tc.Ingest("skymaker:man:pm")
	tc.Ingest("skymaker:man:pm")
	fmt.Println("Appearance('skymaker:man'):", tc.Appearance("skymaker:man"))
	fmt.Println("Appearance('skymaker:lon'):", tc.Appearance("skymaker:lon"))
	fmt.Println("Appearance('london:skymaker'):", tc.Appearance("london:skymaker"))

	// fmt.Println("Customize case")
	// fmt.Println("case 1")
	// tc.Init()
	// tc.Ingest("a:b:c")
	// fmt.Println(tc.Appearance("a"))     // 1
	// fmt.Println(tc.Appearance("b"))     // 1
	// fmt.Println(tc.Appearance("c"))     // 1
	// fmt.Println(tc.Appearance("a:b"))   // 1
	// fmt.Println(tc.Appearance("b:c"))   // 1
	// fmt.Println(tc.Appearance("a:b:c")) // 1
	// fmt.Println(tc.Appearance("a:c"))   // 0

	// fmt.Println("case 2")
	// tc.Init()
	// tc.Ingest("a:b")
	// fmt.Println(tc.Appearance("a"))   // 1
	// fmt.Println(tc.Appearance("b"))   // 1
	// fmt.Println(tc.Appearance("a:b")) // 1
	// fmt.Println(tc.Appearance("c"))   // 0

	// fmt.Println("case 3")
	// tc.Init()
	// tc.Ingest("a")
	// fmt.Println(tc.Appearance("a")) // 1
	// fmt.Println(tc.Appearance("c")) // 0

	// fmt.Println("case 4")
	// tc.Init()
	// tc.Ingest("a:b")
	// tc.Ingest("b:c")
	// fmt.Println(tc.Appearance("a"))     // 0.5
	// fmt.Println(tc.Appearance("b"))     // 1
	// fmt.Println(tc.Appearance("c"))     // 0.5
	// fmt.Println(tc.Appearance("a:b"))   // 0.5
	// fmt.Println(tc.Appearance("a:b:c")) // 0
}
