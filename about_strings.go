package about_strings

import "./koans"
import "strings"
import "fmt"

func TestDoubleQuotedStringsAreStrings(t *koans.T) {
	str := "hello" //String can be defined with literal

	t.AssertEquals("hello", str)
}

func TestBackQuotedStringsAreStrings(t *koans.T) {
     // CHEN apparently everything is literal
	str := `hello\n
world`

	t.AssertEquals(13, len(str))
}

func TestPlusConcatenatesString(t *koans.T) {
	str := "Hello " + "World"
	t.AssertEquals("Hello World", str)
}

func TestPlusWillNotModifyOriginalStrings(t *koans.T) {
	hi := "Hello, "
	there := "world"
	str := hi + there
	t.AssertEquals("Hello, ", hi)
	t.AssertEquals("world", there)
	t.AssertEquals("Hello, world", str)
}

func TestPlusEqualsWillAppendToEndOfString(t *koans.T) {
	hi := "Hello, "
	there := "world"
	hi += there
	t.AssertEquals("Hello, world", hi)
}

func TestPlusEqualsAlsoLeavesOriginalStringUnmodified(t *koans.T) {
	original := "Hello, "
	hi := original
	there := "world"
	hi += there
	t.AssertEquals("Hello, ", original)
}

func TestUseSprintfToInterpolateVaribales(t *koans.T) {
	value1 := "one"
	value2 := 2
	str := fmt.Sprintf("The values are %s and %d", value1, value2)
	t.AssertEquals(koans.String__, str)
}

func TestYouCanGetASubstringFromAString(t *koans.T) {
	str := "Bacon, lettuce and tomato"
	t.AssertEquals(koans.String__, str[7:10])
}

func TestYouCanGetASingleCharacterFromAString(t *koans.T) {
	str := "Bacon, lettuce and tomato"
	t.AssertTrue(koans.Char__ == str[1])
}

func TestCharactersAreBytesActually(t *koans.T) {
	t.AssertTrue(koans.Char__ == 'a'+1)
}

func TestStringsCanBeSplit(t *koans.T) {
	str := "Sausage Egg Cheese"
	words := strings.Split(str, " ")
	t.AssertEqualsStringSlices(koans.StringSlice__, words)
}

func TestStringsCanBeSplitWithDifferentPatterns(t *koans.T) {
	str := "the,|;rain;in,spain"
	words := strings.Split(str, ",|;")
	t.AssertEqualsStringSlices(koans.StringSlice__, words)
}
