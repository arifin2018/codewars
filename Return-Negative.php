<?php


function makeNegative(float $num)
{
    if ($num > 0) {
        return -$num;
    } else {
        return $num;
    }
    // return -abs($num);
    // return $num > 0 ? -$num : $num;
}

// echo makeNegative(7);
// echo makeNegative(-20);
echo makeNegative(0);
// echo makeNegative(0.12);
// class MakeNegativeTestCases extends TestCase
// {
//     // test makeNegative()
//     public function testMakeNegative()
//     {
//         $this->assertEquals(makeNegative(7), -7);
//         $this->assertEquals(makeNegative(-20), -20);
//         $this->assertEquals(makeNegative(0), 0);
//         $this->assertEquals(makeNegative(0.12), -0.12);
//     }
// }
