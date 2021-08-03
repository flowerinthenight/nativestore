package nativestore

import (
	"testing"
)

func TestSetGet(t *testing.T) {
	lbl := "nativestore"
	url := "https://github.com/flowerinthenight/nativestore"
	err := Set(lbl, url, "user", "password")
	if err != nil {
		t.Fatal(err)
	}

	user, secret, err := Get(lbl, url)
	if err == nil {
		if user != "user" {
			t.Errorf("Expecting user, got %s", user)
		}

		if secret != "password" {
			t.Errorf("Expecting password, got %s", secret)
		}
	} else {
		t.Log("got error:", err)
	}
}
