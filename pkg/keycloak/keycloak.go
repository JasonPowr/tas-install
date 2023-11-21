package keycloak

import (
	"fmt"
	"os/exec"
	"work/securesign/tas-install/pkg/kubernetes"
)

func InstallSSOKeycloak() error {

	fmt.Println("Installing keycloak operator")
	if err := applyKustomize("pkg/keycloak/operator/base"); err != nil {
		return err
	}
	if err := kubernetes.WaitForPodStatusRunning("keycloak-system", "rhsso-operator"); err != nil {
		return err
	}

	fmt.Println("Installing keycloak resources")
	if err := applyKustomize("pkg/keycloak/resources/base"); err != nil {
		return err
	}
	if err := kubernetes.WaitForPodStatusRunning("keycloak-system", "keycloak-postgresql"); err != nil {
		return err
	}
	fmt.Println("Keycloak has successfully been installed")
	return nil
}

func applyKustomize(path string) error {
	cmd := exec.Command("oc", "apply", "--kustomize", path)
	err := cmd.Run()
	return err
}
