package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	"k8s.io/client-go/tools/clientcmd"
)

type Query struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	Key       string `json:"key"`
}

type Result struct {
	Value string `json:"value"`
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err.Error())
	}

	var kubeconfig string
	if kubeconfig = os.Getenv("KUBECONFIG"); kubeconfig == "" {
		kubeconfig = filepath.Join(os.Getenv("HOME"), ".kube", "config")
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	var q Query
	err = json.Unmarshal(input, &q)
	if err != nil {
		panic(err.Error())
	}

	secret, err := clientset.CoreV1().Secrets(q.Namespace).Get(q.Name, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}

	val, ok := secret.Data[q.Key]
	if !ok {
		panic("Key not found")
	}

	r := &Result{
		Value: string(val),
	}

	o, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err.Error())
	}

	os.Stdout.Write(o)

	return
}
