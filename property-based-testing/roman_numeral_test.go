package roman

import (
	"fmt"
	"testing"
	"testing/quick"
)

var tt = []struct {
	Arabic uint16
	Roman  string
}{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{7, "VII"},
	{8, "VIII"},
	{9, "IX"},
	{10, "X"},
	{14, "XIV"},
	{18, "XVIII"},
	{20, "XX"},
	{39, "XXXIX"},
	{40, "XL"},
	{47, "XLVII"},
	{49, "XLIX"},
	{50, "L"},
	{90, "XC"},
	{94, "XCIV"},
	{99, "XCIX"},
	{100, "C"},
	{400, "CD"},
	{500, "D"},
	{900, "CM"},
	{1000, "M"},
	{1984, "MCMLXXXIV"},
	{3999, "MMMCMXCIX"},
	{2014, "MMXIV"},
	{1006, "MVI"},
	{798, "DCCXCVIII"},
}

func TestROmanNumerals(t *testing.T) {

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%d gets converted to %q", tc.Arabic, tc.Roman), func(t *testing.T) {
			got, _ := ConvertToRoman(tc.Arabic)

			if got != tc.Roman {
				t.Errorf("got %q, want %q", got, tc.Roman)
			}
		})

	}

}

func TestConvertingToArabic(t *testing.T) {
	for _, tc := range tt {
		t.Run(fmt.Sprintf("%q gets converted to %d", tc.Roman, tc.Arabic), func(t *testing.T) {
			got := ConvertToArabic(tc.Roman)
			if got != tc.Arabic {
				t.Errorf("got %d, want %d", got, tc.Arabic)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {

		if arabic > 3999 {
			return true
		}
		t.Log("testing", arabic)
		roman, _ := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, &quick.Config{MaxCount: 1000}); err != nil {
		t.Error("failed checks", err)
	}
}

func TestMaxRomanValue(t *testing.T) {
	// maximum value you can represent in roman system is 3999
	arabic := 4000
	_, err := ConvertToRoman(uint16(arabic))

	if err == nil {
		t.Error("error shouldn't be nil")
	}

}
