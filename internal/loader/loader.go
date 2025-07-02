package loader

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"tickets/internal/domain"
)

func NewLoaderTicketCSV(filePath string) *LoaderTicketCSV {
	return &LoaderTicketCSV{
		filePath: filePath,
	}
}

type LoaderTicketCSV struct {
	filePath string
}

func (l *LoaderTicketCSV) Load() (t map[int]domain.TicketAttributes, err error) {
	f, err := os.Open(l.filePath)
	if err != nil {
		err = fmt.Errorf("error opening file: %v", err)
		return
	}
	defer f.Close()

	r := csv.NewReader(f)

	t = make(map[int]domain.TicketAttributes)
	for {
		record, err := r.Read()

		if err != nil {
			if err == io.EOF {
				break
			}

			return t, err

		}

		idStr := record[0]
		id, _ := strconv.Atoi(idStr)
		price, _ := strconv.ParseFloat(record[5], 64)
		ticket := domain.TicketAttributes{
			Name:    record[1],
			Email:   record[2],
			Country: record[3],
			Hour:    record[4],
			Price:   price,
		}

		t[id] = ticket
	}

	return
}
