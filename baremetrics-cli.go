/**
 * Baremetrics CLI Client
 *
 * I'm a go newb - so sorry about this in advance. I'd love
 * feedback on what I could do better though. :-)
 */

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const endpoint_url = string("https://dashboard.baremetrics.io/")
const cli_version = string("0.0.1")

func main() {
	cookie := flag.String("cookie", "", "Your baremetrics cookie")
	flag.Parse()

	if *cookie == "" {
		exit_and_print_usage("You must specify your Baremetrics Cookie.")
	}

	start_date, end_date := get_time()
	fmt.Println(fetch_http("stats/mrr.json?start_date="+start_date+"&end_date="+end_date, *cookie))
}

func get_time() (string, string) {
	base_time := time.Now()
	start_date := base_time.Add(-30 * 24 * time.Hour).Format("2006-01-02")
	end_date := base_time.Format("2006-01-02")

	return start_date, end_date
}

func fetch_http(url string, cookie string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", endpoint_url+url, nil)

	req.Header.Set("Cookie", cookie)
	req.Header.Add("User-Agent", "BaremetricsCli/"+cli_version+" (https://github.com/ukd1/baremetrics-cli)")

	resp, err := client.Do(req)
	if err != nil {
		// handle error
		fmt.Println("Could not fetch JSON:")
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		content, _ := ioutil.ReadAll(resp.Body)
		return string(content[:])
	} else {
		fmt.Println("Error: " + resp.Status)
		os.Exit(1)
		return ""
	}
}

func print_usage() {
	fmt.Println("Usage:")
	flag.PrintDefaults()
	fmt.Println("\nExamples:")
	fmt.Println("baremetrics-cli -cookie '__cfduid=dd5e0bb94411y55be2......'")
	fmt.Println("")
}

func exit_and_print_usage(message string) {
	fmt.Println("Error:", message, "\n")
	print_usage()
	os.Exit(1)
}
