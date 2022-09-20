package put_test

import "testing"

func oddOrEven(num int) string {
	if num%2 == 0 {
		return "This is even"
	} else {
		return "This is odd"
	}
}

type testCase struct {
	name string
	in   int
	want string
}

func TestOddOrEven(t *testing.T) {

	tc := testCase{
		name: "even",
		in:   2,
		want: "This is even",
	}

	makeAssertion(t, tc)
}

func makeAssertion(t *testing.T, tc testCase) {
	t.Run(tc.name, func(t *testing.T) {
		got := oddOrEven(tc.in)
		if got != tc.want {
			t.Errorf("Got '%s' but we wanted '%s', the input was %d", got, tc.want, tc.in)
		}
	})
}

func TestOddOrEven2(t *testing.T) {

	tt := []struct {
		name string
		in   int
		want string
	}{
		{name: "even", in: 2, want: "This is even"},
		{name: "odd", in: 1, want: "This is odd"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := oddOrEven(tc.in)
			if got != tc.want {
				t.Errorf("Got '%s' but we wanted '%s', the input was %d", got, tc.want, tc.in)
			}
		})
	}
}
