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
  "runtime"
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
    text := strings.Join(c.Args(), " ")
    url := "http://api.wolframalpha.com/v2/query?input=" + url.QueryEscape(text) + "&appid=KHJ7LL-XU5JJR9HV9"
    resp, _ := http.Get(url)
    defer resp.Body.Close()
    path := xmlpath.MustCompile("/queryresult/pod[2]/subpod/plaintext")
    root, err := xmlpath.Parse(resp.Body)
    if err != nil {
      log.Fatal(err)
    }
    if value, ok := path.String(root); ok {
      parts := strings.SplitAfter(value, "(")
      say("Here is what I found: " + parts[0])
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
  say("Hello. My name is Jarvis.")
  say("I can answer questions, set timers and get the current weather.")
  say("And I am here to win the contest, of course.")
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

	time.Sleep(time.Duration(seconds) * time.Second)

	say("Your timer has finished.")
  }
}

func weather(c *cli.Context) {
	text := strings.Join(c.Args(), " ")
	query := strings.Join(c.Args(), "_")
    url := "http://api.wunderground.com/api/e0930f07065b7998/conditions/q/" + query + ".xml"
	response, _ := http.Get(url)
	defer response.Body.Close()

	response2, _ := http.Get(url)
	defer response2.Body.Close()

	body := response.Body
	body2 := response2.Body

    temp_path := xmlpath.MustCompile("/response/current_observation/temp_c")
    temp_root, _ := xmlpath.Parse(body)
	temperature, _ := temp_path.String(temp_root)

	desc_path := xmlpath.MustCompile("/response/current_observation/weather")
    desc_root, _ := xmlpath.Parse(body2)
	description, _ := desc_path.String(desc_root)

   	string := "It has " + temperature + " degrees in " + text + ", the weather is " + description
	say(string)
}

func say(text string) {
  os := runtime.GOOS
  switch os {
    case "linux":
      exec.Command("tts", text).Run()
    case "darwin":
      exec.Command("say", text).Run()
    default:
      println(text)
  }
}
