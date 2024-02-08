<?php

$matrixArray = [
    [1, 2, 3],
    [4, 5, 6],
    [9, 8, 9]
];

function tes($data)
{
    $data1 = 0;
    $data2 = 0;
    $index = 0;
    for ($i = 0; $i < count($data); $i++) {
        $data1 += $data[$index][$i];
        $index++;
        $data2 += $data[$index][$data--];
    }
    echo $data2;
}


tes($matrixArray);
