package main

import (
  "encoding/json"
  "io/ioutil"
  "os"
)

const (
  SETTINGS = "settings.json"
  DEFAULT_DIVISION_SEPARATOR = ","
  DEFAULT_JOIN_SEPARATOR = ","
  DEFAULT_DATA_PATH = ""
)

type Options struct {
  Options Option `json:"options"`
}

type Option struct {
  Input []string `json:"input"`
  Division_separator string `json:"division_separator"`
  Join_separator string `json:"join_separator"`
  Data_path string `json:"data_path"`
}

func GetOptions() (*Options) {
  var options Options
  json_file, err := os.Open(SETTINGS)
  defer json_file.Close()
  if err != nil {
    Println("Not Found ", SETTINGS)
  } else {
    bytes, _ := ioutil.ReadAll(json_file)
    json.Unmarshal([]byte(bytes), &options)
  }

  setDefault(&options)

  return &options
}

func getValueString(val string, def string) string {
  if val == "" {
    return def
  }
  return val
}

func setDefault(options *Options)  {
  options.Options.Division_separator =
    getValueString(
      options.Options.Division_separator, DEFAULT_DIVISION_SEPARATOR)

  options.Options.Join_separator =
    getValueString(
      options.Options.Join_separator, DEFAULT_JOIN_SEPARATOR)

  options.Options.Data_path =
    getValueString(
      options.Options.Data_path, DEFAULT_DATA_PATH)
}