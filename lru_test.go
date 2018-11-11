package main

import "testing"


func Test(t *testing.T){
	testLru(t)
	testGetIndexOf(t)
	testFindLowest(t)
}

func testLru(t *testing.T) {
	// test case 1
	cacheSize := 3
	accessSeq := []string{"a", "b", "c", "d"}
	solution := []string{"d", "b", "c"}
	test := Lru(cacheSize, accessSeq, false)
	for i := 0; i < len(test); i++{
		if (i >= len(solution)) {
			t.Errorf("testLru(): tested answer has length %d instead of %d", len(test), len(solution))
		}
		if (test[i] != solution[i]) {
			t.Errorf("testLru(): Wrong replacement. Expected %s instead of %s at index %d", solution[i], test[i], i)
		}
	}
	// test case 2
	cacheSize = 4
	accessSeq = []string{"a", "b", "c", "d", "a", "f"}
	solution = []string{"a", "f", "c", "d"}
	test = Lru(cacheSize, accessSeq, false)
	for i := 0; i < len(test); i++{
		if (i >= len(solution)) {
			t.Errorf("testLru(): tested answer has length %d instead of %d", len(test), len(solution))
		}
		if (test[i] != solution[i]) {
			t.Errorf("testLru(): Wrong replacement. Expected %s instead of %s at index %d", solution[i], test[i], i)
		}
	}
}

func testGetIndexOf(t *testing.T) {
	// test case 1
	testSeq := []string{"a", "b", "c", "d"}
	testChar := "c"
	solution := 2
	answer := GetIndexOf(testSeq, testChar)
	if ( answer != solution) {
			t.Errorf("testGetIndexOf(): string found at wrong index. Expected %d, got %d", solution, answer)
	}

	// test case 2
	testSeq = []string{"a", "b", "c", "d"}
	testChar = "x"
	solution = -1
	answer = GetIndexOf(testSeq, testChar)
	if ( answer != solution) {
			t.Errorf("testGetIndexOf(): string found at wrong index. Expected %d, got %d", solution, answer)
	}

	// test case 3
	testSeq = []string{"a", "a", "a", "a"}
	testChar = "a"
	solution = 0
	answer = GetIndexOf(testSeq, testChar)
	if ( answer != solution) {
			t.Errorf("testGetIndexOf(): string found at wrong index. Expected %d, got %d", solution, answer)
	}
}

func testFindLowest(t *testing.T) {
	// test case 1
	testArr := []int{1, 2, 3, 4, 5}
	solution := 0
	answer := FindLowest(testArr)
	if ( answer != solution) {
			t.Errorf("testFindLowest(): got wrong index. Expected %d, got %d", solution, answer)
	}

	// test case 2
	testArr = []int{11, 22, 33, 4, 5}
	solution = 3
	answer = FindLowest(testArr)
	if ( answer != solution) {
			t.Errorf("testFindLowest(): got wrong index. Expected %d, got %d", solution, answer)
	}
}
