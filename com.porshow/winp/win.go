package winp

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)
import (
	"log"

)
var LabelMap =make(map[string]*walk.Label)

type object struct {
	name string
	id string
	weightNow string
}


func Invoke() {

	var mw *walk.MainWindow
	var label *walk.Label
	LabelMap["name"]=label


	if _, err := (MainWindow{
		AssignTo: &mw,
		Title:    "葆秀-体重读取工具",
		MinSize:  Size{300, 300},
		Layout:   VBox{},
		Children: []Widget{
			Label{
				Text: "姓名:",
			},
			Label{
				AssignTo:&label,
				Text: "姓名",
			},
			Label{
				Text: "透析id:",
			},
			Label{
				Text: "目前体重:",
			},
			Label{
				Text: "animal:",
			},
			Label{
				Text: "animal:",
			},
			Label{
				Text: "animal:",
			},
		},
	}.Run()); err != nil {
		log.Fatal(err)
	}
}



