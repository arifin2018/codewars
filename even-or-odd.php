<?php
function even_or_odd(int $n)
{
    return $n & 1 ? 'Odd' : 'Even';
}
$data =  even_or_odd(1);
echo $data;
// class EvenOrOddTest extends TestCase
// {
//     public function testExamples()
//     {
//         $this->assertEquals("Even", even_or_odd(2));
//         $this->assertEquals("Even", even_or_odd(0));
//         $this->assertEquals("Odd", even_or_odd(7));
//         $this->assertEquals("Odd", even_or_odd(1));
//     }
// }
