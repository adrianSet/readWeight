package cache

import "readWeight/jsonEntity"

var(
	//串口读取体重信息
	Weight float32 = 50
    Upload = jsonEntity.DialysisProcess{50,50,50}

)


func ResetCache(){
	Weight=0
	Upload = jsonEntity.DialysisProcess{}
}
