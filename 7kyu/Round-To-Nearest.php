<?php


function roundUp(int $number, array $list): array
{
    if (count($list) == 0) {
        return [];
    }
    sort($list);

    foreach ($list as $key => $value) {
        $index = 0;
        if (abs($number) < abs($value)) {
            $index = $key;
            break;
        } else if (abs($number) == abs($value)) {
            return [$value];
        }
    }
    if ($list[$index] - $number == $number - $list[$index - 1]) {
        return [$list[$index - 1], $list[$index]];
    } elseif ($list[$index] - $number < $number - $list[$index - 1]) {
        return [$list[$index]];
    } else {
        return [$list[$index - 1]];
    }
}

// print_r(roundUp(10, [1, 2, 4, 8, 16, 32]));
// // answer => [8]
// print_r(roundUp(-120, [-10, -20, -40, -80, -160, -320]));
// answer => [-80, -160]
// print_r(roundUp(1, [-32, 12, 8, -10, -2, 4]));
// answer => [-2, 4]

function closestLowerHigherNr($array, $nr)
{
    sort($array);
    $re_arr = array('lower' => min(current($array), $nr), 'higher' => max(end($array), $nr), 'closest' => $nr);
    foreach ($array as $num) {
        if ($nr > $num) $re_arr['lower'] = $num;
        else if ($nr <= $num) {
            $re_arr['higher'] = $num;
            break;
        }
    }
    $re_arr['closest'] = (abs($nr - $re_arr['lower']) < abs($re_arr['higher'] - $nr)) ? $re_arr['lower'] : $re_arr['higher'];

    return $re_arr;
}
print_r(closestLowerHigherNr([1, 2, 4, 8, 16, 32], 10));
// // answer => [8]
print_r(closestLowerHigherNr([-10, -20, -40, -80, -160, -320], -120));
// answer => [-80, -160]
print_r(closestLowerHigherNr([-32, 12, 8, -10, -2, 4], 1));
// answer => [-2, 4]