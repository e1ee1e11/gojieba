package gojieba

import (
	"github.com/e1ee1e11/gojieba/deps/cppjieba"
	"github.com/e1ee1e11/gojieba/deps/limonp"
	"github.com/e1ee1e11/gojieba/dict"
)

func init() {
	dict.Init()
	limonp.Init()
	cppjieba.Init()
}
