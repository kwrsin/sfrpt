package main

import (
  "os"
  "errors"
  "strconv"
)

const (
  TITLE_PAIR = "PAIR"
  TITLE_DATE = "DATE"
  TITLE_TIME = "TIME"
  TITLE_OPEN = "OPEN"
  TITLE_HIGH = "HIGH"
  TITLE_LOW = "LOW"
  TITLE_CLOSE = "CLOSE"
  TITLE_VOL = "VOL"
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

func getIndex(search string, titles []string, defaultIndex int) int {
  for i, title := range titles {
    switch {
    case search == title:
      return i
    }
  }
  return defaultIndex
}

func toBar(record []string, titles []string) (*Bar, error) {
  var iPair, iDate, iTime, iOpen, iHigh, iLow, iClose, iVol int

  iPair = getIndex(TITLE_PAIR, titles, 0)
  iDate = getIndex(TITLE_DATE, titles, 1)
  iTime = getIndex(TITLE_TIME, titles, 2)
  iOpen = getIndex(TITLE_OPEN, titles, 3)
  iHigh = getIndex(TITLE_HIGH, titles, 4)
  iLow = getIndex(TITLE_LOW, titles, 5)
  iClose = getIndex(TITLE_CLOSE, titles, 6)
  iVol = getIndex(TITLE_VOL, titles, 7)

  bar := new(Bar)
  recordCount := len(record)

  if iPair >= 0 && iPair < recordCount {bar.symbol = record[iPair]}
  if iDate >= 0 && iDate < recordCount {bar.date = record[iDate]}
  if iTime >= 0 && iTime < recordCount {bar.time = record[iTime]}
  if iOpen >= 0 && iOpen < recordCount {
    price, err := strconv.ParseFloat(record[iOpen], 64)
    if err != nil {
      return nil, errors.New("converting error[OPEN]")
    }
    bar.open = price
  }

  if iHigh >= 0 && iHigh < recordCount {
    price, err := strconv.ParseFloat(record[iHigh], 64)
    if err != nil {
      return nil, errors.New("converting error[HIGH]")
    }
    bar.high = price
  }

  if iLow >= 0 && iLow < recordCount {
    price, err := strconv.ParseFloat(record[iLow], 64)
    if err != nil {
      return nil, errors.New("converting error[LOW]")
    }
    bar.low = price
  }

  if iClose >= 0 && iClose < recordCount {
    price, err := strconv.ParseFloat(record[iClose], 64)
    if err != nil {
      return nil, errors.New("converting error[CLOSE]")
    }
    bar.close = price
  }

  if iVol >= 0 && iVol < recordCount {
    price, err := strconv.ParseFloat(record[iVol], 64)
    if err != nil {
      return nil, errors.New("converting error[VOL]")
    }
    bar.vol = price
  }

  return bar, nil
}

func toBars(records [][]string, titles []string) []*Bar {
  Println("converting csv records to bars...")
  bars := make([]*Bar, 0, len(records))
  for i, record := range records {
    bar, err := toBar(record, titles)
    if err != nil {
      Printf("Skiped %d line, %v\n", i, err)
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
  bars := toBars(records, options.Options.Input)

  return bars, nil
}
