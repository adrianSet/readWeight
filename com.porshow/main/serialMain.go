package main


import (
	"log"
	"readWeight/com.porshow/winp"
	"time"
	"fmt"
)

func main(){
	log.Println("```````````````````````````running```````````````````````````````")
	//go serial.Listening()
	go func(){
		time.Sleep(10*time.Second)
		label:=winp.LabelMap["name"]
		label.SetText("2222")
		fmt.Println("已替换")
	}()
	winp.Invoke()
}

//func main() {
//	var inTE, outTE *walk.TextEdit
//
//	MainWindow{
//		Title:   "SCREAMO",
//		MinSize: Size{600, 400},
//		Layout:  VBox{},
//		Children: []Widget{
//			HSplitter{
//				Children: []Widget{
//					TextEdit{AssignTo: &inTE},
//					TextEdit{AssignTo: &outTE, ReadOnly: true},
//				},
//			},
//			PushButton{
//				Text: "SCREAM",
//				OnClicked: func() {
//					outTE.SetText(strings.ToUpper(inTE.Text()))
//				},
//			},
//		},
//	}.Run()
//}