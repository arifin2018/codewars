<?php
function past($h, $m, $s)
{
    return (3600 * $h + 60 * $m + $s) * 1000;
}


echo past(0, 1, 1) . '<br>';
echo past(1, 1, 1) . '<br>';
echo past(0, 0, 0) . '<br>';
echo past(1, 0, 1) . '<br>';
echo past(1, 0, 0) . '<br>';
