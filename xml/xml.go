package main

import (
	"encoding/xml"
	"fmt"
)

type S5F3Message struct {
	XMLName       xml.Name `xml:"Request"`
	Header        Header   `xml:"Header"`
	ActionType    string   `xml:"Body>ACTIONTYPE"`
	ActionCode    string   `xml:"Body>ACTIONCODE"`
	ActionComment string   `xml:"Body>ACTIONCOMMENT"`
	HoldOwner     string   `xml:"Body>HOLDOWNER"`
	ComponetList  XMLList  `xml:"Body>COMPONENTLIST>COMPONENT"` // 消息结构暂时待定
}

type Header struct {
	XMLName      xml.Name `xml:"Header"`
	MessageName  string   `xml:"MESSAGENAME"`
	TransationId string   `xml:"TRANSACTIONID"`
	Orgrrn       string   `xml:"ORGRRN"`
	OrgName      string   `xml:"ORGNAME"`
	UserName     string   `xml:"USERNAME"`
	Language     string   `xml:"LANGUAGE"`
}

type XMLList struct {
	XMLName   xml.Name `xml:"COMPONENT"`
	Alst      string   `xml:"ALST"`
	Alcd      string   `xml:"ALCD"`
	AlId      string   `xml:"ALID"`
	Altx      string   `xml:"ALTX"`
	UnitId    string   `xml:"UNITID"`
	ChartId   string   `xml:"CHARTID"`
	EqpId     string   `xml:"EQPID"`
	RuleType  string   `xml:"RULETYPE"`
	GraphType string   `xml:"GRAPHTYPE"`
	LotId     string   `xml:"LOTID"`
	GlsId     string   `xml:"GLSID"`
}

// S5F3MessageDecoder 解析SPC下发的Alarm消息
func S5F3MessageDecoder(b []byte) (*S5F3Message, error) {
	if b == nil {
		return nil, fmt.Errorf("invalid spc message")
	}

	rsp := new(S5F3Message)

	if err := xml.Unmarshal(b, &rsp); err != nil {
		return nil, err
	}

	return rsp, nil
}

// S5F3MessageMarshal 将结构体封装成xml消息
func S5F3MessageMarshal(msg *S5F3Message) ([]byte, error) {
	b, err := xml.MarshalIndent(msg, "", "        ")
	if err != nil {
		return nil, fmt.Errorf("struct marshal failed")
	}

	return b, nil
}
