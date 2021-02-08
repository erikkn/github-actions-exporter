package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-github/v33/github"
	"golang.org/x/oauth2"
)

// Config contains the config for the Github Client.
type Config struct {
	Token        string
	Organization string
	TimeOut      time.Duration
}

// Client wraps the Github Client and some other metadata together.
type Client struct {
	*Config
	Client *github.Client
}

//func NewClient(config) *github.Client <square bracket :P>

// CreateClient creates a Github client
func (c *Config) CreateClient() (*Client, error) {
	if c.Token == "" {
		return nil, fmt.Errorf("unable to create a Github Client: No token was provided")
	}

	ctx, _ := context.WithTimeout(context.Background(), c.TimeOut)

	httpClient := oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: c.Token}))
	ghClient := github.NewClient(httpClient)

	return &Client{
		Client: ghClient,
		Config: ghConfig,
	}, nil
}
