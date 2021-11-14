<?php

class Server
{
    public function serve(\Closure $f)
    {
        if (file_exists('test.sock')) {
            unlink('test.sock');
        }
        $socket = stream_socket_server("unix://test.sock", $errno, $errstr);
        if (!$socket) {
            echo "$errstr ($errno)<br />\n";
        } else {
            while ($conn = stream_socket_accept($socket, -1)) {
                $line = fgets($conn);
                $data = json_decode($line, true);
                $res = $f($data);
                fwrite($conn, json_encode($res) . "\n");
                fclose($conn);
            }
            fclose($socket);
        }
    }
}

(new Server())->serve(function($data){
    $data["hey"] = "you";
    return $data;
});
