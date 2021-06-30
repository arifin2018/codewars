<?php

function getCount($str)
{
    $total = 0;
    $vowels = array('a', 'e', 'i', 'o', 'u');

    for ($i = 0; $i < strlen($str); $i++) {
        for ($j = 0; $j < 5; $j++)
            if ($str[$i] == $vowels[$j]) {
                $total++;
                break;
            }
    }
    return $total;
}

echo getCount("abracadabra");
// Vokal = a, e, i, o, u
// class VovelCountCase extends TestCase
// {
//     public function testBasics()
//     {
//         $this->assertEquals(5, getCount("abracadabra"));
//     }
// }
