/*
Purpose:
- data for Go templates (dagote)

Description:
- Allows usage of arbitrary JSON, YAML, TOML, CSV, XML, TEXT in Go templates.

Releases:
- v1.0.0 - 2022/11/21: initial release

Author:
- Klaus Tockloth

Copyright:
- Copyright (c) 2022 Klaus Tockloth

Contact:
- klaus.tockloth@googlemail.com

Remarks:
- Lint: golangci-lint run --no-config --enable gocritic
- Vulnerability detection: govulncheck ./...

ToDo:
- NN

Links:
- https://github.com/Masterminds/sprig
- https://pkg.go.dev/github.com/Masterminds/sprig/v3
- http://masterminds.github.io/sprig/
- https://gohugo.io/templates/introduction/
- https://stackoverflow.com/questions/57102134/does-go-std-lib-have-a-func-to-read-csv-file-into-mapstringstring
- https://pkg.go.dev/github.com/clbanning/mxj/v2
- https://github.com/clbanning/mxj
- https://stackoverflow.com/questions/41176355/go-template-name
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// general program info
var (
	progName    = os.Args[0]
	progVersion = "v1.0.0"
	progDate    = "2022/11/21"
	progPurpose = "data for Go templates (dagote)"
	progInfo    = "Allows usage of arbitrary JSON, YAML, TOML, CSV, XML, TEXT in Go templates."
)

// command line parameters
var (
	format     *string
	templates  *string
	outputFile *string
	dotfile    *string
	dotstring  *string
	dottype    *string
)

/*
main starts this program.
*/
func main() {
	var err error

	fmt.Printf("\nProgram:\n")
	fmt.Printf("  Name    : %s\n", progName)
	fmt.Printf("  Release : %s - %s\n", progVersion, progDate)
	fmt.Printf("  Purpose : %s\n", progPurpose)
	fmt.Printf("  Info    : %s\n\n", progInfo)

	log.SetFlags(0)
	log.SetPrefix("error: ")

	format = flag.String("format", "text", "format type (text, html)")
	templates = flag.String("templates", "", "name of input template(s) (list of files and/or globs)")
	outputFile = flag.String("output", "", "name of output file")
	dotfile = flag.String("dotfile", "", "dot data from file (injected into start template, accessible via .)")
	dotstring = flag.String("dotstring", "", "dot data from string (injected into start template, accessible via .)")
	dottype = flag.String("dottype", "text", "type of (file/string) dot data (json, yaml, toml, csv, csvmap, xml, text, lines)")

	flag.Usage = printUsage
	flag.Parse()
	if flag.NFlag() == 0 {
		printUsage()
	}
	if *templates == "" {
		log.Fatalf("option '-templates=list' required")
	}
	if *outputFile == "" {
		log.Fatalf("option '-output=file' required")
	}
	if *dotfile != "" && *dotstring != "" {
		log.Fatalf("use either option '-dotfile=file' or option '-dotstring=string'")
	}

	dotdata, err := determineDotData()
	if err != nil {
		log.Fatalf("unable to determine dot data, error=[%v]", err)
	}

	templateFiles, err := determineTemplateFiles()
	if err != nil {
		log.Fatalf("unable to determine template file(s), error=[%v]", err)
	}

	err = processTemplates(templateFiles, dotdata)
	if err != nil {
		log.Fatalf("unable to process template(s), error=[%v]", err)
	}

	fmt.Printf("\n")
}

/*
printUsage prints the usage of this program.
*/
func printUsage() {
	fmt.Printf("Usage:\n")
	fmt.Printf("  %s -templates=list -output=file [-format=string] [-dotfile=file | -dotstring=string] [-dottype=string]\n", os.Args[0])

	fmt.Printf("\nExamples (single template):\n")
	fmt.Printf("  %s -templates=test.tmpl -output=test.txt -format=text\n", os.Args[0])
	fmt.Printf("  %s -templates=category.tmpl -output=category.html -format=html\n", os.Args[0])

	fmt.Printf("\nExamples (set of templates):\n")
	fmt.Printf("  %s -templates='test.tmpl,includes/*' -output=test.txt\n", os.Args[0])
	fmt.Printf("  %s -templates='test.tmpl,templates/*.tmpl,includes/*' -output=test.txt\n", os.Args[0])

	fmt.Printf("\nExamples (dot data from file):\n")
	fmt.Printf("  %s -templates=test.tmpl -output=test.txt -dotfile=test.json -dottype=json\n", os.Args[0])
	fmt.Printf("  %s -templates=test.tmpl -output=test.txt -dotfile=test.yaml -dottype=yaml\n", os.Args[0])

	fmt.Printf("\nExamples (dot data from string):\n")
	fmt.Printf("  %s -templates=test.tmpl -output=test.txt -dotstring='{\"forum\":\"meta.discourse.org\",\"topic\":69776}' -dottype=json\n", os.Args[0])
	fmt.Printf("  %s -templates=test.tmpl -output=test.txt -dotstring='meta.discourse.org\\n69776' -dottype=lines\n", os.Args[0])
	fmt.Printf("  %s -templates=test.tmpl -output=test.txt -dotstring='meta.discourse.org,69776' -dottype=csv\n", os.Args[0])
	fmt.Printf("  %s -templates=test.tmpl -output=test.txt -dotstring='meta.discourse.org,69776' -dottype=text\n", os.Args[0])

	fmt.Printf("\nNotes concerning option '-templates':\n")
	fmt.Printf("  The templates list is a comma separates list of files and/or globs.\n")
	fmt.Printf("  The globs in the templates list will be expanded to a list of files.\n")
	fmt.Printf("  The first template in the list of files is the start template.\n")

	fmt.Printf("\nNotes concerning options '-dotfile, -dotstring, -dottype':\n")
	fmt.Printf("  These options allow to inject arbitrary data into the start template.\n")
	fmt.Printf("  The injected data (.) can be considered as configuration or as content.\n")
	fmt.Printf("    configuration: describes what to do and/or which data to load\n")
	fmt.Printf("    content: represents the data to be processed within the template\n")
	fmt.Printf("  -dotfile: file content represents the data to be injected\n")
	fmt.Printf("  -dotstring: string content represents the data to be injected\n")
	fmt.Printf("  -dottype: data (file/string) will be transformed into 'dottype'\n")

	fmt.Printf("\nOptions:\n")
	flag.PrintDefaults()

	fmt.Printf("\n")
	os.Exit(1)
}
