package postgres_test

import (
	"os"
	"strconv"

	"github.com/nvasilev98/rentals/pkg/repository/postgres"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type Config struct {
	Host     string `envconfig:"DB_HOST" required:"true"`
	Port     int    `envconfig:"DB_PORT" required:"true"`
	Username string `envconfig:"DB_USERNAME" required:"true"`
	Password string `envconfig:"DB_PASSWORD" required:"true"`
	Name     string `envconfig:"DB_NAME" required:"true"`
}

var _ = Describe("Env", func() {
	const (
		hostEnv     = "DB_HOST"
		portEnv     = "DB_PORT"
		usernameEnv = "DB_USERNAME"
		passwordEnv = "DB_PASSWORD"
		nameEnv     = "DB_NAME"
		host        = "127.0.0.1"
		port        = 8080
		username    = "db-user"
		password    = "db-password"
		name        = "db-name"
	)

	BeforeEach(func() {
		Expect(os.Setenv(hostEnv, host)).To(Succeed())
		Expect(os.Setenv(portEnv, strconv.Itoa(port))).To(Succeed())
		Expect(os.Setenv(usernameEnv, username)).To(Succeed())
		Expect(os.Setenv(passwordEnv, password)).To(Succeed())
		Expect(os.Setenv(nameEnv, name)).To(Succeed())
	})

	AfterEach(func() {
		Expect(os.Unsetenv(hostEnv)).To(Succeed())
		Expect(os.Unsetenv(portEnv)).To(Succeed())
		Expect(os.Unsetenv(usernameEnv)).To(Succeed())
		Expect(os.Unsetenv(passwordEnv)).To(Succeed())
		Expect(os.Unsetenv(nameEnv)).To(Succeed())
	})

	When("host environment is missing", func() {
		BeforeEach(func() {
			Expect(os.Unsetenv(hostEnv)).To(Succeed())
		})

		It("should return an error", func() {
			_, err := postgres.LoadDBConfig()
			Expect(err).To(HaveOccurred())
		})
	})

	When("port is invalid", func() {
		BeforeEach(func() {
			Expect(os.Setenv(portEnv, "invalid-port")).To(Succeed())
		})

		It("should return an error", func() {
			_, err := postgres.LoadDBConfig()
			Expect(err).To(HaveOccurred())
		})
	})

	When("database configuration is provided", func() {
		var expectedDBConfig = postgres.Config{
			Host:     host,
			Port:     port,
			Username: username,
			Password: password,
			Name:     name,
		}

		It("should load it", func() {
			dbConfig, err := postgres.LoadDBConfig()
			Expect(err).ToNot(HaveOccurred())
			Expect(dbConfig).To(Equal(expectedDBConfig))
		})
	})
})
