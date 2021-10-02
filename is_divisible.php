<?php
function is_divisible($n, $x, $y)
{
    return ($n % $x == 0) && ($n % $y == 0) ? 'true' : 'false';
}
$data = is_divisible(8, 3, 4);
echo $data;


// class ExampleTest extends TestCase
// {
//     public function testSample()
//     {
//         $this->assertEquals(false, is_divisible(3, 3, 4));
//         $this->assertEquals(true, is_divisible(12, 3, 4));
//         $this->assertEquals(false, is_divisible(8, 3, 4));
//         $this->assertEquals(true, is_divisible(48, 3, 4));
//     }
// }
