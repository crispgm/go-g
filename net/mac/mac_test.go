package mac

import "testing"

func TestGetMac(t *testing.T) {
	addr, err := GetMacAddr()
	if err == nil {
		t.Log(addr)
	} else {
		t.Error(err)
	}
}
