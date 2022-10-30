package env_test

import (
	"os"
	"strconv"

	"github.com/nvasilev98/rentals/cmd/rentals/env"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {

	const (
		hostEnv     = "HOST"
		portEnv     = "PORT"
		validHost   = "127.0.0.1"
		defaultHost = "localhost"
		validPort   = 8000
		defaultPort = 8080
	)

	When("environment is set", func() {
		BeforeEach(func() {
			Expect(os.Setenv(hostEnv, validHost)).To(Succeed())
			Expect(os.Setenv(portEnv, strconv.Itoa(validPort))).To(Succeed())
		})

		AfterEach(func() {
			Expect(os.Unsetenv(hostEnv)).To(Succeed())
			Expect(os.Unsetenv(portEnv)).To(Succeed())
		})

		It("host and port are provided", func() {
			config, err := env.LoadAppConfig()
			Expect(err).ToNot(HaveOccurred())
			Expect(config.Host).To(Equal(validHost))
			Expect(config.Port).To(Equal(validPort))
		})
	})

	When("env is not set", func() {
		It("should assign default values", func() {
			config, err := env.LoadAppConfig()
			Expect(err).ToNot(HaveOccurred())
			Expect(config.Host).To(Equal(defaultHost))
			Expect(config.Port).To(Equal(defaultPort))
		})
	})

})
