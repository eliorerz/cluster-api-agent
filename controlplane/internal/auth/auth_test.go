package auth

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAuth(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecs(t, "Auth Suite")
}

var _ = Describe("GenerateFakePullSecret", func() {
	const (
		pullSecretName      = "test-pull-secret"
		pullSecretNamespace = "test-namespace"
	)

	It("should successfully generate a fake pull secret", func() {
		secret := GenerateFakePullSecret(pullSecretName, pullSecretNamespace)
		Expect(secret).NotTo(BeNil())
		Expect(secret.Data).To(HaveKey(".dockerconfigjson"))
	})
})
