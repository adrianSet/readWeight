package cache

import "readWeight/jsonEntity"

var(
	//串口读取体重信息
    Upload = jsonEntity.DialysisProcess{}

)


func ResetCache(){
	Upload = jsonEntity.DialysisProcess{}
}
