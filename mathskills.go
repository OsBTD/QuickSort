package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	Data := Read()
	Sorted := QuickSort(Data)
	resfile, err := os.Create("ResultName.txt")
	if err != nil {
		log.Fatal("Error creating file : ", err)
	}
	defer resfile.Close()
	var s string
	for _, v := range Sorted {
		s += strconv.FormatFloat(v, 'f', 0, 64) + " "
	}
	_, err = io.WriteString(resfile, s)
	if err != nil {
		log.Fatal("Error writing to file : ", err)
	}
	fmt.Println("File processed successfully :) ")
}

func Read() []float64 {
	var Population []float64

	content, err := os.ReadFile("datasettest.txt")
	if err != nil {
		log.Fatal("couldn't read file")
	}

	Split := strings.Split(strings.TrimSpace(string(content)), " ")

	for _, v := range Split {
		s, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Print("Parse failed", string(v), "in this value")
		} else {
			Population = append(Population, s)
		}
	}
	return Population
}

func QuickSort(Population []float64) []float64 {
	var before, after, pivotlist []float64

	if len(Population) < 2 {
		return Population
	}
	first := Population[0]
	middle := Population[(len(Population)-1)/2]
	last := Population[len(Population)-1]
	var Pivot float64

	if (first > middle) && (first < last) {
		Pivot = first
	} else if (middle > first) && (middle < last) {
		Pivot = middle
	} else {
		Pivot = last
	}

	for _, value := range Population {
		if value < Pivot {
			before = append(before, value)
		} else if value > Pivot {
			after = append(after, value)
		} else {
			pivotlist = append(pivotlist, value)
		}
	}
	var s []float64
	s = append(s, QuickSort(before)...)
	s = append(s, pivotlist...)
	s = append(s, QuickSort(after)...)

	return s
}
