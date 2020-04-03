package pkg

import (
	"io"
	"io/ioutil"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"log"
	"os"
	"strings"
)

func getPodName() string {
	podName, _ := os.Hostname()
	return podName
}

func getNamespace() string {
	ns, err := os.Stat("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
	if err != nil && ns.IsDir() {
		log.Fatalf("Not find /var/run/secrets/kubernetes.io/serviceaccount/namespace")
	}

	namespace, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
	if err != nil {
		log.Fatalf("Cann't open and read file")
	}

	return string(namespace)
}

func GetUID(k8sConnect *kubernetes.Clientset) bool {
	var uid string
	pod, err := k8sConnect.CoreV1().Pods(getNamespace()).Get(getPodName(), metav1.GetOptions{})
	if err != nil {
		log.Fatal(err)
		return false
	}

	node, err := k8sConnect.CoreV1().Nodes().Get(pod.Spec.NodeName, metav1.GetOptions{})
	if err != nil {
		log.Fatal(err)
		return false
	}

	if node.Status.NodeInfo.SystemUUID != "" {
		uid = node.Status.NodeInfo.SystemUUID
	}
	uid = string(node.ObjectMeta.UID)

	if err = writeUIDToFile("/tmp/uid", uid); err != nil {
		return false
	}
	return true
}

func writeUIDToFile(filepath, s string) error {
	fo, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer fo.Close()

	_, err = io.Copy(fo, strings.NewReader(s))
	if err != nil {
		return err
	}

	return nil
}
