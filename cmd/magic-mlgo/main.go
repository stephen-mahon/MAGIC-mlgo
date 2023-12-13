package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"

	"github.com/stephen-mahon/MAGIC-mlgo/internal/read"
)

func main() {
	fileName := flag.String("f", "input.txt", "input file name")
	flag.Parse()

	names := dataframe.Names(
		"fLength",
		"fWidth",
		"fSize",
		"fConc",
		"fConc1",
		"fAsym",
		"fM3Long",
		"fM3Trans",
		"fAlpha",
		"fDist",
		"class",
	)

	noHeader := dataframe.HasHeader(false)

	df, err := read.File(*fileName, names, noHeader)
	if err != nil {
		log.Fatalf("could not read %s: %v", *fileName, err)
	}

	// Convert "class" column values to 1 if equal to "g", and 0 otherwise
	df = df.Mutate(series.New([]int{}, series.Int, "class"), func(d dataframe.Series) dataframe.Series {
		classCol := d.(*series.Int)
		for i := 0; i < d.Len(); i++ {
			if df.Elem(i, "class").(string) == "g" {
				classCol.Set(i, 1)
			} else {
				classCol.Set(i, 0)
			}
		}
		return classCol
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(df)

}
