<?php

function solution($str)
{
    $string = '';
    for ($i = strlen($str) - 1; $i >= 0; $i--) {
        $string .= $str[$i];
    }
    return $string;
}

echo solution("world") . '<br>';
echo solution("hello") . '<br>';
echo solution("") . '<br>';
echo solution("h") . '<br>';

// class ReversedStringsTest extends TestCase {
//     public function testExamples() {
//       $this->assertEquals("dlrow", solution("world"));
//       $this->assertEquals("olleh", solution("hello"));
//       $this->assertEquals("", solution(""));
//       $this->assertEquals('h', solution("h"));
//     }
//   }