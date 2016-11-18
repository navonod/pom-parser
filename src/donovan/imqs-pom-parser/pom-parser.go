package main

import (
	"io/ioutil"
	"flag"
	"fmt"
	"encoding/xml"
	"bytes"
)

type project struct {

}

type Project struct {
	dependencies []Dependency `xml:"dependencies"`
}

type Dependency struct {
	groupId string
	artifactId string
}

// parse a pom's xml
func (p *pomParser) parse(xmlFilepath string) (string, error) {

	fmt.Println("Loading pom file:", xmlFilepath)

	byteArray, err := ioutil.ReadFile(xmlFilepath)
	if (err != nil) {
		fmt.Print(err)
		return "", err
	}

	var project Project
	xml.Unmarshal(byteArray, &project)
	fmt.Printf("%+v\n", project)

	var buffer bytes.Buffer

	fmt.Println("Found", len(project.dependencies), "dependencies")
	for _, dependency := range project.dependencies {
		buffer.WriteString(dependency.artifactId)
	}

	return buffer.String(), nil
}

func main() {
	var rootPath string

	flag.StringVar(&rootPath, "path", "pom.xml", "Please enter the root path for the parser")
	flag.Parse()

	parser := &pomParser{}
	fmt.Println("Parsing pom file:", rootPath)
	fmt.Print(parser.parse(rootPath))
}