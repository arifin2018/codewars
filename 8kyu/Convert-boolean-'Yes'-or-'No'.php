<?php

function boolToWord($bool)
{
    // if ($bool = true) {
    //     return 'yes';
    // } else {
    //     return "no";
    // }
    // return $bool;

    return $bool === true ? 'yes' : 'no';
}

echo boolToword(false);
// echo boolToword(false);


// PHPUnit Test Examples:
// TODO: Replace examples and use TDD development by writing your own tests
// class MyTestCases extends TestCase
// {
//     // test function names should start with "test"
//     public function testThatSomethingShouldHappen()
//     {
//         $this->assertEquals(boolToword(true), "Yes");
//         $this->assertEquals(boolToword(false), "No");
//     }
// }
