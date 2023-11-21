package main

import (
	"fmt"
	"log"
	"work/securesign/tas-install/pkg/keycloak"
	"work/securesign/tas-install/pkg/kubernetes"
)

func main() {
	// Initialize the Kubernetes client
	err := kubernetes.InitKubeClient()
	if err != nil {
		log.Fatalf("Failed to initialize Kubernetes client: %v", err)
	}

	if err := keycloak.InstallSSOKeycloak(); err != nil {
		fmt.Println(err.Error())
	}
}
