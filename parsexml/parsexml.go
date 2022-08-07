package parsexml

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"xml2pb/generatepbfile"
)

const TAP_BLANK_NUM = "    "
const (
	FIELD_SINGULAR_TYPE = "singular"
	FIELD_REPEATED_TYPE = "repeated"
)

type PbMsgFile struct {
	ResConfig xml.Name `xml:"ResConfig"`
	Name      string   `xml:"name,attr"`
	MsgList   []PbMsg  `xml:"struct"`
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
	Rule  string   `xml:"rule,attr"`
	Type  string   `xml:"type,attr"`
	CName string   `xml:"cname,attr"`
	Desc  string   `xml:"desc,attr"`
}

type PBMessage struct {
	FieldIdx   int
	MsgContext string
}

type XmlParse struct {
	MessageSet map[string]*PBMessage
}

func (pXmlParse *XmlParse) Parse(xmlFile string) {
	pXmlParse.MessageSet = make(map[string]*PBMessage)
	data, err := ioutil.ReadFile(xmlFile)
	if err != nil {
		fmt.Println("读文件出错！", err)
		return
	}
	// fmt.Println(string(bytes))
	var pbMsgFileData PbMsgFile
	err = xml.Unmarshal(data, &pbMsgFileData)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Println(pbMsgFileData)

	for _, msg := range pbMsgFileData.MsgList {
		for _, field := range msg.FieldList {
			pXmlParse.createPbMsg(msg.Name, field.Name, field.Rule, field.Type, field.Desc)
		}
	}
	pXmlParse.createMsgDefine(pbMsgFileData.Name)
}

func (pXmlParse *XmlParse) createPbMsg(msgName string, filedName string, fieldRule string,
	fieldType string, fieldDesc string) {
	_, ok := pXmlParse.MessageSet[msgName]
	if !ok {
		pPbMsg := new(PBMessage)
		pXmlParse.MessageSet[msgName] = pPbMsg
	}

	pXmlParse.createDesc(msgName, fieldDesc)
	pXmlParse.createField(msgName, filedName, fieldRule, fieldType)
}

func (pXmlParse *XmlParse) createDesc(msgName string, desc string) {
	msg := pXmlParse.MessageSet[msgName]
	if strings.Count(desc, "\n") > 1 {
		if desc[len(desc)-1] != '\n' {
			desc = desc + "\n"
		}
		desc = strings.Replace(desc, "\n", "\n"+TAP_BLANK_NUM, -1)
		desc = strings.Replace(desc, "\n\n", "\n", -1)
		msg.MsgContext += TAP_BLANK_NUM + "/** " + desc + TAP_BLANK_NUM + "*/\n"
	} else {
		msg.MsgContext += TAP_BLANK_NUM + "/** " + desc + " */\n"
	}
}

func (pXmlParse *XmlParse) createField(msgName string, fieldName string, fieldRule string, fieldType string) {
	msg := pXmlParse.MessageSet[msgName]
	var fileRuleStr string
	if fieldRule == FIELD_SINGULAR_TYPE {
		fileRuleStr = ""
	} else {
		fileRuleStr = fieldRule + " "
	}

	msg.FieldIdx++
	msg.MsgContext += TAP_BLANK_NUM + fileRuleStr + fieldType + " " + fieldName + " = " + strconv.Itoa(msg.FieldIdx) + ";\n"
}

func (pXmlParse *XmlParse) createMsgDefine(pbFileName string) {
	var pbFileContext string
	for msgName, msg := range pXmlParse.MessageSet {
		pbFileContext += "message " + msgName + "{\n" + msg.MsgContext + "\n}\n\n"
	}

	generatepbfile.GeneratePbFile(pbFileName, pbFileContext)
}
