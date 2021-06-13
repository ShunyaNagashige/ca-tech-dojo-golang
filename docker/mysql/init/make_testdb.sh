#!/bin/sh

# [@]は，引数を伴った変数展開
echo "CREATE DATABASE IF NOT EXISTS game_db_test;" | "${mysql[@]}"

# $MYSQL_USERに，game_db_testに対する権限を与える
echo "GRANT ALL ON game_db_test.* TO '"$MYSQL_USER"';" | "${mysql[@]}"