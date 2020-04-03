package pkg

import (
	"io/ioutil"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"log"
	"os"
)

func getPodName () string{
	podName, _ := os.Hostname()
	return podName
}

func getNamespace() string{
	ns, err := os.Stat("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
	if err != nil && ns.IsDir(){
		log.Fatalf("Not find /var/run/secrets/kubernetes.io/serviceaccount/namespace")
	}

	namespace, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
	if err != nil {
		log.Fatalf("Cann't open and read file")
	}

	return string(namespace)
}

func GetUID( k8sConnect *kubernetes.Clientset) (uid string) {
	pod, err := k8sConnect.CoreV1().Pods(getNamespace()).Get(getPodName(),metav1.GetOptions{})
	if err != nil {
		log.Fatal(err)
	}

	node, err :=k8sConnect.CoreV1().Nodes().Get(pod.Spec.NodeName, metav1.GetOptions{})
	if err != nil {
		log.Fatal(err)
	}

	if node.Status.NodeInfo.SystemUUID != "" {
		return node.Status.NodeInfo.SystemUUID
	}
	return string(node.ObjectMeta.UID)
}

func WriteUID(uid string)  bool{

	return true
}