<?php

function remove_char(string $s)
{
    return substr($s, 1, -1);
}
echo  remove_char('eloquent') . "<br>";
echo  remove_char('country') . "<br>";
echo  remove_char('person') . "<br>";
echo  remove_char('place') . "<br>";

    // class RemoveCharTest extends TestCase {
    //     public function testFixed() {
    //       $this->assertEquals('loquen', remove_char('eloquent'));
    //       $this->assertEquals('ountr', remove_char('country'));
    //       $this->assertEquals('erso', remove_char('person'));
    //       $this->assertEquals('lac', remove_char('place'));
    //     }
    //   }