package service

import (
	"testing"

	"github.com/JUkhan/goapp/db"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestVideoService(t *testing.T) {
	db.InitDB()

	RegisterFailHandler(Fail)
	RunSpecs(t, "Video Service Test Suite")
}
