package kgoutil

import "testing"

var localAddress = "192.168.0.116"

func TestGetHostIP(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		// change to your own local ip address.
		{"TestGetHostIP", localAddress},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetHostIP(); got != tt.want {
				t.Errorf("GetHostIP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetHostPort(t *testing.T) {
	tests := []struct {
		name     string
		addr     string
		wantHost string
		wantPort string
	}{
		// TOD: Add test cases.
		{"TestGetHostPort", ":9090", localAddress, "9090"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHost, gotPort := GetHostPort(tt.addr)
			if gotHost != tt.wantHost {
				t.Errorf("GetHostPort() returns Host: %v, want %v.", gotHost, tt.wantHost)
			}

			if gotPort != tt.wantPort {
				t.Errorf("GetHostPort() returns Port: %v, want %v.", gotPort, tt.wantPort)
			}
		})
	}
}
