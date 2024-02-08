package Go

import (
	"bytes"
	"encoding/csv"
	"log"
	"strconv"
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
