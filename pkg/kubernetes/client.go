package kubernetes

import (
	"fmt"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var Clientset *kubernetes.Clientset

func InitKubeClient() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("error getting user home dir: %w", err)
	}
	kubeConfigPath := filepath.Join(homeDir, ".kube", "config")
	fmt.Printf("Using kube config found at %s\n", kubeConfigPath)

	kubeConfig, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		return fmt.Errorf("error getting Kubernetes config: %w", err)
	}

	Clientset, err = kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		return fmt.Errorf("error getting Kubernetes clientset: %w", err)
	}

	return nil
}
