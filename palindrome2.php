<?php

function is_palindrome($line)
{
    $temp = '';

    for ($i = strlen($line) - 1; $i >= 0; $i--) {
        $temp = $temp . substr($line, $i);
    }

    return true;
}

echo is_palindrome('anna');
echo is_palindrome('123456');
echo is_palindrome('!anna');
echo is_palindrome('.!!.*');


// class IsPalindromeTest extends TestCase
// {
//   /**
//    * @dataProvider getValid
//    */
//   public function testItIs($word)
//   {
//     $this->assertTrue(is_palindrome($word));
//   }
  
//   /**
//    * @dataProvider getNotValid
//    */
//   public function testItIsNot($word)
//   {
//     $this->assertFalse(is_palindrome($word));
//   }
  
//   public function getNotValid()
//   {
//     return [
//       ["walter"],
//       [123456],
//       ["!anna"],
//       [".!!.*"],
//     ];
//   }
    
//   public function getValid()
//   {
//     return [
//       ["anna"],
//       [12321],
//       ["."],
//       [".!!."],
//     ];
//   }
// }