package xml

import (
	"io"

	"github.com/clbanning/mxj"
	"github.com/devopsfaith/krakend/encoding"
)

// Register registers the xml decoder
func Register() error {
	return encoding.Register(Name, NewDecoder)
}

// Name is the key for the xml encoding
const Name = "xml"

// NewDecoder return the right XML decoder
func NewDecoder(isCollection bool) func(io.Reader, *map[string]interface{}) error {
	if isCollection {
		return CollectionDecoder
	}
	return Decoder
}

// Decoder implements the Decoder interface
func Decoder(r io.Reader, v *map[string]interface{}) error {
	mv, err := mxj.NewMapXmlReader(r)
	if err != nil {
		return err
	}
	*v = mv
	return nil

}

// CollectionDecoder implements the Decoder interface over a collection
func CollectionDecoder(r io.Reader, v *map[string]interface{}) error {
	mv, err := mxj.NewMapXmlReader(r)
	if err != nil {
		return err
	}
	*(v) = map[string]interface{}{"collection": mv}
	return nil
}
