package xgo

import (
	"fmt"
	"github.com/gotomicro/ego/core/elog"
	"runtime"
)

func try(fn func() error, cleaner func()) (ret error) {
	if cleaner != nil {
		defer cleaner()
	}
	defer func() {
		if err := recover(); err != nil {
			_, file, line, _ := runtime.Caller(2)
			elog.Error("recover", elog.Any("err", err), elog.FieldAddr(fmt.Sprintf("%s:%d", file, line)))
			if _, ok := err.(error); ok {
				ret = err.(error)
			} else {
				ret = fmt.Errorf("%+v", err)
			}
		}
	}()
	return fn()
}

func try2(fn func(), cleaner func()) (ret error) {
	if cleaner != nil {
		defer cleaner()
	}
	defer func() {
		_, file, line, _ := runtime.Caller(2)
		if err := recover(); err != nil {
			elog.Error("recover", elog.Any("err", err), elog.FieldAddr(fmt.Sprintf("%s:%d", file, line)))
			if _, ok := err.(error); ok {
				ret = err.(error)
			} else {
				ret = fmt.Errorf("%+v", err)
			}
		}
	}()
	fn()
	return nil
}
