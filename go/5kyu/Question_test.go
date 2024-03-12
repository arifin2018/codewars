package kyu

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertStringToCamelCase(t *testing.T) {
	test1 := ConvertStringToCamelCase("The_Stealth_Warrior")
	assert.Equal(t, "TheStealthWarrior", test1)
	test2 := ConvertStringToCamelCase("the-stealth-Warrior")
	assert.Equal(t, "theStealthWarrior", test2)
	// test3 := ConvertStringToCamelCase("You_have_chosen_to_translate_this_kata_For_your_convenience_we_have_provided_the_existing_test_cases_used_for_the_language_that_you_have_already_completed_as_well_as_all_of_the_other_related_fields")
	// assert.Equal(t, "YouhavechosentotranslatethiskataForyourconveniencewehaveprovidedtheexistingtestcasesusedforthelanguagethatyouhavealreadycompletedaswellasalloftheotherrelatedfields", test3)
}
func TestConverta2(t *testing.T) {
	var text = "banana burger soup"
	var regex, _ = regexp.Compile(`[banana]+`)
	var res1 = regex.FindAllString(text, -1)
	fmt.Printf("%#v \n", res1)
}

func TestIpsBetween(t *testing.T) {
	// test1 := IpsBetween("10.0.0.0", "10.0.0.50")
	// assert.Equal(t, 50, test1)
	// test2 := IpsBetween("10.0.0.0", "10.0.1.0")
	// assert.Equal(t, 256, test2)
	// test3 := IpsBetween("20.0.0.10", "20.0.1.0")
	// assert.Equal(t, 246, test3)
	// test4 := IpsBetween("50.0.0.0", "50.1.1.1")
	// assert.Equal(t, 65793, test4)
	// test5 := IpsBetween("10.11.12.13", "10.11.13.0")
	// assert.Equal(t, 243, test5)
	// test6 := IpsBetween("143.218.244.41", "205.55.162.96")
	// assert.Equal(t, 102948439, test6)
}
