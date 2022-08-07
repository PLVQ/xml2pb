package main

import (
	"encoding/xml"
	"io/ioutil"
	"strings"
	"xml2pb/config"
	"xml2pb/log"
	"xml2pb/parsexml"
)

type PbMsgList struct {
	ResConfig xml.Name `xml:"ResConfig"`
	Msg       []PbMsg  `xml:"struct"`
}

type PbMsg struct {
	Msg       xml.Name     `xml:"struct"`
	Name      string       `xml:"name,attr"`
	Desc      string       `xml:"desc,attr"`
	FieldList []PbMsgField `xml:"entry"`
}

type PbMsgField struct {
	Field xml.Name `xml:"entry"`
	Name  string   `xml:"name,attr"`
	Type  string   `xml:"type,attr"`
	CName string   `xml:"cname,attr"`
	Desc  string   `xml:"desc,attr"`
}

func main() {
	xmlFileList, err := ioutil.ReadDir(config.ConfigData.XmlPath)
	if err != nil {
		log.Log.WithField("err", err.Error()).Fatal("Read Xml Failed")
	}

	for _, file := range xmlFileList {
		if file.IsDir() {
			continue
		}

		if strings.Contains(file.Name(), ".xml") {
			log.Log.WithField("FileName", file.Name()).Info()
			xmlParse := new(parsexml.XmlParse)
			xmlParse.Parse(config.ConfigData.XmlPath + file.Name())
		}
	}
}
