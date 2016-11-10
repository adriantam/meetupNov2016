package main

import "testing"

type mockProviderBad struct {
	mockInfo Information
	err      error
}

func (m mockProviderBad) GetInfo() (Information, int, error) {
	return m.mockInfo, 0, m.err
}

func TestGetInfoBad(t *testing.T) {
	m := mockProviderBad{}
	user := User2{MyProvider: &m}
	info, err := user.ProcessInfo()
	if err != nil {
		t.Fatal(err)
	}
	if (info != Information{}) {
		t.Fatal("Info not expected")
	}

}
