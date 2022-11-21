# dagote - data for Go templates

## Context
Go, text template, html template, dynamic data, arbitrary JSON, YAML, TOML, CSV, XML, TEXT

## Purpose
Allows usage of arbitrary JSON, YAML, TOML, CSV, XML, TEXT in Go templates.

## General concept
'dagote' executes your text or html template set, and injects (optional) arbitrary content as 'dot' (.) data.

**Scenario 1:**
We inject no 'dot' data into the start template. We load all data (readJSON, readXML, ...) dynamically within our template set. We use the functions for data loading for this.

**Scenario 2:**
We inject configuration data as 'dot' data into the start template. Data transforming is applied before injection. In other words, the configuration data can be arbitrary JSON, YAML, CSV, CSVMap, Text, Lines, XML or TOML. Within the template set, we process the configuration data, e.g. to load arbitrary data (readJSON, readXML, ...).

**Scenario 3:**
We inject content data as 'dot' data into the start template. Data transforming is applied before injection. In other words, the content data can be arbitrary JSON, YAML, CSV, CSVMap, Text, Lines, XML or TOML. Within the template set, we process the content data. It is typically not nesseccary to load more data.

## Functions
'dagote' provides a rich function set for usage within your template set.

**Functions for data loading:**
* readJSON : reads JSON from file and unmarshals to 'map of any' (Go: map[string]any)
* readYAML : reads YAML from file and unmarshals to 'map of any' (Go: map[string]any)
* readCSV : reads all records of csv file into 'two-dimensional slice of strings' (Go: [][]string)
* readCSVMap : reads all records of csv file into 'slice of maps of strings' (Go: []map[string]string)
* readText : reads full text file into 'string' (Go: string)
* readLines : reads all lines of text file into 'slice of strings' (Go: []string)
* readXML : reads XML from file and unmarshals to 'map of any' (Go: map[string]any)
* readTOML : reads TOML from file and unmarshals to 'map of any' (Go: map[string]any)

**Functions for general purposes:**
* http://masterminds.github.io/sprig : general functions (sprig)
* https://pkg.go.dev/text/template : default functions (Go)

**Functions for basic file handling:**
* fileExists : checks whether file or directory exists (false, true)
* fileStat : returns FileInfo structure (Go: FileInfo {Name, Size, Mode, ModTime, IsDir, Sys})
* fileRead : reads arbitrary file into 'slice of bytes' (Go: []byte)

**Functions for html templates:**
* toTypeHTML : avoids autoescaping of HTML string (Go: template.HTML)
* toTypeCSS : avoids autoescaping of CSS string (Go: template.CSS)
* toTypeJS : avoids autoescaping of JS string (Go: template.JS)
* toTypeURL : avoids autoescaping of URL string (Go: template.URL)

**Note**: Use of the 'toType' functions presents a security risk. The encapsulated content should come from a trusted source, as it will be included verbatim in the html template output.

## 'dot' (.) data
The options '-dotfile, -dotstring, -dottype' are useful to (optionally) inject arbitrary data into the start template. The injected 'dot' data (.) can be considered as configuration or as content.

* configuration: describes what to do and/or which data to load
* content: represents the data to be processed within the template
* -dotfile: file content represents the data to be injected
* -dotstring: string content represents the data to be injected
* -dottype: data (file/string) will be transformed into 'dottype'

## Template files
For simple cases, a single template is often sufficient. Extensive or complex applications
usually require a large number of templates. The '-templates' option can be used to represent both.

**Scenario 1:**
One template: The '-templates' option defines one file.

**Scenario 2:**
Multiple templates (file list): The '-templates' option defines a list of files.

**Scenario 3:**
Multiple templates (wildcards): The '-templates' option defines a list of files and/or wildcard patterns (globs). The wildcard patterns are expanded to lists of files.

The first template in the template set is in all scenarios the start template. 

## Basic use within a Go template
``` text
{{ $json := readJSON "test.json" }}
{{ $yaml := readYAML "test.yaml" }}

{{ printf "name = %v\n" $json.name }}
{{ printf "age = %v\n" $json.age }}

{{ range $yaml.Comments }}
  {{ printf "%v, %v, %v\n" .Action .Date .Priority }}
{{ end }}
```

## Examples
* text-example : demonstrates usage for all supported data sources
* html-example : demonstrates usage of complex JSON data

## Binaries
Precompiled binaries for many operating systems can be found here : **Releases -> Assets**

## Program usage
``` text
dagote

Program:
  Name    : dagote
  Release : v1.0.0 - 2022/11/21
  Purpose : data for Go templates (dagote)
  Info    : Allows usage of arbitrary JSON, YAML, TOML, CSV, XML, TEXT in Go templates.

Usage:
  dagote -templates=list -output=file [-format=string] [-dotfile=file | -dotstring=string] [-dottype=string]

Examples (single template):
  dagote -templates=test.tmpl -output=test.txt -format=text
  dagote -templates=category.tmpl -output=category.html -format=html

Examples (set of templates):
  dagote -templates='test.tmpl,includes/*' -output=test.txt
  dagote -templates='test.tmpl,templates/*.tmpl,includes/*' -output=test.txt

Examples (dot data from file):
  dagote -templates=test.tmpl -output=test.txt -dotfile=test.json -dottype=json
  dagote -templates=test.tmpl -output=test.txt -dotfile=test.yaml -dottype=yaml

Examples (dot data from string):
  dagote -templates=test.tmpl -output=test.txt -dotstring='{"forum":"meta.discourse.org","topic":69776}' -dottype=json
  dagote -templates=test.tmpl -output=test.txt -dotstring='meta.discourse.org\n69776' -dottype=lines
  dagote -templates=test.tmpl -output=test.txt -dotstring='meta.discourse.org,69776' -dottype=csv
  dagote -templates=test.tmpl -output=test.txt -dotstring='meta.discourse.org,69776' -dottype=text

Notes concerning option '-templates':
  The templates list is a comma separates list of files and/or globs.
  The globs in the templates list will be expanded to a list of files.
  The first template in the list of files is the start template.

Notes concerning options '-dotfile, -dotstring, -dottype':
  These options allow to inject arbitrary data into the start template.
  The injected data (.) can be considered as configuration or as content.
    configuration: describes what to do and/or which data to load
    content: represents the data to be processed within the template
  -dotfile: file content represents the data to be injected
  -dotstring: string content represents the data to be injected
  -dottype: data (file/string) will be transformed into 'dottype'

Options:
  -dotfile string
    	dot data from file (injected into start template, accessible via .)
  -dotstring string
    	dot data from string (injected into start template, accessible via .)
  -dottype string
    	type of (file/string) dot data (json, yaml, toml, csv, csvmap, xml, text, lines) (default "text")
  -format string
    	format type (text, html) (default "text")
  -output string
    	name of output file
  -templates string
    	name of input template(s) (list of files and/or globs)
```

