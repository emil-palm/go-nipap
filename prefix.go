package nipap


import (
	"time"
	"github.com/fatih/structs"
	sq "github.com/mrevilme/go-nipap/search_query"
	"fmt"
)


type PrefixType string

const (
        PrefixTypeAssignment                  PrefixType	= "assignment"
        PrefixTypeReservation					= "reservation"
        PrefixTypeHost						= "host"
)

type Prefix struct {
	Family int `xmlrpc:"family" structs:"family,omitempty"`
	Vrf string `xmlrpc:"vrf" structs:"vrf,omitempty"`
	VrfId int `xmlrpc:"vrf_id" structs:"vrf_id,omitempty"`
	Prefix string `xmlrpc:"prefix" structs:"prefix,omitempty"`
	DisplayPrefix string `xmlrpc:"display_prefix" structs:"display_prefix,omitempty"`
	Description string `xmlrpc:"description" structs:"description,omitempty"`
	Comment string `xmlrpc:"comment" structs:"comment,omitempty"`
	Node string `xmlrpc:"node" structs:"node,omitempty"`
	Pool string `xmlrpc:"pool" structs:"pool,omitempty"`
	Type PrefixType `xmlrpc:"type" structs:"type,omitempty"`
	Indent int `xmlrpc:"indent" structs:"indent,omitempty"`
	Id int `xmlrpc:"id" structs:"id,omitempty"`
	Country string`xmlrpc:"country" structs:"country,omitempty"`
	ExternalKey string `xmlrpc:"external_key" structs:"external_key,omitempty"`
	OrderId int `xmlrpc:"order_id" structs:"order_id,omitempty"`
	CustomerId int `xmlrpc:"customer_id" structs:"customer_id,omitempty"`
	AuthoritativeSource string `xmlrpc:"authoritative_source" structs:"authoritative_source,omitempty"`
	AlarmPriority string `xmlrpc:"alarm_priority" structs:"alarm_priority,omitempty"`
	Monitor bool `xmlrpc:"monitor" structs:"monitor,omitempty"`
	Display bool `xmlrpc:"display" structs:"display,omitempty"`
	Match bool `xmlrpc:"match" structs:"match,omitempty"`
	Children int `xmlrpc:"children" structs:"children,omitempty"`
	Vlan int `xmlrpc:"vlan" structs:"vlan,omitempty"`
	Added time.Time `xmlrpc:"added" structs:"added,omitempty"`
	LastModified time.Time `xmlrpc:"last_modified" structs:"last_modified,omitempty"`
	TotalAddresses string `xmlrpc:"total_addresses" structs:"total_addresses,omitempty"`
	UsedAddreses string `xmlrpc:"used_addresses" structs:"used_addresses,omitempty"`
	FreeAddreses string `xmlrpc:"free_addresses" structs:"free_addresses,omitempty"`
	Status string `xmlrpc:"status" structs:"status,omitempty"`
	Expires string `xmlrpc:"expires" structs:"expires,omitempty"`
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

func (client *Client) PrefixSmartSearch(query string, options *SearchOptions) (error,[]Prefix) {
	args := make(map[string]interface{},0)
	args["query_string"] = query

	if options != nil {
		args["search_options"] = structs.Map(options)
	}

	response := struct{
		Error bool `xmlrpc:"error"`
		Result []Prefix `xmlrpc:"result"`
	}{}

	err := client.Run("smart_search_prefix", args, &response)
	return err, response.Result
}

func (client *Client) SearchPrefix(query sq.SearchQuery, options *SearchOptions) (error, []Prefix) {
	args := make(map[string]interface{},0)
	args["query"] = sq.Map(query)

	if options != nil {
		args["search_options"] = structs.Map(options)
	}

	response := struct{
		Error bool `xmlrpc:"error"`
		Result []Prefix `xmlrpc:"result"`
	}{}

	err := client.Run("search_prefix", args, &response)
	return err, response.Result
}

func (client *Client) DeletePrefix(prefix Prefix, recursive bool) error {
	args := make(map[string]interface{},0)
	args["prefix"] = structs.Map(prefix)
	args["recursive"] = recursive

	err := client.Run("remove_prefix", args, nil)
	return err
}

