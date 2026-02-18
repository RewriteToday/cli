package api

import (
	"github.com/RewriteToday/cli/internal/profile"
)

type Client struct {
	Profile string
}

func New() (*Client, error) {
	name, _, err := profile.GetActive()
	if err != nil {
		return nil, err
	}

	return &Client{Profile: name}, nil
}
