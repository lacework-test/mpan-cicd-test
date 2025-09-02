$weakHashSalt = rand();
password_hash('123', PASSWORD_DEFAULT, ['salt'=> $weakHashSalt]);
