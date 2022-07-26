<?php

define('__ROOT__', dirname(dirname(__FILE__)));
// require_once __ROOT__ . '/8kyu/Convert-a-Number-to-a-String.php';
require_once '/Users/arifinlenna/Documents/belajar/codewars/8kyu/Convert-a-Number-to-a-String.php';

use PHPUnit\Framework\TestCase;

class TestingConvertANumberToAString extends TestCase
{
    public function testingNumberToStrings()
    {
        $this->assertEquals("23", numberToString(23));
        $this->assertSame('22', numberToString(22));
        $this->assertIsString(numberToString(22));
    }
}
