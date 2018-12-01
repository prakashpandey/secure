package secure

import (
	"testing"
)

func TestSha512HashGeneration(t *testing.T) {
	input := "Beautiful Days"
	// base64 encoded
	expectedOutput := "94009566dec17318603ff29cbc359f39856b1d40f6f2ffce0394c56f3dc8530038c0212e631b7f95a5619425213821aff3820365d8826690d0bba92765f4fbdb"
	actualOutput, err := Hash(input, SHA512)
	if err != nil {
		t.Errorf("test failed: %s", err.Error())
	}
	if actualOutput != expectedOutput {
		t.Errorf("Expedted[%s], Actual[%s]", expectedOutput, actualOutput)
	}
}

func TestMatchHash(t *testing.T) {
	input := "Beautiful Days"
	hashedInput := "94009566dec17318603ff29cbc359f39856b1d40f6f2ffce0394c56f3dc8530038c0212e631b7f95a5619425213821aff3820365d8826690d0bba92765f4fbdb"
	if !MatchHash(input, hashedInput, SHA512) {
		t.Error("Hash do not match")
	}
}
