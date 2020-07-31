package helpers

import (
	. "github.com/onsi/gomega"
	"testing"
	"fmt"
)

func TestParseUInt64(t *testing.T){
	g := NewGomegaWithT(t)

	stringToUint, err := ParseUInt64("5")

	if err != nil {
		fmt.Errorf("could not parse string to int")
	}
	
	g.Expect(stringToUint).To(BeIdenticalTo(uint64(5)), "String should have become a int")
}