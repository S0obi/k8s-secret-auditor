package commands

import (
	"context"
	"fmt"
	"os"

	"github.com/S0obi/k8s-secret-auditor/pkg/config"
	"github.com/S0obi/k8s-secret-auditor/pkg/utils"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	coreV1Types "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/clientcmd"
)

// API client for managing secrets
var secretsClient coreV1Types.SecretInterface

// Audit : Audit Kubernetes secrets
func Audit(config *config.Config, namespace string) {
	clientset := initClient(namespace)
	fmt.Printf("Audit for namespace %s\n", namespace)
	secrets, err := clientset.CoreV1().Secrets(namespace).List(context.TODO(), metav1.ListOptions{FieldSelector: "metadata.namespace!=kube-system,type==Opaque"})
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, secret := range secrets.Items {
		isSecretCompliant(&secret, config)
	}
}

func isSecretCompliant(secret *v1.Secret, config *config.Config) {
	for key, value := range secret.Data {
		if utils.IsPassword(string(key)) {
			password := string(value)
			entropy := utils.ComputeEntropy(password)
			if len(password) < config.Policy.Length || entropy < config.Policy.Entropy {
				fmt.Printf("%s is not compliant (entropy=%f, length=%d)\n", secret.GetName(), entropy, len(password))
			} else {
				fmt.Printf("%s is compliant (entropy=%f, length=%d)\n", secret.GetName(), entropy, len(password))
			}
		}
	}
}

func initClient(namespace string) *kubernetes.Clientset {
	kubeconfig := os.Getenv("HOME") + "/.kube/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientset
}
