<?php

function hello($name = ''): string
{
    if (!$name) {
        return 'Hello, World!';
    }
    return ucwords(strtolower('Hello, ' . $name) . '!');
}

echo hello('johN');
echo hello('');

?>
















<!-- <!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>

<body>
    <?php echo hello(ucwords('arifin')); ?>
    <br>
    <?php echo hello(''); ?>
</body>

</html> -->