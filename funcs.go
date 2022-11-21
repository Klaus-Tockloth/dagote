package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"os"
	"strings"

	xml "github.com/clbanning/mxj/v2"
	toml "github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
)

/*
readJSON reads JSON from file and unmarshals to map of any.
*/
func readJSON(filename string) (map[string]any, error) {
	if filename == "" {
		return nil, errors.New("readJSON needs a filename")
	}
	jsonMap := make(map[string]any)
	jsonRaw, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to read JSON file, file=[%v], error=[%w]", filename, err)
	}
	err = json.Unmarshal(jsonRaw, &jsonMap)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal JSON data, file=[%v], error=[%w]", filename, err)
	}
	return jsonMap, nil
}

/*
readYAML reads YAML from file and unmarshals to map of any.
*/
func readYAML(filename string) (map[string]any, error) {
	if filename == "" {
		return nil, errors.New("readYAML needs a filename")
	}
	yamlMap := make(map[string]any)
	yamlRaw, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to read YAML file, file=[%v], error=[%w]", filename, err)
	}
	err = yaml.Unmarshal(yamlRaw, &yamlMap)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal YAML data, file=[%v], error=[%w]", filename, err)
	}
	return yamlMap, nil
}

/*
readCSV reads all records of csv file into two-dimensional slice of strings.
*/
func readCSV(filename string) ([][]string, error) {
	if filename == "" {
		return nil, errors.New("readCSV needs a filename")
	}
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to read CSV file, file=[%v], error=[%w]", filename, err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("unable to read all CSV records, file=[%v], error=[%w]", filename, err)
	}
	return records, nil
}

/*
readCSVMap reads all records of csv file into slice of maps.
*/
func readCSVMap(filename string) ([]map[string]string, error) {
	if filename == "" {
		return nil, errors.New("readCSVMap needs a filename")
	}
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to read CSV file, file=[%v], error=[%w]", filename, err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("unable to read all CSV records, file=[%v], error=[%w]", filename, err)
	}
	returnMap := []map[string]string{}
	header := []string{} // holds first row (header)
	for lineNum, record := range rawCSVdata {
		// for first row, build the header slice
		if lineNum == 0 {
			for i := 0; i < len(record); i++ {
				header = append(header, strings.TrimSpace(record[i]))
			}
		} else {
			// for each cell, map[string]string k=header v=value
			line := map[string]string{}
			for i := 0; i < len(record); i++ {
				line[header[i]] = record[i]
			}
			returnMap = append(returnMap, line)
		}
	}
	return returnMap, nil
}

/*
readText reads text file into string.
*/
func readText(filename string) (string, error) {
	if filename == "" {
		return "", errors.New("readText needs a filename")
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("unable to read text file, file=[%v], error=[%w]", filename, err)
	}
	return string(data), nil
}

/*
readLines reads all lines of text file into slice of strings.
*/
func readLines(filename string) ([]string, error) {
	if filename == "" {
		return nil, errors.New("readTextLines needs a filename")
	}
	var lines []string
	file, err := os.ReadFile(filename)
	if err != nil {
		return lines, fmt.Errorf("unable to read text lines file, file=[%v], error=[%w]", filename, err)
	}
	buf := bytes.NewBuffer(file)
	for {
		line, err := buf.ReadString('\n')
		if len(line) == 0 {
			if err != nil {
				if err == io.EOF {
					break
				}
				return lines, fmt.Errorf("unable to read string, file=[%v], error=[%w]", filename, err)
			}
		}
		line = strings.TrimSuffix(line, "\n")
		lines = append(lines, line)
		if err != nil && err != io.EOF {
			return lines, fmt.Errorf("unable to read string, file=[%v], error=[%w]", filename, err)
		}
	}
	return lines, nil
}

/*
readXML reads XML from file and unmarshals to map of any.
*/
func readXML(filename string) (map[string]any, error) {
	if filename == "" {
		return nil, errors.New("readXML needs a filename")
	}
	xmlRaw, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to read XML file, file=[%v], error=[%w]", filename, err)
	}
	var xmlRoot xml.Map
	xmlRoot, err = xml.NewMapXml(xmlRaw)
	if err != nil {
		return nil, fmt.Errorf("unable to parse XML data, file=[%v], error=[%w]", filename, err)
	}
	xmlRootName, err := xmlRoot.Root()
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal XML data, file=[%v], error=[%w]", filename, err)
	}
	return xmlRoot[xmlRootName].(map[string]any), nil
}

/*
readTOML reads TOML from file and unmarshals to map of any.
*/
func readTOML(filename string) (map[string]any, error) {
	if filename == "" {
		return nil, errors.New("readTOML needs a filename")
	}
	tomlMap := make(map[string]any)
	tomlRaw, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to read TOML file, file=[%v], error=[%w]", filename, err)
	}
	err = toml.Unmarshal(tomlRaw, &tomlMap)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal TOML data, file=[%v], error=[%w]", filename, err)
	}
	return tomlMap, nil
}

/*
fileExists checks whether file or directory exists under given path.
*/
func fileExists(filename string) (bool, error) {
	if filename == "" {
		return false, errors.New("fileExists needs a filename")
	}
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, nil
}

/*
fileStat returns file info structure describing file.
*/
func fileStat(filename string) (os.FileInfo, error) {
	if filename == "" {
		return nil, errors.New("fileStat needs a filename")
	}
	info, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}
	return info, nil
}

/*
fileRead reads arbitrary file into slice of bytes.
*/
func fileRead(filename string) ([]byte, error) {
	if filename == "" {
		return nil, errors.New("fileRead needs a filename")
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to read file, file=[%v], error=[%w]", filename, err)
	}
	return data, nil
}

/*
toTypeHTML avoids autoescaping of HTML string (by html template engine.)
*/
func toTypeHTML(s string) template.HTML {
	return template.HTML(s)
}

/*
toTypeCSS avoids autoescaping of CSS string (by html template engine).
*/
func toTypeCSS(s string) template.CSS {
	return template.CSS(s)
}

/*
toTypeJS avoids autoescaping of JS string (by html template engine).
*/
func toTypeJS(s string) template.JS {
	return template.JS(s)
}

/*
toTypeURL avoids autoescaping of URL string (by html template engine).
*/
func toTypeURL(s string) template.URL {
	return template.URL(s)
}
