package heif

import "testing"

func TestGetVersion(t *testing.T) {
	t.Log("libheif version: ", GetVersion())
}
