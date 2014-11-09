package main

import (
  "os"
  "github.com/codegangsta/cli"
  "time"
  "os/exec"
  "net/http"
  "net/url"
  "launchpad.net/xmlpath"
  "log"
  "strings"
  "strconv"
)

func main() {
  app := cli.NewApp()
  app.Name = "jarvis"
  app.Usage = "I am a nice butler, who can speak and listen"
  app.Action = func(c *cli.Context) {
    println("Hello friend, how can I help you!")
  }
  app.Commands = []cli.Command{
    {
      Name:      "ask",
      ShortName: "a",
      Usage:     "Ask me anything",
      Action: ask,
    },
    {
      Name:      "clock",
      ShortName: "c",
      Usage:     "tells you the current time",
      Action: clock,
    },
    {
      Name:      "date",
      ShortName: "d",
      Usage:     "tells you the current date",
      Action: date,
    },
    {
      Name:      "introduce",
      ShortName: "i",
      Usage:     "Introduces yourself",
      Action: introduce,
    },
    {
      Name:      "timer",
      ShortName: "t",
      Usage:     "set the timer to X seconds, minutes or hours",
      Action: timer,
    },
    {
      Name:      "timer_finish",
      ShortName: "tf",
      Usage:     "informs everyone that the timer has finished",
      Action: timer_finish,
    },
    {
      Name:      "weather",
      ShortName: "w",
      Usage:     "Checks todays weather, e.g. Hamburg, Germany",
      Action: weather,
    },
  }
  app.Run(os.Args)
}

func ask(c *cli.Context) {
  say("Let me look that up for you.")
  // query wolfram alpha
  if len(c.Args()) > 0 {
    url := "http://api.wolframalpha.com/v2/query?input=" + url.QueryEscape(c.Args()[0]) + "&appid=KHJ7LL-XU5JJR9HV9"
    resp, _ := http.Get(url)
    defer resp.Body.Close()
    path := xmlpath.MustCompile("/queryresult/pod[2]/subpod/plaintext")
    root, err := xmlpath.Parse(resp.Body)
    if err != nil {
      log.Fatal(err)
    }
    if value, ok := path.String(root); ok {
      say("Here is what I found: " + value)
    }
  }
}

func clock(c *cli.Context) {
  const layout = "03:04PM"
  t := time.Now()
  say("It is " + t.Format(layout))
}

func date(c *cli.Context) {
  const layout = "January 2, 2006"
  t := time.Now()
  say("Today is " + t.UTC().Format(layout))
}

func introduce(c *cli.Context) {

}

func timer(c *cli.Context) {
  if len(c.Args()) > 0 {
    text := c.Args()[0]
    arr := strings.Split(text, " ")
    count, _ := strconv.Atoi(arr[0])
    unit := arr[1]
    seconds := 0
    switch unit {
      case "seconds":
        seconds = count
      case "minutes":
        seconds = count * 60
      case "hours":
        seconds = count * 60 * 60
    }
    say("I set your timer to fire in " + strconv.Itoa(seconds) + " seconds")
  }
}

func timer_finish(c *cli.Context) {
  say("Your timer has finished.")
}

func weather(c *cli.Context) {

}

func say(text string) {
  exec.Command("say", text).Run()
}
