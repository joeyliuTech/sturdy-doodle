/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/*
modification history
--------------------
2021/10/19,  create
*/
/*
DESCRIPTION
TODO mobile_test.go
*/
package testdata_generator

import "testing"

func TestGenerateMobilePhoneNum(t *testing.T) {
	mobileNum := Gen()
	if "" == mobileNum {
		t.Error("phone numbers is empty")
	}

	if len(mobileNum) != 11 {
		t.Error("phone numbers is invalid", mobileNum)
	}

	t.Log("generate mobile num:", mobileNum)
}