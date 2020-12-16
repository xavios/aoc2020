package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(t *testing.T) {
	t.Run("One passport", func(t *testing.T) {
		serialized := `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
		byr:1937 iyr:2017 cid:147 hgt:183cm`

		var p []Passport = ParsePassports(serialized)
		require.Equal(t, 1, len(p))
		require.Equal(t, "gry", p[0].EyeColor)
		require.Equal(t, "860033327", p[0].PassportID)
		require.Equal(t, 2020, p[0].ExpirationYear)
		require.Equal(t, "#fffffd", p[0].HairColor)
		require.Equal(t, 1937, p[0].BirthYear)
		require.Equal(t, 2017, p[0].IssueYear)
		require.Equal(t, "147", p[0].CountryID)
		require.Equal(t, "183cm", p[0].Height)
		require.True(t, p[0].IsValid())
	})

	t.Run("Two passports", func(t *testing.T) {
		s := `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
		byr:1937 iyr:2017 cid:147 hgt:183cm
		
		iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
		hcl:#cfa07d byr:1929`

		p := ParsePassports(s)
		require.Equal(t, 2, len(p))
		require.Equal(t, 1, CountValids(p))
	})
}
