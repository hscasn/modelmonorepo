package utils

import (
	"fmt"
	"io/ioutil"
)

func Must(bt interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return bt
}

func ReadModuleFile(module string) ([]byte, error) {
	return ioutil.ReadFile(fmt.Sprintf("../modules/%s.yaml", module))
}
