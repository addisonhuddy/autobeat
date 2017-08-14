package main_test

import (
	. "github.com/addisonhuddy/autobeat"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"path/filepath"
)

func fixturePath(filename string) string {
	path, err := filepath.Abs(filepath.Join("fixtures", filename+".json"))
	Ω(err).ShouldNot(HaveOccurred())
	return path
}

var _ = Describe("Autobeat", func() {
	Describe("JSON parsing functionality", func() {
		It("Prints out a list of stories", func() {
			jsonFile, _ := ioutil.ReadFile(fixturePath("current"))
			_, err := ReadStoryResponse(jsonFile)
			Expect(err).ToNot(HaveOccurred())
		})
	})

})
