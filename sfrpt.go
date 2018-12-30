package main


func main() {
  options := GetOptions()
  bars, err := GetData(options)
  if err != nil {
    panic(err)
  }
  report(bars, options)
  //TODO: goroutine１行処理 reporter
  //TODO: View
}