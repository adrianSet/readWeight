package main

import (
	"net/http"
	"log"
	"fmt"
	"html/template"
	"readWeight/config"
	"io/ioutil"
	. "readWeight/log1"
	"encoding/json"
	"readWeight/cache"
	"net/url"
	"strconv"
	"readWeight/serial"
)

var (
	patientInfoUrl  = config.M["host"] + config.M["patient_info_url"]
	uploadWeightUrl = config.M["host"] + config.M["upload_weight"]
)

func init() {

}

//func sayhelloName(w http.ResponseWriter, r *http.Request) {
//	r.ParseForm() //解析参数，默认是不会解析的
//	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
//	fmt.Println("path", r.URL.Path)
//	fmt.Println("scheme", r.URL.Scheme)
//	fmt.Println(r.Form["url_long"])
//	for k, v := range r.Form {
//		fmt.Println("key:", k)
//		fmt.Println("val:", strings.Join(v, ""))
//	}
//	fmt.Fprintf(w, "Hello Wrold!") //这个写入到w的是输出到客户端的
//}

//***********************request handle *********************************
func index(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("html/index.html")
	locals := make(map[string]interface{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmp.Execute(w, locals)

}

func initScaleWeight(){
	cache.ResetCache()
	serial.Flag=1

}

func getPatientInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	//称重初始化
	initScaleWeight()
	dialysisCode := r.FormValue("dialysisCode")
	resp, err := http.Get(patientInfoUrl + "?dialysisCode=" + dialysisCode)
	if err != nil {
		// handle error
		Logger.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	json := string(body)
	parseJsonOfPatientInfo(json)
	Logger.Println(json)
	fmt.Fprintf(w, json)
}

/**
呼吸事件
 */
func breathe(w http.ResponseWriter, r *http.Request) {
	if cache.Upload.PreBodyWeight!=0{
		fmt.Fprint(w,cache.Upload.PreBodyWeight)
	}else{
		fmt.Fprint(w,0)
	}
}

/**
病人完成称重过程 初始化环境变量
 */
func finished(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json;charset=utf-8")

	patientId:=strconv.Itoa(cache.Upload.PatientId)
	dialysisProcessId:=strconv.Itoa(cache.Upload.DialysisProcessId)
	preBodyWeight:=strconv.FormatFloat(cache.Upload.PreBodyWeight, 'E', -1, 32)
	data:=url.Values{"patientId":{patientId},"preBodyWeight":{preBodyWeight},"dialysisProcessId":{dialysisProcessId}}

	resp,err:=http.PostForm(uploadWeightUrl,data)
	//重置
	cache.ResetCache()
	if err!=nil {
		Logger.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	json := string(body)
	fmt.Fprint(w,json)
}
//**********************private method **********************************

/**
病人信息解析并存储
 */
func parseJsonOfPatientInfo(jsonstr string) {
	var c map[string]interface{}
	err := json.Unmarshal([]byte(jsonstr), &c)
	if err != nil {
		Logger.Println(err)
	}
	code:=c["code"].(float64)
	if code == 666{
		result:=c["result"].(map[string]interface{})
		cache.Upload.PatientId= int(result["patientId"].(float64))
		cache.Upload.DialysisProcessId= int(result["id"].(float64))
	}

}
func main() {
	Logger.Println("************服务开始启动***************")
	//称重协程启动
	go serial.Listening()

	//设置访问的路由

	//获取病人信息
	http.HandleFunc("/getPatientInfo", getPatientInfo)
	//呼吸获取体重信息
	http.HandleFunc("/breathe", breathe)
	http.HandleFunc("/finished", finished)


	http.HandleFunc("/", index)
	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir("./html"))))

	err := http.ListenAndServe(":9091", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	//go serial.Listening()
}
