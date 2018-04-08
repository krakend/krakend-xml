package main

import (
	"fmt"
	"io"

	"github.com/devopsfaith/krakend-xml"
)

var Registrable registrable

type registrable int

func (r *registrable) RegisterDecoder(setter func(name string, dec func(bool) func(io.Reader, *map[string]interface{}) error) error) error {
	fmt.Println("registrable", r, "from plugin 'krakend_xml' is registering its decoders at", setter)

	setter.Register(xml.Name, xml.NewDecoder)

	return nil
}

func init() {
	fmt.Println("plugin 'krakend_xml' loaded!")
}

func main() {

}
