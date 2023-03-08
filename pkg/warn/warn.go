package warn

import (
	"github.com/golang/glog"
)

func EXIT() {
	glog.Errorf("White DB EXIT!")
}

func ERRORF(msg string) {
	glog.Errorf("is error! msg : [%s]", msg)
}

func ERROR() {
	glog.Errorf("this is a error command!")
}
