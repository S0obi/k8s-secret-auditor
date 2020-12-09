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
	"k8s.io/client-go/tools/clientcmd"
)

// PasswordInfo : Password information
type PasswordInfo struct {
	name      string
	value     string
	compliant bool
	entropy   float64
}

// NewPasswordInfo : Constructor of PasswordInfo struct
func NewPasswordInfo(key string, value string) *PasswordInfo {
	s := PasswordInfo{name: key, value: value, entropy: utils.ComputeEntropy(value)}
	return &s
}

// Audit : Audit Kubernetes secrets
func Audit(config *config.Config, namespace string) {
	clientset := initClient(namespace)
	fmt.Printf("Audit secrets containing passwords\n\n")
	secrets, err := clientset.CoreV1().Secrets(namespace).List(context.TODO(), metav1.ListOptions{FieldSelector: "metadata.namespace!=kube-system,type==Opaque"})
	if err != nil {
		fmt.Println(err)
		return
	}
	var results [][]string
	for _, secret := range secrets.Items {
		passwordInfo := evaluateSecret(&secret, config)
		if passwordInfo != nil && !passwordInfo.compliant {
			secretName := fmt.Sprintf("%s/%s", secret.Namespace, passwordInfo.name)
			info := fmt.Sprintf("entropy=%f, length=%d", passwordInfo.entropy, len(passwordInfo.value))
			results = append(results, []string{secretName, passwordInfo.value, info})
		}
	}
	fmt.Printf("%d/%d are not compliant to the policy\n", len(results), len(secrets.Items))
	utils.PrintResultTable(results)
}

func evaluateSecret(secret *v1.Secret, config *config.Config) *PasswordInfo {
	for key, value := range secret.Data {
		if utils.IsPassword(string(key)) {
			s := NewPasswordInfo(key, string(value))
			s.compliant = isCompliant(s, config)
			return s
		}
	}
	return nil
}

func isCompliant(passwordInfo *PasswordInfo, config *config.Config) bool {
	if len(passwordInfo.value) < config.Policy.Length || passwordInfo.entropy < config.Policy.Entropy {
		return false
	} else {
		return true
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
