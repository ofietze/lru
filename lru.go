package main

import (
	"fmt"
	"time"
	"os"
	"strconv"
	"strings"
	"math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyz"
const maxLengthSeq = 12

func main() {
	fmt.Println("Usage: ./lru {cacheSize int} {printIterations y/n} {cacheElem1 string} [...] {cacheElemN string}")
	fmt.Println("Passing only one argument will generate a random accessSeq of lenght [1 .. maxLengthSeq] and solve it.")

	if(len(os.Args[:]) == 2){
		cacheSize, e := strconv.Atoi(os.Args[1])
		if e != nil {
		   fmt.Println(e)
		}
		rand.Seed(time.Now().UTC().UnixNano())
		accessSeq := RandString(rand.Intn(maxLengthSeq))
		fmt.Println("Access Sequence:")
		fmt.Println(strings.Join(accessSeq, ", "))
		answerSeq := Lru(cacheSize, accessSeq, true)
		fmt.Println("Result")
		fmt.Println(strings.Join(answerSeq, ", "))
		return
	}

	if(len(os.Args[:]) < 4){
		fmt.Printf("Wrong number of arguments.")
		return
	}

	cacheSize, e := strconv.Atoi(os.Args[1])
	if e != nil {
	   fmt.Println(e)
	}

	prints := true

	if(os.Args[2] == "n") {
		prints = false
	} else if(os.Args[2] != "y"){
		fmt.Println("Wrong printIterations argument.")
	}

	inputArr := os.Args[3:]

	cacheSeq := Lru(cacheSize, inputArr, prints)
	fmt.Println("Result")
	fmt.Println(strings.Join(cacheSeq, ", "))
}

func RandString(n int) []string {
    s := make([]string, n)
    for i := range s {
        s[i] = string(letterBytes[rand.Intn(len(letterBytes))])
    }
    return s
}

func Lru(cacheSize int, inSeq []string, prints bool) (retSeq []string) {
	lastAccess := make([]int, cacheSize)

	if (len(inSeq) <= cacheSize) {
		retSeq = inSeq
	} else {
		retSeq = make([]string, cacheSize)
		cacheElems := 0

		for i := 0; i < len(inSeq); i++{
			current := inSeq[i]
			cacheIndex := GetIndexOf(retSeq, current)

			if (cacheIndex > -1) {
				// cache hit
				// refresh last access
				lastAccess[cacheIndex] = i
			} else if (i <= cacheElems && cacheElems < cacheSize/* && cacheIndex == -1*/) {
				// cache is not full and cache miss on current data
				retSeq[i] = current
				lastAccess[i] = i;
				cacheElems++
			} else {
				// cache replacement
				oldestAccess := FindLowest(lastAccess);
				retSeq[oldestAccess] = current
				lastAccess[oldestAccess] = i;
			    }

			if (prints) {
				fmt.Println("iteration: ",i)
				fmt.Println(lastAccess)
				fmt.Println(strings.Join(retSeq, " | "))
				}
		}
	}
	return
}

/*
* Finds the first index of {c} in string {in}
*/
func GetIndexOf(in []string, c string) (index int) {
	index = -1

	for i := 0; i < len(in); i++ {
		if (in[i] == c){
			index = i
			return
		}
	}
	return
}

/**
* Finds the index of the lowest element in the given array.
*
* @param accesses the array to search through
* @return index with lowest value
*/
func FindLowest(accesses []int) (lowest int) {
	lowest = 0;

	for i := 0; i < len(accesses); i++ {
	    if (accesses[lowest] > accesses[i]) {
		lowest = i
	    }
	}
	return lowest
}