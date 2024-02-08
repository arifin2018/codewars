<?php

function wordsToMarks(string $str): int
{
    $huruf = "abcdefghijklmnopqrstuvwxyz";
    $hurufArray = [];
    for ($i = 0; $i < strlen($huruf); $i++) {
        array_push($hurufArray, $huruf[$i]);
    }

    $result = 0;

    for ($i = 0; $i < strlen($str); $i++) {
        $result += array_search($str[$i], $hurufArray) + 1;
    }


    return $result;
}
// l + o + v + e = 54
print_r(wordsToMarks("love"));
