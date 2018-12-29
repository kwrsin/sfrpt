package main

import (
  "os"
  "errors"
  "encoding/csv"
)

type TextReader interface {
  textRead(fp *os.File) ([][]string, error)
}

type CSVReader struct {
  options *Options
}

func newCSVReader(options *Options) *CSVReader{
  csvReaer := new(CSVReader)
  csvReaer.options = options
  return csvReaer
}

func (csvReaer *CSVReader) textRead(fp *os.File) ([][]string, error) {
  Println("reading csv file...")
  reader := csv.NewReader(fp)
  reader.Comma = []rune(csvReaer.options.Options.Division_separator)[0]
  reader.LazyQuotes = true
  records, err := reader.ReadAll()
  if err != nil {
    return nil, errors.New("not read csv")
  }

  return records, nil
}

