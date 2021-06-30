<?php
function square_sum($numbers)
{
    $total = 0;
    for ($i = 0; $i < count($numbers); $i++) {
        // echo $i . '<br>';
        // print_r($numbers[$i]);
        $total += $numbers[$i] * $numbers[$i];
    }
    return $total;
}

// print_r(square_sum([1, 2]));
print_r(square_sum([0, 3, 4, 5]));
// // echo square_sum(2, 2);

// function myfunction($v)
// {
//     return ($v * $v);
// }

// $a = array(1, 2, 3, 4, 5);
// print_r(array_map("myfunction", $a));


// class SquareSumTestCases extends TestCase
// {
// public function testBasic()
// {
// $this->assertEquals(square_sum([1, 2]), 5);
// $this->assertEquals(square_sum([0, 3, 4, 5]), 50);
// $this->assertEquals(square_sum([]), 0);
// $this->assertEquals(square_sum([-1, -2]), 5);
// $this->assertEquals(square_sum([-1, 0, 1]), 2);
// }
// }


// Complete the square sum function so that it squares each
// number passed into it and then sums the results together.

// For example, for [1, 2, 2] it should return 9 because 1^2 + 2^2 + 2^2 = 9.