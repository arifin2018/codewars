<?php

define('__ROOT__', dirname(dirname(__FILE__)));
// require_once __ROOT__ . '/8kyu/Convert-a-Number-to-a-String.php';
require_once '/Users/arifinlenna/Documents/belajar/codewars/7kyu/Love-vs-friendship.php';

use PHPUnit\Framework\TestCase;

// ./vendor/bin/phpunit --bootstrap vendor/autoload.php testing/7kyu/Lovevsfriendship.php


class Lovevsfriendship extends TestCase
{
    /** @test */
    public function Lovefriendship()
    {
        $this->assertEquals(100, wordsToMarks("attitude"));
        $this->assertEquals(75, wordsToMarks("friends"));
        $this->assertEquals(66, wordsToMarks("family"));
        $this->assertEquals(99, wordsToMarks("selfness"));
        $this->assertEquals(96, wordsToMarks("knowledge"));
    }
}
