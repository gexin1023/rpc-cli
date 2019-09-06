package main

import "github.com/urfave/cli"

var (
	urlFlag = cli.StringFlag{
		Name:"url, u",
		Usage:"url",
		Value : "http://127.0.0.1:6791",
	}

	methodFlag = cli.StringFlag{
		Name:"method, m",
		Usage:"method to call",
	}

	paramFlag = cli.StringFlag{
		Name: "param, p",
		Usage:"param to call",
	}

	rpcFlags = []cli.Flag{
		urlFlag,
		methodFlag,
		paramFlag,
	}
)
