package repl

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  HELLO  WORld  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  what  beauTifuL tests  ",
			expected: []string{"what", "beautiful", "tests"},
		},
		{
			input:    "  Testing is  quite  important   ",
			expected: []string{"testing", "is", "quite", "important"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("Lenght does not match. Expected: %d, Actual: %d", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test

			if word != expectedWord {
				t.Errorf("\nExpected: %v\nGot: %v", expectedWord, word)
			}
		}
	}
}
