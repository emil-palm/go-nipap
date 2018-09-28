package nipap


import (
	"time"
	"github.com/fatih/structs"
	"fmt"
)

type PrefixType string

const (
        PrefixTypeAssignment                  PrefixType	= "assignment"
        PrefixTypeReservation					= "reservation"
)

type Prefix struct {
	Family int `xmlrpc:"family" structs:",omitempty"`
	Vrf string `xmlrpc:"vrf" structs:",omitempty"`
	VrfId int `xmlrpc:"vrf_id" structs:",omitempty"`
	Prefix string `xmlrpc:"prefix" structs:"prefix,omitempty"`
	DisplayPrefix string `xmlrpc:"display_prefix" structs:",omitempty"`
	Description string `xmlrpc:"description" structs:"description,omitempty"`
	Comment string `xmlrpc:"comment" structs:",omitempty"`
	Node string `xmlrpc:"node" structs:",omitempty"`
	Pool string `xmlrpc:"pool" structs:",omitempty"`
	Type PrefixType `xmlrpc:"type" structs:"type,omitempty"`
	Indent int `xmlrpc:"indent" structs:",omitempty"`
	Id int `xmlrpc:"id" structs:",omitempty"`
	Country string`xmlrpc:"country" structs:",omitempty"`
	ExternalKey string `xmlrpc:"external_key" structs:",omitempty"`
	OrderId int `xmlrpc:"order_id" structs:",omitempty"`
	CustomerId int `xmlrpc:"customer_id" structs:",omitempty"`
	AuthoritativeSource string `xmlrpc:"authoritative_source" structs:",omitempty"`
	AlarmPriority string `xmlrpc:"alarm_priority" structs:",omitempty"`
	Monitor bool `xmlrpc:"monitor" structs:",omitempty"`
	Display bool `xmlrpc:"display" structs:",omitempty"`
	Match bool `xmlrpc:"match" structs:",omitempty"`
	Children int `xmlrpc:"children" structs:",omitempty"`
	Vlan int `xmlrpc:"vlan" structs:",omitempty"`
	Added time.Time `xmlrpc:"added" structs:",omitempty"`
	LastModified time.Time `xmlrpc:"last_modified" structs:",omitempty"`
	TotalAddresses string `xmlrpc:"total_addresses" structs:",omitempty"`
	UsedAddreses string `xmlrpc:"used_addresses" structs:",omitempty"`
	FreeAddreses string `xmlrpc:"free_addresses" structs:",omitempty"`
	Status string `xmlrpc:"status" structs:",omitempty"`
	Expires string `xmlrpc:"expires" structs:",omitempty"`
}

func (p Prefix) stripped () map[string]interface{} {
	return structs.Map(p)
}



func (client *Client) addPrefix(prefix Prefix, spec map[string]interface{}) (error, Prefix) {
	args := make(map[string]interface{},0)
	if spec != nil {
		args["args"] = spec
	}

	args["attr"] = prefix.stripped()

	var prefixResp Prefix

	err := client.Run("add_prefix",args, &prefixResp)
	if err != nil {
		return err, Prefix{}
	}

	return nil, prefixResp

}
func (client *Client) AddPrefix(prefix Prefix) (error, Prefix) {
	return client.addPrefix(prefix, nil)
}

func (client *Client) AddPrefixFromPrefix(newPrefix, parentPrefix Prefix, prefixLength int) (error, Prefix) {
	if parentPrefix.Prefix == "" {
		return fmt.Errorf("Cannot add a prefix within parentPrefix when parentPrefix.prefix is empty"), Prefix{}
	}
	if prefixLength <= 0 || prefixLength > 128 {
		return fmt.Errorf("Cannot add a prefix without a PrefixLength set"), Prefix{}
	}

	spec := NewSpec()
	spec.Set("from-prefix", []string{parentPrefix.Prefix})
	spec.Set("prefix_length", prefixLength)

	return client.addPrefix(newPrefix,spec)
}

func (client *Client) ListPrefix(spec map[string]string) (error,[]Prefix) {
	args := make(map[string]interface{},0)
	if spec == nil {
		spec = make(map[string]string)
	}
	args["prefix"] = spec

	prefixes := make([]Prefix, 0)
	err := client.Run("list_prefix",args, &prefixes)
	if err != nil {
		return err, nil
	}
	return nil, prefixes
}

func (client *Client) PrefixSmartSearch(query string, options map[string]interface{}) (error,[]Prefix) {
	args := make(map[string]interface{},0)
	args["query_string"] = query

	if options != nil {
		args["search_options"] = options
	}

	response := struct{
		Error bool `xmlrpc:"error"`
		Result []Prefix `xmlrpc:"result"`
	}{}

	err := client.Run("smart_search_prefix", args, &response)
	return err, response.Result
}
