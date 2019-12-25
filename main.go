package main

import (
	"code.izhaohu.com/dolphin/demo/protoc-gen-enum/enumext"
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

func main() {
	pgs.Init(
		pgs.DebugEnv("DEBUG"),
	).RegisterModule(
		enumext.EnumExt(),
	).RegisterPostProcessor(
		pgsgo.GoFmt(),
	).Render()
}
