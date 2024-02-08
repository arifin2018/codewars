<?php

function index($target, array $params)
{
    $idx_start = 0;
    $idx_end = count($params) - 1;
    $idx_mid = floor((count($params) - 1) / 2) + $idx_start;
    $step = 0;

    while ($idx_start <= $idx_mid) {
        if ($params[$idx_mid] == $target) {
            return $step;
        } elseif ($params[$idx_mid] < $target) {
            $idx_start = $idx_mid;
            $idx_end = $idx_end;
            print_r($idx_end - $idx_start);
            die;
            $idx_mid = $idx_end;
        }
    }
}
index(55, [1, 4, 11, 25, 32, 37, 40, 43, 47, 49, 53, 55]);
