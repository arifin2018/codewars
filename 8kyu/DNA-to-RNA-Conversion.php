<?php 

function dnaToRna($str) {
    return str_replace('U', 'T', $str);
}

echo dnaToRna("GAUUCCACCGACUUCCCAAGUACCGGAAGCGCGACCAACUCGCACAGC");

//   class MyTestCases extends TestCase {
//     public function testFixedTests() {
//       $this->assertEquals('UUUU', dnaToRna("TTTT"));
//       $this->assertEquals('GCAU', dnaToRna("GCAT"));
//       $this->assertEquals("", dnaToRna(""));
//       $this->assertEquals('U', dnaToRna("T"));
//       $this->assertEquals('GACCGCCGCC', dnaToRna("GACCGCCGCC"));
//       $this->assertEquals('GAUUCCACCGACUUCCCAAGUACCGGAAGCGCGACCAACUCGCACAGC', dnaToRna("GATTCCACCGACTTCCCAAGTACCGGAAGCGCGACCAACTCGCACAGC"));
//       $this->assertEquals('CACGACAUACGGAGCAGCGCACGGUUAGUACAGCUGUCGGUGAACUCCAUGACA', dnaToRna("CACGACATACGGAGCAGCGCACGGTTAGTACAGCTGTCGGTGAACTCCATGACA"));
//       $this->assertEquals('CACGACAUACGGAGCAGCGCACGGUUAGUACAGCUGUCGGUGAACUCCAUGACA', dnaToRna("CACGACATACGGAGCAGCGCACGGTTAGTACAGCTGTCGGTGAACTCCATGACA"));
//       $this->assertEquals('AACCCUGUCCACCAGUAACGUAGGCCGACGGGAAAAAUAAACGAUCUGUCAAUG', dnaToRna("AACCCTGTCCACCAGTAACGTAGGCCGACGGGAAAAATAAACGATCTGTCAATG"));
//     }
//   }

?>