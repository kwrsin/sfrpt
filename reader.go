package main

import (
  "os"
  "errors"
  "strconv"
)

type Bar struct {
  symbol string
  date string
  time string
  open float64
  high float64
  low float64
  close float64
  vol float64
}

func readData(data_path string, textReader TextReader) ([][]string, error) {
  info, err := os.Stdin.Stat()
  var fp *os.File
  if err != nil || info.Size() <= 0 {
    if data_path != "" {
      fp, err = os.Open(data_path)
      if err != nil {
        return nil, err
      }
      defer fp.Close()
    }
  } else {
    if info.Mode() & os.ModeCharDevice != 0 {
      return nil, errors.New("device error")
    }
    fp = os.Stdin
  }
  data, err := textReader.textRead(fp)
  if err != nil {
    return nil, err
  }
  return data, nil
}

func toBar(record []string) (*Bar, error) {
  bar := new(Bar)
  bar.symbol = record[0]
  bar.date = record[1]
  bar.time = record[2]
  open, err_open := strconv.ParseFloat(record[3], 64)
  bar.open = open

  high, err_high := strconv.ParseFloat(record[4], 64)
  bar.high = high

  low, err_low := strconv.ParseFloat(record[5], 64)
  bar.low = low

  clos, err_close := strconv.ParseFloat(record[6], 64)
  bar.close = clos

  vol, err_vol := strconv.ParseFloat(record[7], 64)
  bar.vol = vol

  if err_open != nil ||
     err_high != nil ||
     err_low != nil ||
     err_close != nil ||
     err_vol != nil {
    return nil, errors.New("bar converting error")
  }

  return bar, nil
}

func toBars(records [][]string) []*Bar {
  Println("converting csv records to bars...")
  bars := make([]*Bar, 0, len(records))
  for i, record := range records {
    bar, err := toBar(record)
    if err != nil {
      Printf("%d line, %v\n", i, err)
      continue
    }
    bars = append(bars, bar)
  }

  return bars
}

func GetData(options *Options) ([]*Bar, error) {
  // ReadFile
  records, err := readData(
    options.Options.Data_path, newCSVReader(options))
  if err !=  nil {
    return nil, errors.New("csv reading error")
  }
  bars := toBars(records)

  return bars, nil
}
