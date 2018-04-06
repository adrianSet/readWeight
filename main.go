package main

import (
	"net/http"
	"log"
	"fmt"
	"html/template"
	"readWeight/config"
	"io/ioutil"
   ."readWeight/log1"
)
var (
	patientInfoUrl = config.M["host"]+config.M["patient_info_url"]
)


func init(){


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

func index(w http.ResponseWriter, r *http.Request){
	tmp, err := template.ParseFiles("html/index.html")
	locals := make(map[string]interface{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmp.Execute(w, locals)

}

func getPatientInfo(w http.ResponseWriter, r *http.Request){
	id:=r.FormValue("id")
	resp, err := http.Get(patientInfoUrl+"?dialysisCode="+id)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	json:=string(body)
	fmt.Println(json)
	fmt.Fprintf(w,json)
}


func main(){
	Logger.Println("************服务开始启动***************")
	//设置访问的路由
	http.HandleFunc("/", index)
	http.HandleFunc("/getPatientInfo", getPatientInfo)

	http.Handle("/html/",http.StripPrefix("/html/", http.FileServer(http.Dir("./html"))))

	err := http.ListenAndServe(":9091", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	//go serial.Listening()
}

