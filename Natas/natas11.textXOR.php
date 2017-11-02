<?php
function xor_encrypt2($k, $te) {
    $key = $k;
    $text = $te;
    $outText = '';
    
    // Iterate through each character
    for($i=0;$i<strlen($text);$i++) {
        $outText .= $text[$i] ^ $key[$i % strlen($key)];
    }
    return $outText;
}

$testoChiaro = "Nella Vecchia Fattoria";
$chiave = "bazinga";
$cifrato = xor_encrypt2($chiave, $testoChiaro);
echo "testo chiaro: ".$testoChiaro."\n"."Chiave: ".$chiave."\n"."Cifrato: ".$cifrato."\n---\n";

$chiaveTrovata = xor_encrypt2($testoChiaro, $cifrato);
$decifroTesto = xor_encrypt2($chiaveTrovata, $cifrato);
echo "Chiave trovata: ".$chiaveTrovata."\n"."Testo decifrato: ".$decifroTesto."\n";
?>
