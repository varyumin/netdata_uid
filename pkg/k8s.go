package pkg

import (
	"fmt"
	"github.com/prometheus/common/log"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

func possible() bool {
	fi, err := os.Stat("/var/run/secrets/kubernetes.io/serviceaccount/token")
	return os.Getenv("KUBERNETES_SERVICE_HOST") != "" &&
		os.Getenv("KUBERNETES_SERVICE_PORT") != "" &&
		err == nil && !fi.IsDir()
}

func getConfigPOD() (connect *kubernetes.Clientset) {

	config, err := rest.InClusterConfig()
	if err != nil {

	}
	connect, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error get config: %v\n", err)

	}
	return connect
}

func getConfigHost() (connect *kubernetes.Clientset) {
	loader := clientcmd.NewDefaultClientConfigLoadingRules()
	config, err := clientcmd.BuildConfigFromKubeconfigGetter("", loader.Load)

	if err != nil {
		log.Errorf("ERROR: %v\n", err)
		os.Exit(1)
	}
	connect = kubernetes.NewForConfigOrDie(config)
	return connect
}

func GetConnect() (k8sConnect *kubernetes.Clientset) {
	if possible() {
		k8sConnect = getConfigPOD()
	} else {
		k8sConnect = getConfigHost()
	}
	return k8sConnect
}

func GetSelfUID() {
	for _, pair := range os.Environ() {
		fmt.Println(pair)
	}
}
