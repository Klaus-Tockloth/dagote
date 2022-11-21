package main

import (
	"fmt"
	"os"
	"strings"
)

/*
determineDotData idetermines 'dot' (.) data from string or file.
*/
func determineDotData() (any, error) {
	var dotdata any
	var err error

	if *dotstring != "" {
		// create temporary file (to unify dot data processing)
		f, err := os.CreateTemp("", "dotstring.*.txt")
		if err != nil {
			return nil, fmt.Errorf("-dotstring: unable to create temporary file, error=[%v]", err)
		}
		tempFilename := f.Name()
		defer os.Remove(tempFilename)
		ds := strings.ReplaceAll(*dotstring, "\\n", "\n")
		_, err = f.WriteString(ds)
		if err != nil {
			f.Close()
			return nil, fmt.Errorf("-dotstring: unable to write to temporary file, file=[%v], error=[%v]", tempFilename, err)
		}
		err = f.Close()
		if err != nil {
			return nil, fmt.Errorf("-dotstring: unable to close temporary file, file=[%v], error=[%v]", tempFilename, err)
		}
		*dotfile = tempFilename
	}

	if *dotfile != "" {
		switch strings.ToLower(*dottype) {
		case "json":
			dotdata, err = readJSON(*dotfile)
			if err != nil {
				return nil, fmt.Errorf("unable to transform dot file to JSON, file=[%v], error=[%v]", *dotfile, err)
			}
		case "yaml":
			dotdata, err = readYAML(*dotfile)
			if err != nil {
				return nil, fmt.Errorf("unable to transform dot file to YAML, file=[%v], error=[%v]", *dotfile, err)
			}
		case "csv":
			dotdata, err = readCSV(*dotfile)
			if err != nil {
				return nil, fmt.Errorf("unable to transform dot file to CSV, file=[%v], error=[%v]", *dotfile, err)
			}
		case "csvmap":
			dotdata, err = readCSVMap(*dotfile)
			if err != nil {
				return nil, fmt.Errorf("unable to transform dot file to CSVMap, file=[%v], error=[%v]", *dotfile, err)
			}
		case "text":
			dotdata, err = readText(*dotfile)
			if err != nil {
				return nil, fmt.Errorf("unable to transform dot file to TEXT, file=[%v], error=[%v]", *dotfile, err)
			}
		case "lines":
			dotdata, err = readLines(*dotfile)
			if err != nil {
				return nil, fmt.Errorf("unable to transform dot file to LINES, file=[%v], error=[%v]", *dotfile, err)
			}
		case "xml":
			dotdata, err = readXML(*dotfile)
			if err != nil {
				return nil, fmt.Errorf("unable to transform dot file to XML, file=[%v], error=[%v]", *dotfile, err)
			}
		case "toml":
			dotdata, err = readTOML(*dotfile)
			if err != nil {
				return nil, fmt.Errorf("unable to transform dot file to TOML, file=[%v], error=[%v]", *dotfile, err)
			}
		default:
			return nil, fmt.Errorf("unsupported dot type, type=[%v]", *dottype)
		}
	}

	return dotdata, nil
}
