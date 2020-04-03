package main

import (
	"fmt"
	"hetdata_uid/pkg"
	"time"
)

func main() {
	k8s := pkg.GetConnect()
	pkg.GetUID(k8s)
	fmt.Printf("POD_NAME: %s\n", pkg.GetPodName())
	fmt.Printf("NAMESPACE_NAME: %s\n", pkg.GetNamespace())

	for {
		pkg.GetSelfUID()
		time.Sleep(1 * time.Second)
	}

}
