package learnlab

import (
	"fmt"
	"strings"
)

func StringUtilities() {

	// strings.split functions splits based on delimiter
	st := "apple-banana-golang"
	fmt.Println(strings.Split(st, "-"))

	// strings.join function joins based on delimiter
	towns := []string {"Rome", "Baramati", "Lucknow", "Kotdwar"}
	joined_towns := strings.Join(towns, "::")
	fmt.Println(joined_towns)

	// strings.contains returns true or false for substring
	fmt.Println(strings.Contains(joined_towns, "::Baramati"))

	// strings.replace Replace returns a copy of the string s with the first n non-overlapping instances of old replaced by new.
	fmt.Println(strings.Replace(joined_towns, "::", "  ||  ", 2))

	// Trim Space, To Lower and To Upper Methods
	stSpace := "    Hello There :)          "
	fmt.Println(strings.TrimSpace(stSpace))
	fmt.Println(strings.ToLower(stSpace))
	fmt.Println(strings.ToUpper(stSpace))

	// Repeat substrings
	fmt.Println(strings.Repeat("Foo-BarTaz", 4))

	// HasPrefix , HasSuffix and Count occurences function
	st = "WWE_WWF_WWW"
	fmt.Println(strings.HasSuffix(st, "WWW"))
	fmt.Println(strings.HasPrefix(st, "WWE_WWF_"))
	fmt.Println(strings.Count(st, "WW"))

}