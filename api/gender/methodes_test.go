package gender

import "testing"

func TestGender_IsValid(t *testing.T) {
    testCases := []struct {
        gender Gender
        isValid bool
    }{
        {Male, true},
        {Female, true},
        {OtherGender, true},
        {NoGender, true},
        {Gender("InvalidGender"), false},
    }

    for _, tc := range testCases {
        isValid := tc.gender.IsValid()
        if isValid != tc.isValid {
            t.Errorf("Expected %v.IsValid() to be %v, but got %v", tc.gender, tc.isValid, isValid)
        }
    }
}

