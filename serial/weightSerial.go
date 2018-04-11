package serial

import (
	"github.com/tarm/serial"
	"readWeight/config"
	"time"
	. "readWeight/log1"
	"strconv"
	"readWeight/cache"
)

type temp struct {

	index                  int
	//多次体重数组
	list                   []float64
	//体重连续稳定次数
	continueSameWeightFlag int
}





var (
	//是否尝试读取串口信息
	Flag = 0
	//体重稳定对象
	t    = temp{0, make([]float64, 100), 0}
	//串口指针
	s *serial.Port
)

func init() {
	Logger.Println("串口初始化!!!")
	openSerial()

}

func openSerial() {
	c := &serial.Config{Name: config.M[config.WeightMachine], Baud: 9600}
	var port, err = serial.OpenPort(c)
	s = port
	if err != nil {
		Logger.Println("体重器串口打开失败:", err)
		time.Sleep(time.Second * 5)
		openSerial()
	}
}

func reOpen(){
	s.Close()
	time.Sleep(time.Second)
	temp1 = make([]byte, 0, 14)
	openSerial()
}

/**
监听串口信息
 */
 
 var temp1 []byte
func Listening() {
	//go func(){
	//	for{
	//		s.Write([]byte("R"))
	//		time.Sleep(time.Second*1)
	//	}
	//}()
	temp1 = make([]byte, 0, 14)
	for {
		//是否尝试读取串口信息
		//if Flag == 0 {
		//	time.Sleep(time.Second)
		//	continue
		//}
		time.Sleep(time.Millisecond * 10)
		buf := make([]byte, 14)
		n, err := s.Read(buf)
		if err != nil {
			Logger.Println("体重读取失败，尝试重新打开串口", err)
			openSerial()
		} else {
			tlen := len(temp1)
			//Logger.Printf("%q|len:%d ", buf[:n], n)

			if n+tlen < 14 {
				temp1 = append(temp1, buf[:(n)]...)
				continue
			}

			if n+tlen == 14 {
				temp1 = append(temp1, buf[:(n)]...)
				result := temp1
				temp1 = make([]byte, 0, 14)
				//Logger.Printf("result : %q ", result)
				saveWeightAndVaildate(parseWeight(result))
			}
			if n == 14 && tlen == 1 {
				temp1 = append(temp1, buf[:(n - tlen)]...)
				result := temp1
				temp1 = make([]byte, 0, 14)
				temp1 = append(temp1, buf[n-tlen:]...)
				//Logger.Printf("result : %q ", result)
				saveWeightAndVaildate(parseWeight(result))
			}
		}
	}
}

func parseWeight(byte []byte) float64{
	//Logger.Printf("result : %q ", byte)
	var string1 string
	 s3:=string(byte[5])
	weight:=string(byte[6:10])
	if "0" != s3{
		string1+=s3
	}
	string1+=weight
	v2, err := strconv.ParseFloat(string1, 64)
	if err!=nil {
		Logger.Println("体重信息读取异常",err,"尝试重启串口")
		reOpen()

	}else{
		//Logger.Println("体重为:",v2)
	}
	return v2

}


func saveWeightAndVaildate(weight float64) {
		if weight ==0 || weight<10{
			return
		}
		Logger.Printf("获取体重:%dkg\n", weight)
		isStability := continueSameWeight(weight)
		if isStability {
			Logger.Printf("连续%d次 成功获取稳定体重%dkg", t.continueSameWeightFlag, weight)
			t = temp{0, make([]float64, 50), 0}
			cache.Upload.PreBodyWeight = float64(weight)
			Flag = 0
		}
}

func continueSameWeight(w float64) bool {
	if t.index < len(t.list) {
		t.list[t.index] = w

		if t.index-1 < 0 {
			t.continueSameWeightFlag++
		} else if t.list[t.index-1] == t.list[t.index] {
			t.continueSameWeightFlag++
		} else {
			t.continueSameWeightFlag = 0
		}
		t.index++
		Logger.Printf("此次index[%d] 连续[%d]", t.index, t.continueSameWeightFlag)

	} else {
		t = temp{0, make([]float64, 50), 0}
	}

	if t.continueSameWeightFlag >= 3 {
		return true
	}
	return false
}
