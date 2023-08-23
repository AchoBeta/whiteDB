package warn

import (
	"fmt"

	"github.com/golang/glog"
)

func EXIT() {
	fmt.Printf(">>>>>>>>>> NekoKV EXIT!\n")
}

func ERRORF(msg string) {
	glog.Errorf("is error! msg : [%s]", msg)
}

func ERROR() {
	glog.Errorf("this is a error command!")
}

func DEBUG(f interface{}) {
	glog.Infof("%+v\n", f)
}
