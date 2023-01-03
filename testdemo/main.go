package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/dualm/common"
	"github.com/spf13/viper"
)

// var (
// 	dbKey string
// )

// func main() {
// dbKey = time.Now().Format("01022006")
// log.Println(dbKey)
// result := strings.Compare("09052022", "09022022")
// log.Println(result)
// }

// func main() {
// 	fmt.Println("----------------引用传递----------------")
// 	for i := 0; i < 10; i++ {
// 		go func() {
// 			fmt.Println(i)
// 		}()
// 	}
// 	time.Sleep(10 * time.Millisecond)

// 	fmt.Println("----------------值传递----------------")
// 	for i := 0; i < 10; i++ {
// 		go func(x int) {
// 			fmt.Println(x)
// 		}(i)
// 	}
// 	time.Sleep(10 * time.Millisecond)

// 	return
// }

// type Student struct {
// 	Class string     `json:"class"`
// 	Info  []*StuInfo `json:"studentInfo"`
// }

// type StuInfo struct {
// 	Name  string `json:"name"`
// 	Age   int    `json:"age"`
// 	Score int    `json:"score"`
// }

func main() {
	// patternMain := `^([a-zA-Z]{1,2})([0-9]+)`
	// reg, err := regexp.Compile(patternMain)
	// if err != nil {
	// 	log.Println(err)
	// }

	// addr := "D110"
	// matches := reg.FindStringSubmatch(addr) // FindStringSubmatch方法是提取出匹配的字符串，然后通过[]string返回
	// log.Println(matches)

	// if len(matches) != 3 {
	// 	log.Printf("bad device address, %s", addr)
	// }

	// patternMain := `[bB]([01]{1,16})$`
	// reg, err := regexp.Compile(patternMain)
	// if err != nil {
	// 	log.Println(err)
	// }

	// addr := "D110b2"
	// matches := reg.FindStringSubmatch(addr) // FindStringSubmatch方法是提取出匹配的字符串，然后通过[]string返回
	// log.Println(matches)

	// if len(matches) != 2 {
	// 	log.Printf("bad device address, %s", addr)
	// }

	// s := "11111111"
	// mask := "10010000"
	// _s, err := strconv.ParseInt(s, 2, 16)
	// if err != nil {
	// 	log.Println(err)
	// }

	// _mask, err := strconv.ParseInt(mask, 2, 16)
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println(strconv.FormatInt(_s&_mask, 2))

	// var x int = 0
	// if x&0x80 == 0x80 {
	// 	fmt.Println("true")
	// } else {
	// 	fmt.Println("false")
	// }

	// var ip string = "www.baidu.com"
	// var port string = "8000"
	// addr, err := verifyAddr(ip, port)
	// if err != nil {
	// 	fmt.Println("err: ", err)
	// }
	// fmt.Println("addr: ", addr)

	// type Student struct {
	// 	grade int
	// 	age   int
	// }
	// stu := &Student{
	// 	grade: 100,
	// 	age:   18,
	// }
	// log.Println(stu)
	// stubyte := *(*[]byte)(unsafe.Pointer(stu))
	// log.Println(stubyte)

	// var a string = "3"
	// b := []byte(a)

	// log.Println(b)
	// var a []byte = []byte{179, 2}
	// var c uint8 = a[0]
	// d := c & 0x7F
	// log.Println(d)
	// log.Println(string(d))
	// var a int8 = 1
	// var b int16 = 1
	// var c int32 = 1
	// var d int64 = 1
	// log.Printf("a: %v", a)
	// log.Printf("b: %v", b)
	// log.Printf("c: %v", c)
	// log.Printf("d: %v", d)

	// stu := new(Student)
	// var stuInfo []*StuInfo
	// stuInfo = append(stuInfo, &StuInfo{Name: "小莉", Age: 19, Score: 100})
	// stuInfo = append(stuInfo, &StuInfo{Name: "平平", Age: 18, Score: 98})

	// stu.Class = "大学一年级1班"
	// stu.Info = append(stu.Info, stuInfo...)

	// data, _ := json.Marshal(stu)

	// log.Println(stu)
	// log.Printf("%#v\n", &stu.Info)

	// log.Println(string(data))
	// header := []byte{0b01000001, '1'}
	// ack := []byte("0")
	// fmt.Println(append(header, ack...))

	// x := false
	// if x {
	// 	fmt.Println("true")
	// }
	// for i := 0; i < 10; i++ {
	// 	log.Println("Hi~ ZhangWei~ Hello World!")
	// 	time.Sleep(time.Second * 1)
	// }

	// fn1 := "E:/众凌存储盘/9工程制造部/03整合测试科/28.CONDATA/20221207/Z2500/B4AC016BC5030S004-B.x1sx"
	// fn2 := "E:/众凌存储盘/9工程制造部/03整合测试科/28.CONDATA/20221207/Z2500/B4AC016BC5030S004_20221207.piece.scsv"
	// fn3 := "E:/众凌存储盘/9工程制造部/03整合测试科/28.CONDATA/20221207/Z2500/B4AC016BC5030S004.xlsx"

	// fns1 := strings.Split(strings.Split(strings.Split(filepath.Base(fn1), "-")[0], "_")[0], ".")
	// fns2 := strings.Split(strings.Split(strings.Split(filepath.Base(fn2), "-")[0], "_")[0], ".")
	// fns3 := strings.Split(strings.Split(strings.Split(filepath.Base(fn3), "-")[0], "_")[0], ".")

	// s1 := strings.TrimSuffix(strings.TrimSuffix(fns1[0], "-A"), "-B")
	// s2 := strings.TrimSuffix(strings.TrimSuffix(fns2[0], "-A"), "-B")

	// s3 := strings.TrimSuffix(strings.TrimSuffix(fns3[0], "-A"), "-B")

	// fmt.Println(s1)
	// fmt.Println(s2)
	// fmt.Println(s3)

	conf, err := InitConfig("config")
	if err != nil {
		log.Println(err)
	}
	str := common.GetStringSlice(conf, "SvFields", "Float")
	log.Println("Float的长度", len(str))
	log.Println("Float:", str)

}

func InitConfig(configId string) (*viper.Viper, error) {
	conf := viper.New()
	conf.SetConfigType("toml")
	conf.SetConfigName(strings.ToLower(configId))
	conf.AddConfigPath(".")
	err := conf.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("initialize config error, ConfigId: %s, Error:%w", configId, err)
	}

	return conf, nil
}

// func verifyAddr(ip, port string) (*net.TCPAddr, error) {
// 	addr, err := net.ResolveTCPAddr("tcp", ip+":"+port)
// 	if err != nil {
// 		return nil, fmt.Errorf("invalid ip or port, ip: %s, port: %s", ip, port)
// 	}

// 	return addr, nil
// }
