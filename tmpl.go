package main

import (
	"bufio"
	"fmt"
	htmltemplate "html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
	texttemplate "text/template"

	"github.com/Masterminds/sprig/v3"
)

/*
determineTemplateFiles determines template files for parsing.
*/
func determineTemplateFiles() ([]string, error) {
	globs := strings.Split(*templates, ",")
	var templateFiles []string
	for _, glob := range globs {
		tmpfiles, err := filepath.Glob(glob)
		if err != nil {
			return nil, fmt.Errorf("error [%v] at filepath.Glob()", err)
		}
		// ignore directories
		for _, tmpfile := range tmpfiles {
			info, err := os.Stat(tmpfile)
			if err != nil {
				return nil, fmt.Errorf("error [%v] at os.Stat()", err)
			}
			if info.IsDir() {
				continue
			}
			templateFiles = append(templateFiles, tmpfile)
		}
	}
	if len(templateFiles) == 0 {
		return nil, fmt.Errorf("no template file found for parsing")
	}
	fmt.Printf("Files for template parsing:\n")
	for i := range templateFiles {
		fmt.Printf("- %s\n", templateFiles[i])
	}

	return templateFiles, nil
}

/*
processTemplates processes (parse, execute) template file set.
*/
func processTemplates(templateFiles []string, dotdata any) error {
	var err error

	*format = strings.ToLower(*format)
	switch *format {
	case "text":
		// create text template with functions
		templ := texttemplate.New(templateFiles[0]).Funcs(sprig.FuncMap()).Funcs(
			texttemplate.FuncMap{
				"readJSON":   readJSON,
				"readYAML":   readYAML,
				"readCSV":    readCSV,
				"readCSVMap": readCSVMap,
				"readText":   readText,
				"readLines":  readLines,
				"readXML":    readXML,
				"readTOML":   readTOML,
				"fileExists": fileExists,
				"fileStat":   fileStat,
				"fileRead":   fileRead,
				"toTypeHTML": toTypeHTML,
				"toTypeCSS":  toTypeCSS,
				"toTypeJS":   toTypeJS,
				"toTypeURL":  toTypeURL,
			})

		// parse template
		fmt.Printf("\nParsing text template(s) ...\n")
		templ, err = templ.ParseFiles(templateFiles...)
		if err != nil {
			return fmt.Errorf("unable to parse text template(s), error=[%v]", err)
		}

		startTemplate := templ.Name()
		fmt.Printf("\nTemplates defined after parsing:\n")
		for _, template := range templ.Templates() {
			fmt.Printf("-  %s\n", template.Name())
		}

		// execute template
		fmt.Printf("\nExecuting text template [%s] -> [%s] ...\n", startTemplate, *outputFile)
		file, err := os.OpenFile(*outputFile, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
		if err != nil {
			return fmt.Errorf("unable to open output file, file=[%v], error=[%v]", *outputFile, err)
		}
		writer := bufio.NewWriter(file)
		err = templ.Execute(writer, dotdata)
		if err != nil {
			return fmt.Errorf("unable to execute text template, name=[%v], error=[%v]", startTemplate, err)
		}
		err = writer.Flush()
		if err != nil {
			log.Fatalf("unable to flush output file, file=[%v], error=[%v]", *outputFile, err)
		}
		err = file.Close()
		if err != nil {
			return fmt.Errorf("unable to close output file, file=[%v], error=[%v]", *outputFile, err)
		}
		fmt.Printf("Done.\n")

	case "html":
		// create html template with functions
		templ := htmltemplate.New(templateFiles[0]).Funcs(sprig.FuncMap()).Funcs(
			htmltemplate.FuncMap{
				"readJSON":   readJSON,
				"readYAML":   readYAML,
				"readCSV":    readCSV,
				"readCSVMap": readCSVMap,
				"readText":   readText,
				"readLines":  readLines,
				"readXML":    readXML,
				"readTOML":   readTOML,
				"fileExists": fileExists,
				"fileStat":   fileStat,
				"fileRead":   fileRead,
				"toTypeHTML": toTypeHTML,
				"toTypeCSS":  toTypeCSS,
				"toTypeJS":   toTypeJS,
				"toTypeURL":  toTypeURL,
			})

		// parse template
		fmt.Printf("\nParsing html template(s) ...\n")
		templ, err = templ.ParseFiles(templateFiles...)
		if err != nil {
			return fmt.Errorf("unable to parse html template(s), error=[%v]", err)
		}

		startTemplate := templ.Name()
		fmt.Printf("\nTemplates defined after parsing:\n")
		for _, template := range templ.Templates() {
			fmt.Printf("-  %s\n", template.Name())
		}

		// execute template
		fmt.Printf("\nExecuting html template [%s] -> [%s] ...\n", startTemplate, *outputFile)
		file, err := os.OpenFile(*outputFile, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
		if err != nil {
			return fmt.Errorf("unable to open output file, file=[%v], error=[%v]", *outputFile, err)
		}
		writer := bufio.NewWriter(file)
		err = templ.Execute(writer, dotdata)
		if err != nil {
			return fmt.Errorf("unable to execute html template, name=[%v], error=[%v]", startTemplate, err)
		}
		err = writer.Flush()
		if err != nil {
			return fmt.Errorf("unable to flush output file, file=[%v], error=[%v]", *outputFile, err)
		}
		err = file.Close()
		if err != nil {
			return fmt.Errorf("unable to close output file, file=[%v], error=[%v]", *outputFile, err)
		}
		fmt.Printf("Done.\n")

	default:
		return fmt.Errorf("option '-format=%s' not supported", *format)
	}

	return nil
}
