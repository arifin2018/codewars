package Go

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//https://www.codewars.com/kata/5a34af40e1ce0eb1f5000036/train/go

func CSV_representation_of_array(data [][]int) string {
	var records [][]string
	var convertData []string
	var result string

	for i := 0; i < len(data); i++ {
		convertData = nil
		for _, v := range data[i] {
			convertData = append(convertData, strconv.Itoa(v))
		}
		records = append(records, convertData)
	}

	b := new(bytes.Buffer)
	w := csv.NewWriter(b)
	w.WriteAll(records) // calls Flush internally

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
	s := b.String()
	result = s[:len(s)-1]

	return result
}

func CSV_representation_of_array1(array [][]int) string {
	array2 := []string{}
	for _, a := range array {
		k := []string{}
		for _, n := range a {
			k = append(k, fmt.Sprint(n))
		}
		array2 = append(array2, strings.Join(k, ","))
	}
	return strings.Join(array2, "\n")
}
