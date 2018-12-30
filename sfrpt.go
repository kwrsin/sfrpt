package main


func main() {
  options := GetOptions()
  bars, err := GetData(options)
  if err != nil {
    panic(err)
  }
  report(bars, options)
}