package main

import (
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./test.xml") // For read access.
	if err != nil {
		log.Printf("open file failed, error: %v", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Printf("read file failed, error: %v", err)
	}

	res, err := S5F3MessageDecoder(data)
	if err != nil {
		log.Println("Unmarshal xml message failed, err: ", err)
	}

	log.Printf("%#v\n", res)

	msg := &S5F3Message{

		Header: Header{
			MessageName:  "MES.COMPONENTHOLD",
			TransationId: "GLLy65t9NyjQJurT1PBo8Y",
			Orgrrn:       "+orgrrn+",
			OrgName:      "+factoryName+",
			UserName:     "PDP",
			Language:     "ZH",
		},
		ActionType:    "HOLD",
		ActionCode:    "SPCHD",
		ActionComment: "alarmTypeList",
		HoldOwner:     "holdOwner",
		ComponetList: XMLList{
			Alst:      "Alst",
			Alcd:      "Alcd",
			AlId:      "AlId",
			Altx:      "Altx",
			UnitId:    "UnitId",
			ChartId:   "ChartId",
			EqpId:     "EqpId",
			RuleType:  "RuleType",
			GraphType: "GraphType",
			LotId:     "LotId",
			GlsId:     "GlsId",
		},
	}

	Msg, err := S5F3MessageMarshal(msg)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(Msg))
}
