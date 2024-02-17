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
