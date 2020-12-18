package commands

import (
	"context"
	"fmt"
	"os"

	"github.com/S0obi/k8s-secret-auditor/pkg/config"
	"github.com/S0obi/k8s-secret-auditor/pkg/password"
	"github.com/S0obi/k8s-secret-auditor/pkg/utils"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// Audit : Audit Kubernetes secrets
func Audit(config *config.Config, namespace string) {
	clientset := initClient(namespace)
	fmt.Printf("Audit secrets containing passwords\n\n")

	ignoredNamespaces := ""
	for _, ignoredNamespace := range config.IgnoredNamespaces {
		ignoredNamespaces += fmt.Sprintf("metadata.namespace!=%s,", ignoredNamespace)
	}

	secrets, err := clientset.CoreV1().Secrets(namespace).List(context.TODO(), metav1.ListOptions{FieldSelector: ignoredNamespaces + "type==Opaque"})
	if err != nil {
		fmt.Println(err)
		return
	}

	var results [][]string
	for _, secret := range secrets.Items {
		passwordInfo := evaluateSecret(&secret, config)
		if passwordInfo != nil && !passwordInfo.Compliant {
			secretName := fmt.Sprintf("%s/%s", secret.Namespace, passwordInfo.Name)
			info := fmt.Sprintf("entropy=%f, length=%d", passwordInfo.Entropy, len(passwordInfo.Value))

			results = append(results, []string{secretName, passwordInfo.Value, info})
		}
	}

	fmt.Printf("%d/%d are not compliant to the policy\n", len(results), len(secrets.Items))
	utils.PrintResultTable(results)
}

func evaluateSecret(secret *v1.Secret, config *config.Config) *password.Password {
	for key, value := range secret.Data {
		if password.IsPassword(string(key), config) {
			s := password.NewPassword(key, string(value))
			s.Compliant = s.IsCompliant(config)
			return s
		}
	}
	return nil
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
