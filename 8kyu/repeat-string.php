<?php

function repeatStr($n, $str)
{
    // $hasil = '';
    // for ($i = 0; $i < $n; $i++) {
    //     $hasil .= $str;
    //     // echo $i;
    // }
    // return $hasil;

    echo str_repeat($str, $n);
}

// echo repeatStr(3, "*");
echo repeatStr(2, "@");

// class RepeatStrTest extends TestCase
// {
//     public function testThatSomethingShouldHappen()
//     {
//       $this->assertEquals("***", repeatStr(3, "*"));
//       $this->assertEquals("@@", repeatStr(2, "@"));
//       $this->assertEquals("!", repeatStr(1, "!"));
//     }
// }