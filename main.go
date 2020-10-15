package main

import (
  "os"
  "fmt"
  "time"
)

const TIME_FORMAT = "15:04";

func main() {
  args := os.Args[1:];

  startTime, _ := time.Parse(TIME_FORMAT, args[0]);

  fmt.Printf("startTime: %s\n", startTime);
  fmt.Printf("endTime: %s\n", startTime.Add(8 * time.Hour));
  fmt.Println(args);
}
