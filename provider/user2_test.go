package main

import (
	"fmt"
	"testing"
)

type mockProvider struct {
	mockInfo Information
	err      error
}

func (m mockProvider) GetInfo() (Information, error) {
	return m.mockInfo, m.err
}

func TestGetInfo(t *testing.T) {
	// Here, I set the mocker to provide pre-configure data
	m := mockProvider{mockInfo: Information{}, err: fmt.Errorf("Some error")}

	user := User2{MyProvider: &m}
	info, err := user.ProcessInfo()
	if err != nil {
		t.Fatal(err)
	}
	if (info != Information{}) {
		t.Fatal("Info not expected")
	}

}
