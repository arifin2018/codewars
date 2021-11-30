<!-- <?php

        function is_palindrome($line)
        {
            $panjang_kata = strlen($line) - 1;
            $hasil = '';

            // for ($i = 0; $i < $panjang_kata; $i++) {
            //     $hasil = $line[$i];
            // }
            for ($i = $panjang_kata; $i >= 0; $i--) {
                $hasil .= $line[$i];
                // echo $i . '<br>';
                // $hasil += $line[$i];
            }

            if ($hasil == $line) {
                echo 'Palindrome ' . $line;
            } else {
                echo 'Not Palindrome ' . $line;
            }
            // var_dump($panjang_kata);
            // return;
        }

        echo is_palindrome('kontol') . "<br>";
        echo is_palindrome('12321');

// class IsPalindromeTest extends TestCase
// {
//     /**
//      * @dataProvider getValid
//      */
//     public function testItIs($word)
//     {
//         $this->assertTrue(is_palindrome($word));
//     }

//     /**
//      * @dataProvider getNotValid
//      */
//     public function testItIsNot($word)
//     {
//         $this->assertFalse(is_palindrome($word));
//     }

//     public function getNotValid()
//     {
//         return [
//             ["walter"],
//             [123456],
//             ["!anna"],
//             [".!!.*"],
//         ];
//     }

//     public function getValid()
//     {
//         return [
//             ["anna"],
//             [12321],
//             ["."],
//             [".!!."],
//         ];
//     }
// } -->
