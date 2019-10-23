package tests

import "github.com/revel/revel/testing"

type ItemInfoTest struct {
	testing.TestSuite
}

func (t *ItemInfoTest) Before() {
	println("Set up")
}

///////////////////////// Mocks //////////////////////////
////////////////////// Unit Tests ////////////////////////
////////////////// Functional Tests //////////////////////

func (t *ItemInfoTest) After() {
	println("Tear down")
}
