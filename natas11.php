<?


$defaultdata = array( "showpassword"=>"no", "bgcolor"=>"#ffffff");
$final = "ClVLIh4ASCsCBE8lAxMacFMZV2hdVVotEhhUJQNVAmhSEV4sFxFeaAw%3D";
$encoded = base64_decode($final);
$decoded = json_encode($defaultdata);

echo $encoded."\n";
echo $decoded."\n";

function xor_encrypt2($en, $de) {
    $key = $en;
    $text = $de;
    $outText = '';

    // Iterate through each character
    for($i=0;$i<strlen($text);$i++) {
        $outText .= $text[$i] ^ $key[$i % strlen($key)];
    }

    return $outText;
}

echo xor_encrypt2($encoded, $decoded)."\n\n\n";

function xor_encrypt($in) {
    $key = 'qw8Jqw8Jqw8Jqw8Jqw8Jqw8Jqw8Jqw8Jqw8Jqw8Jq';
    $text = $in;
    $outText = '';

    // Iterate through each character
    for($i=0;$i<strlen($text);$i++) {
        $outText .= $text[$i] ^ $key[$i % strlen($key)];
    }

    return $outText;
}

$test = base64_encode(xor_encrypt(json_encode($defaultdata)));

echo $final."\n";
echo $test."\n";

if ($test == $final) {
    echo "Chiave trovata!\n";
}
else {
    echo "Ritenta!\n";
}

?>
