package main

import "net"

type Information struct {
}

type Provider struct {
	server net.IPAddr
}

func (provider *Provider) GetInfo() (Information, error) {
	var info Information
	var err error
	info = Information{}
	err = nil
	/* Connect to other server */
	/* Very complex method with high cyclomatic complexity */
	return info, err
}
