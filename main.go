package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	ghConfig *Config
	ghClient *Client

	metricsPath    = flag.String("path", "/metrics", "HTTP Path that Prometheus will scrape, e.g. /metrics")
	githubToken    = flag.String("token", "", "Your private/organization Github Token to connect to Github")
	rateLimit      = flag.Duration("ratelimit", 5*time.Minute, "The delay (in minutes) the exporter should keep between the API calls")
	requestTimeOut = flag.Duration("timeout", 5*time.Second, "After this time we cut the connection with Github.")
	organization   = flag.String("organization", "", "The name of your Github Organization")
)

func rootHandler(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte(`
<html>
<head><title>GitHub Actions Exporter</title></head>
<body>
<h1>GitHub Actions Exporter</h1>
</body>
</html>`))
}

func main() {
	flag.Parse()

	if len(os.Args) < 2 {
		flag.PrintDefaults()
		log.Fatalf("\n\n No user-input provided; exiting..")
	}

	if *organization == "" {
		log.Fatalln("no organization name was provided, exiting")
	}

	ghConfig = &Config{
		Token:        *githubToken,
		Organization: *organization,
	}

	var err error
	ghClient, err = ghConfig.CreateClient()
	if err != nil {
		log.Fatalf("error creating client: %s", err)
	}

	go collectMetrics(*rateLimit)

	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.Handle(*metricsPath, promhttp.Handler())

	http.ListenAndServe(":9870", mux)
}
