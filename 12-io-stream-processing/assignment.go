/* using the bufio.Scanner to read lines from the source */

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	var evenSum, oddSum int
	splitter(&evenSum, &oddSum, "data1.dat")
	splitter(&evenSum, &oddSum, "data2.dat")
	saveResults(&evenSum, &oddSum)
}

func splitter(evenSum *int, oddSum *int, fileName string) {
	fileHandle, fileError := os.Open(fileName)
	if fileError != nil {
		log.Fatalln(fileError)
	}
	defer fileHandle.Close()
	scanner := bufio.NewScanner(fileHandle)
	for scanner.Scan() {
		str := scanner.Text()
		intVar, err := strconv.Atoi(str)
		if err != nil {
			log.Fatalln(err)
		}
		if intVar%2 == 0 {
			fmt.Println("even", intVar)
			*evenSum += intVar
		} else {
			fmt.Println("odd", intVar)
			*oddSum += intVar
		}
	}
}

func saveResults(evenSum, oddSum *int) {
	err2 := ioutil.WriteFile("result.txt", []byte(fmt.Sprintf("Even Total: %d\nOdd Total: %d", *evenSum, *oddSum)), 0x777)
	if err2 != nil {
		log.Fatalln(err2)
	}
	fmt.Println("file copied")
}
