package main

import (
  "testing"
)

func TestGetOptions(t *testing.T) {
  options := GetOptions()
  if options == nil {
    t.Fatal("Can't get options")
  }
  if options.Options.Division_separator == "" ||
     options.Options.Join_separator == "" {
      t.Fatalf("members are empty, %v", options)
  }
}