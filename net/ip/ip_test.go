package ip

import "testing"

func TestGetLocalIP(t *testing.T) {
	t.Log(GetLocalIP())
}
