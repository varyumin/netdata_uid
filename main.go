package main

import (
	"hetdata_uid/pkg"
	"time"
)

func main() {
	k8s := pkg.GetConnect()
	pkg.GetUID(k8s)
	for {
		pkg.GetSelfUID()
		time.Sleep(1 * time.Second)
	}

}
