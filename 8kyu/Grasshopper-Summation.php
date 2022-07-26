<?php

// function summation($n)
// {
//     $number = 0;
//     for ($i = 0; $i <= $n; $i++) {
//         echo $number += $i;
//     }
//     return $number;
// }

function summation($n)
{
    // return range(1, $n); 
    // print_r(range(1, $n));
    $arrayNumber = range(1, $n);
    return array_sum($arrayNumber);
}

// echo summation(1);
// echo summation(2);
// echo summation(3);
// echo summation(4);
echo summation(5);

// class SummationTest extends TestCase
// {
//     public function testThatSummationWorksForExampleTests()
//     {
//         $this->assertEquals(summation(1), 1);
//         $this->assertEquals(summation(2), 3);
//         $this->assertEquals(summation(3), 6);
//         $this->assertEquals(summation(4), 10);
//         $this->assertEquals(summation(5), 15);
//     }
// }
