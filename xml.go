package xml

import (
	"io"

	"golang.org/x/net/html/charset"
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
	mxj.XmlCharsetReader = charset.NewReaderLabel
	mv, err := mxj.NewMapXmlReader(xmlReader{r: r})
	if err != nil {
		return err
	}
	*v = mv
	return nil

}

// CollectionDecoder implements the Decoder interface over a collection
func CollectionDecoder(r io.Reader, v *map[string]interface{}) error {
	mxj.XmlCharsetReader = charset.NewReaderLabel
	mv, err := mxj.NewMapXmlReader(xmlReader{r: r})
	if err != nil {
		return err
	}
	*(v) = map[string]interface{}{"collection": mv}
	return nil
}

type xmlReader struct {
	r io.Reader
}

func (x xmlReader) Read(p []byte) (n int, err error) {
	n, err = x.r.Read(p)

	if err != io.EOF {
		return n, err
	}

	if len(p) == n {
		return n, nil
	}

	p[n] = ([]byte("\n"))[0]
	return n + 1, err
}
