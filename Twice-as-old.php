<?php
function twice_as_old($dad_years_old, $son_years_old)
{
    $data = $dad_years_old - ($son_years_old * 2);
    return trim($data, '-');
}
$data = twice_as_old(55, 30);
// echo abs($data);
echo $data;
// class MyTestCases extends TestCase
// {
//     public function testBasicTests()
//     {
//         $this->assertEquals(twice_as_old(36, 7), 22);
//         $this->assertEquals(twice_as_old(55, 30), 5);
//         $this->assertEquals(twice_as_old(42, 21), 0);
//         $this->assertEquals(twice_as_old(22, 1), 20);
//         $this->assertEquals(twice_as_old(29, 0), 29);
//     }
// }
