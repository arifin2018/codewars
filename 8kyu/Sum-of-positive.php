<?php

function positive_sum($arr)
{
    $pos_array = 0;
    for ($i = 0; $i < count($arr); $i++) {
        if ($arr[$i] > 0) {
            // var_dump($i);
            $pos_array += $arr[$i];
        }
    }
    return $pos_array;
}
echo positive_sum([1, 2, 3, 4, 5]);
echo positive_sum([1, -2, 3, 4, 5]);


/* 

1 2 3 4 5 = 15 
12345(4)



*/
