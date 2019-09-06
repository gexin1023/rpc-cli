package main

import (
	"encoding/json"
	"fmt"
	"github.com/powerman/rpc-codec/jsonrpc2"
	"github.com/urfave/cli"
)

var (

	callCmd = cli.Command {
		Name:"call",
		Flags:rpcFlags,
		Action:post,
	}

	subcmds = []cli.Command{
		callCmd,
	}
)

var Version = "2.0"

func post (c *cli.Context)  error{

	url := c.String("url")
	method := c.String("method")
	paramJson := c.String("param")

	var params interface{}
	if err := json.Unmarshal([]byte(paramJson), &params); err !=nil{
		panic(err)
	}

	client := jsonrpc2.NewHTTPClient(url)

	var result interface{}

	if err := client.Call(method, params, &result); err != nil{
		panic(err)
	}

	fmt.Println(result)

	return nil
}