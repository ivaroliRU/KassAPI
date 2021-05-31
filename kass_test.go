package KassAPI

import (
	"testing"
)

func TestCreateChargeSuccess(t *testing.T) {
	c := New(false, "kass_test_auth_token", "")
	response1, err := c.CreateCharge(2199, "Kass bolur", "https://photos.kassapi.is/kass/kass-bolur.jpg", "ABC12332öiö23iö3", "1001001", 1, 90, "https://example.com/callbacks/kass")

	if !response1.Success || err != nil {
		t.Errorf("Failed response from sandbox, wanted success")
	}
}

func TestCreateChargeBadRecipient(t *testing.T) {
	c := New(false, "kass_test_auth_token", "")
	response1, _ := c.CreateCharge(2199, "Kass bolur", "https://photos.kassapi.is/kass/kass-bolur.jpg", "ABC12332öiö23iö3", "1234123", 1, 90, "https://example.com/callbacks/kass")

	if response1.Success {
		t.Errorf("Success response from sandbox, wanted failed (wrong recipient)")
	}
}

func TestBadToken(t *testing.T) {
	c := New(false, "asdf", "")
	response1, err := c.CreateCharge(2199, "Kass bolur", "https://photos.kassapi.is/kass/kass-bolur.jpg", "ABC12332öiö23iö3", "1001001", 1, 90, "https://example.com/callbacks/kass")

	if response1.Success || err != nil {
		t.Errorf("Success response from sandbox, wanted failed (wrong access token)")
	}
}
