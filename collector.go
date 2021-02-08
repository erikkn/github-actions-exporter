package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/go-github/v33/github"
	"github.com/prometheus/client_golang/prometheus"
)

type githubRunners struct {
	*github.Runners
}

var orgRunnerStatus = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "github_organization_runner_status",
	Help: "Status of the self-hosted Github runners in the entire organization.",
},
	[]string{"name", "id"},
)

func collectMetrics(t time.Duration) {
	for {

		orgRunners, err := listOfOrgRunners()
		if err != nil {
			log.Fatalf("error setting listOfOrgRunners metric: %s", err)
		}

		err = orgRunners.setRunnerStatusMetric()
		if err != nil {
			log.Fatalf("error setting the runner status metrics:", err)
		}

		time.Sleep(t)
	}
}

func init() {
	prometheus.MustRegister(orgRunnerStatus)
}

// listOfOrgRunners returns the runners in an organizations
func listOfOrgRunners() (*githubRunners, error) {
	ctx := context.Background()

	r, _, err := ghClient.Client.Actions.ListOrganizationRunners(ctx, ghClient.Organization, nil)
	if err != nil {
		return nil, fmt.Errorf("error retrieving runners: %s", err)
	}

	return &githubRunners{
		Runners: r,
	}, nil
}

// setRunnerStatusMetric sets the status of the runners; 0=offline, 1=idle/online, 2=active/busy.
func (r *githubRunners) setRunnerStatusMetric() error {

	for _, v := range r.Runners.Runners {
		if *v.Status == "idle" {
			orgRunnerStatus.WithLabelValues(*v.Name, string(*v.ID)).Set(1)
		} else if *v.Status == "offline" {
			orgRunnerStatus.WithLabelValues(*v.Name, string(*v.ID)).Set(0)
		} else if *v.Status == "active" {
			orgRunnerStatus.WithLabelValues(*v.Name, string(*v.ID)).Set(2)
		} else {
			return fmt.Errorf("runner %s is in an unknown state: %s", *v.Name, *v.Status)
		}
	}

	return nil
}
