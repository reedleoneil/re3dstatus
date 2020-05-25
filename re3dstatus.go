package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  emojis := []string {"👹", "👺", "🍱", "🍜", "🍙", "🍥", "🌸", "🌊", "🎏", "🎋", "🎎", "🗼", "🏯", "🗻", "🏮", "🔴", "🗾", "🍣", "🍵", "🍡", "🍘", "🍢"}
  s := rand.NewSource(time.Now().Unix())
  r := rand.New(s)
  rnd := r.Intn(len(emojis))
  fmt.Println(time.Now().Format("Mon Jan 2 15:04:05 MST 2006") + " " + emojis[rnd] + "(日本)")
}
