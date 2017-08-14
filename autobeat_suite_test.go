package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/gexec"
	"io/ioutil"
	"os"
	"testing"
)

func TestAutobeat(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Autobeat Suite")
}

var autobeat string
var workingDir string

var _ = BeforeSuite(func() {
	var err error
	autobeat, err = gexec.Build("github.com/addisonhuddy/autobeat")
	Ω(err).ShouldNot(HaveOccurred())
})

var _ = BeforeEach(func() {
	var err error
	workingDir, err = ioutil.TempDir("", "autobeat")
	Ω(err).ShouldNot(HaveOccurred())
})

var _ = AfterEach(func() {
	Ω(os.RemoveAll(workingDir)).Should(Succeed())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})
