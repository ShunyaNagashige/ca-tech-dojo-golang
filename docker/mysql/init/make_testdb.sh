#!/bin/sh

# [@]は，引数を伴った変数展開
echo "CREATE DATABASE IF NOT EXISTS game_db_test;" | "${mysql[@]}"

# game_db_testに対する権限を，　$MYSQL_USERに与える
echo "GRANT ALL ON game_db_test.* TO '"$MYSQL_USER"';" | "${mysql[@]}"