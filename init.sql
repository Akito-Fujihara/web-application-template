-- ユーザーが存在しない場合のみ作成
CREATE USER IF NOT EXISTS 'webapp_ro'@'%' IDENTIFIED BY 'password';

-- 読み取り専用権限を付与
GRANT SELECT ON *.* TO 'webapp_ro'@'%';

-- 権限を反映
FLUSH PRIVILEGES;
