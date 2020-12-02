package main

import "testing"

func TestPassword(t *testing.T) {
	t.Run("first task - old password rule", func(t *testing.T) {
		cases := map[string]struct {
			pass Password
			want bool
		}{
			"valid password": {
				Password{Min: 1, Max: 3, Pass: "abcde", Char: "a"},
				true,
			},
			"invalid password": {
				Password{Min: 1, Max: 3, Pass: "abcde", Char: "f"},
				false,
			},
			"test from file": {
				// 14-15 h: xpwhdjhhjhrdjkhfh
				Password{Min: 14, Max: 15, Pass: "xpwhdjhhjhrdjkhfh", Char: "h"},
				false,
			},
		}

		for _, tc := range cases {
			got := IsPassowrdValid(tc.pass)
			if got != tc.want {
				t.Errorf("Got: %+v, Want: %+v", got, tc.want)
			}
		}
	})

	t.Run("second task - new rule", func(t *testing.T) {
		cases := map[string]struct {
			Pass Password
			Want bool
		}{
			"valid pass": {
				Password{Min: 1, Max: 3, Pass: "abcde", Char: "a"},
				true,
			},
			"easy invalid": {
				Password{Min: 1, Max: 2, Pass: "aa", Char: "a"},
				false,
			},
			"invalid pass": {
				Password{Min: 1, Max: 3, Pass: "cdefg", Char: "b"},
				false,
			},
		}

		for _, tc := range cases {
			got := IsPassValidNewRule(tc.Pass)
			if got != tc.Want {
				t.Errorf("Got: %+v, Want: %+v. %+v", got, tc.Want, tc.Pass)
			}
		}
	})
}
