CREATE TABLE User (
    sn INTEGER PRIMARY KEY AUTOINCREMENT,
    user_name TEXT UNIQUE NOT NULL,
    privilege INTEGER DEFAULT 1,
    user_password TEXT NOT NULL
);
-- 至少新增user_name跟user_password, privilege默認是1