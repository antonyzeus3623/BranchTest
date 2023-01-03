package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type PlcData struct {
	NodeID     string `json:"node_id"`
	DeviceData []Data `json:"device_data"`
}

type Data struct {
	DeviceNo     int32   `json:"device_no"`
	DeviceValues []int16 `json:"device_values"`
}

func main() {
	plcData := &PlcData{
		NodeID: "node03",
		DeviceData: []Data{{
			DeviceNo:     1,
			DeviceValues: []int16{1, 2, 3},
		},
			{
				DeviceNo:     2,
				DeviceValues: []int16{4, 5, 6},
			},
		},
	}

	// log.Println(plcData)

	b, err := json.Marshal(*plcData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

	// var data PlcData
	// err = json.Unmarshal(b, &data)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(data)
}
