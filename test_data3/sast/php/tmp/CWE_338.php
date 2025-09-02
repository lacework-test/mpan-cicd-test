<?php
  $pwd = rand();

  $token_0 = crypt($pwd);
  $token_1 = session_id($pwd);
?>
