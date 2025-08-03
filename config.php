<?php
$host = getenv("DB_HOST");
$db   = getenv("DB_NAME");
$user = getenv("DB_USER");
$pass = getenv("DB_PASS");

$pdo = new PDO("mysql:host=$host;dbname=$db", $user, $pass);
?>
