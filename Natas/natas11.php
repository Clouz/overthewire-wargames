<?php

//input ed output
$defaultdata = array( "showpassword"=>"no", "bgcolor"=>"#ffffff");
$final = "ClVLIh4ASCsCBE8lAxMacFMZV2hdVVotEhhUJQNVAmhSEV4sFxFeaAw";

$json = json_encode($defaultdata);
$xorResult = base64_decode($final);

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

$key = xor_encrypt2($xorResult, $json );
echo "Ciave: ".$key."\n";

$defaultdata2 = array( "showpassword"=>"yes", "bgcolor"=>"#ffffff");
$test = base64_encode(xor_encrypt2( "qw8J", json_encode($defaultdata2)));
$tempdata = xor_encrypt2("qw8J", base64_decode($test));

//echo "Originale: ".$final."\n";
echo "Risultato: ".$test."\n";
echo "Decodeificato :".$tempdata."\n";

?>
