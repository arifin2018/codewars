<?php 

function bump($x) {
    if (substr_count($x, 'n') > 15) {
        return "Car Dead";
    }else {
        # code...
        return "Woohoo!";
    }
}

echo bump("nnnnnnnnnnnnnnnnnnnnn");

// class MyTestCases extends TestCase
// {
//     public function testSampleTests() {
//         $this->assertEquals(bump("n"), "Woohoo!");
//         $this->assertEquals(bump("_nnnnnnn_n__n______nn__nn_nnn"), "Car Dead");
//         $this->assertEquals(bump("______n___n_"), "Woohoo!");
//         $this->assertEquals(bump("nnnnnnnnnnnnnnnnnnnnn"), "Car Dead");
//     }
// }

?>