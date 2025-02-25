package csv

import (
	"encoding/csv"
	"io"
)

func Parser(fileIo io.Reader) (data [][]string, _ int, err error) {
	r := csv.NewReader(fileIo)
	rows, err := r.ReadAll() // csvを一度に全て読み込む
	if err != nil {
		return nil, 0, err
	}

	return rows, len(rows), nil
}
