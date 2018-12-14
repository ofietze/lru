package main

import (
	"fmt"
	"time"
	"os"
	"strconv"
	"strings"
	"math/rand"
)

const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	fmt.Println("Usage 1 Solve: ./lru {cacheSize int} {printIterations y/n} {cacheElem1 string} [...] {cacheElemN string}")
	fmt.Println("Usage 2 Generate: ./lru {cacheSize int} {numReplacements int}")
	fmt.Println("Passing only two arguments will generate a random access sequence of length {cacheSize} + {numReplacements} and solve it.")

	// Check which usage case was chosen
	if(len(os.Args[:]) == 3){
		cacheSize, e := strconv.Atoi(os.Args[1])
		if e != nil {
		   fmt.Println(e)
		}

		numReplacements, err := strconv.Atoi(os.Args[2])
		if err != nil {
		   fmt.Println(err)
		}
		rand.Seed(time.Now().UTC().UnixNano())
		// choose a random sequence length from [cacheSize, cacheSize + numReplacments]
		accessSeq := RandString(cacheSize + numReplacements)
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
				// uncomment for more output of algo
			//	fmt.Println(lastAccess)
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
 Finds the index of the lowest element in the given array.
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
