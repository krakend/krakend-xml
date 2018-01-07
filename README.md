Krakend XML
====

Enable XML encoding in the [KrakenD](github.com/devopsfaith/krakend) framework

**Notice: this package requires KrakenD >= 0.4**

## Documentation

For more details, check the auto-generated documentation: https://godoc.org/github.com/devopsfaith/krakend-xml

## Using it

Just register the module before parsing your config and everything will work smoothly

	xml.Register()

## Build the example

Go 1.8 is a requirement

	$ make

## Run

Running it as a common executable, logs are send to the stdOut and some options are available at the CLI

	$ ./krakend_xml_example
	Usage of ./krakend_xml_example:
	  -c string
	    	Path to the configuration filename (default "/etc/krakend/configuration.json")
	  -d	Enable the debug
	  -p int
	    	Port of the service
