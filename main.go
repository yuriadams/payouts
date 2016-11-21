package main

import (
  "os"
  "encoding/json"
  "github.com/yuriadams/payout/models"
  "github.com/yuriadams/payout/services"
  // "github.com/DavidHuie/quartz/go/quartz"
)

type Finalizer struct{}

type Args struct {
    GameId string
    Rankings string
}

type Response struct {
  // Results []models.Result
}

func (t *Finalizer) Finalize(args Args, response *Response) error {
    *response = Response{}

    bytes := []byte(args.Rankings)
    var rankings []models.Result
    json.Unmarshal(bytes, &rankings)
    services.Finalize(args.GameId, rankings)
    return nil
}

func main() {
  // finalizer := &Finalizer{}
  // quartz.RegisterName("finalizer", finalizer)
  // quartz.Start()

  args := os.Args[1:]

  bytes := []byte(args[1])
  var rankings []models.Result
  json.Unmarshal(bytes, &rankings)

  services.Finalize(args[0], rankings)
}
