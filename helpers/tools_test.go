package helpers

import (
	"github.com/onsi/gomega"
)

func TestParseUInt64(t *testing.T){
	g := NewGomegaWithT(t)

	stringToUint, err := ParseUInt64("5")
	
	g.Expect(stringToUint).To(BeString(), "String should have become a int")
}