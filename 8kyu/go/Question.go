package IDENT

import (
	"encoding/csv"
	"log"
	"os"
)

func CSV_representation_of_array(data [][]string) {
	w := csv.NewWriter(os.Stdout)

	for _, v := range data {
		if err := w.Write(v); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
		w.UseCRLF = true
	}
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

}
