package about_asserts

import "./koans"

/* 
    This unit testing infrastructure doesn't corellate gotest infrastructure.
    Please refer to golang.org for documentation about gotest. Here we are using
    unit testing infrastrucutre special for go koans.

    This file contains koans for go koans unit testing infrastructure.
    */
func TestAssertTruth(t *koans.T) {
	t.AssertTrue(true) // This should be true
}

/*  Go doesn't allow method and function overloading. That is why 
    AssertTrue with message argument should have another name.
    */
func TestAssertWithMessage(t *koans.T) {
	t.AssertTrueWithMessage(true, "This should be true. Please fix.") // This should be true
}

/*  Function FailNow stops execution of the current test. */
func TestFillInValues(t *koans.T) {
	t.AssertEquals(2, 1 + 1)
}

func TestAssertEquality(t *koans.T) {
	expected_value := 2
	actual_value := 1 + 1
	t.AssertTrue(expected_value == actual_value)
}

func TestABetterWayOfAssertingEquality(t *koans.T) {
	expected_value := 2
	actual_value := 1 + 1
	t.AssertEquals(expected_value, actual_value)
}
