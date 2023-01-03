package main

import "fmt"

func main() {
	// var _devtyp int32
	// var err error
	// var j int64
	// str := "ZR1000"

	// x, _x := deviceTypeNo(str)
	// fmt.Println("type: ", x, "no: ", _x)
	// switch x {
	// case "L":
	// 	j, err = strconv.ParseInt("0003", 16, 32)
	// 	_devtyp = int32(j)
	// case "M":
	// 	j, err = strconv.ParseInt("0004", 16, 32)
	// 	_devtyp = int32(j)
	// case "D":
	// 	j, err = strconv.ParseInt("000D", 16, 32)
	// 	_devtyp = int32(j)
	// case "B":
	// 	j, err = strconv.ParseInt("0017", 16, 32)
	// 	_devtyp = int32(j)
	// case "W":
	// 	j, err = strconv.ParseInt("0018", 16, 32)
	// 	_devtyp = int32(j)
	// case "R":
	// 	j, err = strconv.ParseInt("0016", 16, 32)
	// 	_devtyp = int32(j)
	// case "ZR":
	// 	j, err = strconv.ParseInt("00DCH", 16, 32)
	// 	_devtyp = int32(j)
	// }
	// if err != nil {
	// 	log.Println("deviceType转换为int32失败, err: ", err)
	// }
	// fmt.Printf("%#v", _devtyp)

	// t := time.Now()
	// fmt.Println(t)
	// _t, err := t.Parse("20060102")
	// _today, err := strconv.Atoi(time.Now().Format("20060102"))
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(_t)
	// fmt.Println(_today)
	// fmt.Println(_today - 1)

	// s := make(map[string]string)
	// s["力力"] = "lili"
	// s["力力2"] = "lili2"
	// s["力力3"] = "lili3"
	// for k := range s {
	// 	fmt.Println(k)
	// }

	// s1 := make([]string, 0, len(s))
	// for k := range s {
	// 	s1 = append(s1, k)
	// }

	s := "OK1"
	len := len([]byte(s + " "))

	fmt.Println(len)
}

// func deviceTypeNo(str string) (string, int32) {
// 	var deviceType string
// 	for i := range str {
// 		if !unicode.IsNumber(rune(str[i])) {
// 			deviceType = deviceType + string(str[i])
// 		}
// 	}
// 	deviceNo := strings.Trim(str, deviceType)
// 	_deviceNo, err := strconv.ParseInt(deviceNo, 10, 32)
// 	if err != nil {
// 		log.Println("string转int32转换失败, err:", err)
// 	}
// 	return deviceType, int32(_deviceNo)
// }
