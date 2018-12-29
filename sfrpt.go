package main


func main() {
  options := GetOptions()
  bars, err := GetData(options)
  if err != nil {
    panic(err)
  }
  Printf("%v", bars)
  //TODO: goroutine１行処理
}