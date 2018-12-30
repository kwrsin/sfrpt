// reporter.go
package main

import (
  "fmt"
  "strings"
  "strconv"
)

func enumerate(num int, pBars *[]*Bar, options *Options) {
  myBar := (*pBars)[num]
  barCount := len((*pBars))
  var sign int = 0
  var maxPrice float64 = myBar.close
  var minPrice float64 = myBar.close
  var bar_counter int = 0
  var pips float64
  var isFind bool = false
  var isBack bool = false

  for i := num + 1; i < barCount; i++ {
    bar_counter++
    nextBar := (*pBars)[i]
    sub := nextBar.close - myBar.close
    if sub > 0 {
      if sign == 0 {sign++}
      if maxPrice < nextBar.close {
        maxPrice = nextBar.close
      }
    } else if(sub < 0) {
      if sign == 0 {sign--}
      if minPrice > nextBar.close {
        minPrice = nextBar.close
      }
    }
    if sign > 0 && sub <= 0 {
      pips = (maxPrice - myBar.close) * 1000
      isBack = true
    } else if sign < 0 && sub >= 0 {
      pips = (minPrice - myBar.close) * 1000
      isBack = true
    }
    if(isBack == true) {
      record := toArray(myBar, pips, bar_counter)
      fmt.Println(strings.Join(record, options.Options.Join_separator))
      isFind = true
      break
    }
  }
  if isFind == false {
    record := toArray(myBar, pips, -1)
    fmt.Println(strings.Join(record, options.Options.Join_separator))
  }
}

func toArray(pBar *Bar, pips float64, bar_counter int) []string {
  var record []string
  if pBar.symbol != "" {
    record = append(record, pBar.symbol)
  }
  if pBar.date != "" {
    record = append(record, pBar.date)
  }
  if pBar.time != "" {
    record = append(record, pBar.time)
  }
  if pBar.open != 0 {
    record = append(record, strconv.FormatFloat(pBar.open, 'f', 3, 64))
  }
  if pBar.high != 0 {
    record = append(record, strconv.FormatFloat(pBar.high, 'f', 3, 64))
  }
  if pBar.low != 0 {
    record = append(record, strconv.FormatFloat(pBar.low, 'f', 3, 64))
  }
  if pBar.close != 0 {
    record = append(record, strconv.FormatFloat(pBar.close, 'f', 3, 64))
  }
  if pBar.vol != 0 {
    record = append(record, strconv.FormatFloat(pBar.vol, 'g', 4, 64))
  }
  record = append(record, strconv.FormatFloat(pips, 'g', 4, 64))
  record = append(record, strconv.Itoa(bar_counter))

  return record
}

func report(bars []*Bar, options *Options) {
  barCount := len(bars)
  for i := 0; i < barCount - 1; i++ {
    enumerate(i, &bars, options)
  }
  record := toArray(bars[barCount -1], 0, -1)
  fmt.Println(strings.Join(record, options.Options.Join_separator))

}