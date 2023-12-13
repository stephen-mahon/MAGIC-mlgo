package read

import (
	"os"

	"github.com/go-gota/gota/dataframe"
)

func File(filename string, names dataframe.LoadOption, noHeader dataframe.LoadOption) (dataframe.DataFrame, error) {

	f, err := os.Open(filename)
	if err != nil {
		return dataframe.DataFrame{}, err
	}
	defer f.Close()

	return dataframe.ReadCSV(f, names, noHeader), nil

}
