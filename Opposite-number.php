<?php


function opposite($n)
{
    if ($n < 0) {
        return $n * -1;
    } else {
        return $n * -1;
    }
    // return $n *= -1;
}

// echo opposite(112) . "<br>";
echo opposite(-1) . "<br>";
echo opposite(125) . "<br>";

// class OppositeTest extends TestCase
// {
//     public function testThatOppositeWorksForExamplesProvidedInDescription()
//     {
//         $this->assertEquals(opposite(1), -1);
//         $this->assertEquals(opposite(14), -14);
//         $this->assertEquals(opposite(-34), 34);
//     }
//     public function testThatOppositeWorksForMyCustomTests()
//     {
//         // TODO: Implement TDD cycle: write failing test, write code to pass it, repeat
//     }
// }
