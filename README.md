# Github Actions Prometheus Exporter
Yet another quick and dirty tGithub Actions exporter for Prometheus.

## Getting started
Download or build the binary yourself. Either build the Docker container yourself or run the binary directly straight from the command line.

Available arguments:

```
  -organization string
    	The name of your Github Organization
  -path string
    	HTTP Path that Prometheus will scrape, e.g. /metrics (default "/metrics")
  -ratelimit duration
    	The delay (in minutes) the exporter should keep between the API calls (default 5m0s)
  -timeout duration
    	After this time we cut the connection with Github. (default 5s)
  -token string
    	Your private/organization Github Token to connect to Github
```

## Installing

```go
go get -u github.com/erikkn/github-actions-exporter
```

## Building

```
make build
```

### Running

```
./build/bin/github-actions-exporter
```

## Todo
A lot, but most importantly verbose output of the net.HTTP serverMux.
