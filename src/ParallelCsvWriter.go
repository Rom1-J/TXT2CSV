package main

import (
	"encoding/csv"
	"os"
	"sync"
)

type CsvWriter struct {
	mutex     *sync.Mutex
	csvWriter *csv.Writer
}

func ParallelCsvWriter(fileName string) (*CsvWriter, error) {
	csvFile, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	w := csv.NewWriter(csvFile)
	return &CsvWriter{csvWriter: w, mutex: &sync.Mutex{}}, nil
}

func (w *CsvWriter) Write(row []string) {
	w.mutex.Lock()
	err := w.csvWriter.Write(row)
	if err != nil {
		return
	}
	w.mutex.Unlock()
}

func (w *CsvWriter) WriteAll(row [][]string) {
	w.mutex.Lock()
	err := w.csvWriter.WriteAll(row)
	if err != nil {
		return
	}
	w.mutex.Unlock()
}

func (w *CsvWriter) Flush() {
	w.mutex.Lock()
	w.csvWriter.Flush()
	w.mutex.Unlock()
}
