<?php
function divisors(int $n)
{
    // solve solution from myself
    $data = [];
    for ($i = 1; $i <= $n; $i++) {
        if ($n % $i == 0) {
            // array_push($data, $i);
            $data[] = $i;
        }
    }
    return count($data);

    // solve solution from codewars
    // $count = 0;
    // for ($i = 1; $i <= $n; $i++) {
    //     if ($n % $i == 0) {
    //         $count++;
    //     }
    // }
    // return $count;
}
// $data = 'ad';
$data = divisors(54);
print_r($data);


// class MyTestCases extends TestCase
// {
//     public function testBasicTests()
//     {
//         $this->assertEquals(1, divisors(1));
//         $this->assertEquals(4, divisors(10));
//         $this->assertEquals(2, divisors(11));
//         $this->assertEquals(8, divisors(54));
//         $this->assertEquals(7, divisors(64));
//     }
// }
