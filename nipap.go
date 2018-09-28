package nipap

import (
	"github.com/kodburn/xmlrpc"
	"encoding/base64"
	"fmt"
	"strings"
)

var userAgent = "go-nipap 0.1"

type Client struct {
	*xmlrpc.Client
	AuthoritativeSource 	string
}

func NewClient(url,username,password string) (error, *Client) {
	headers := make(map[string]string,0)
	headers["Authorization"] = fmt.Sprintf("Basic %s", basicAuth(username,password))

	client, err := xmlrpc.NewClient(url, headers, nil)
	if err != nil {
		return err,nil
	}

	return nil,&Client{client,""}
}

func (c *Client) Run(serviceMethod string, args interface{}, reply interface{}) error {
	auth := make(map[string]string, 0)
	auth["authoritative_source"] = c.AuthoritativeSource

	if args == nil {
		args := make(map[string]interface{},0)
		args["auth"] = auth
	} else {
		t := fmt.Sprintf("%T", args)
		if strings.HasPrefix(t,"map[string]") {
			args.(map[string]interface{})["auth"] = auth
		} else {
			return fmt.Errorf("NIPAP must either recieve nil args or a map[string]")
		}
	}

	return c.Call(serviceMethod, args, reply)
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

