package main

import "testing"

func Test_GetPersonID(t *testing.T) {
	var persons person
	personExists, selectedPerson, _ := getPersonID("1")
	if persons != selectedPerson {
		t.Fatalf("selected person is not empty when selecting from an empty set")
	}
	if personExists {
		t.Fatalf("got personExists is not false when selecting from an empty set")
	}
}
