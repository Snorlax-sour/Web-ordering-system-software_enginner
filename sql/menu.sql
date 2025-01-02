CREATE TABLE Menu (
    menu_sn INTEGER PRIMARY KEY AUTOINCREMENT,
    menu_name TEXT UNIQUE NOT NULL,
    menu_describe_and_material TEXT,
    menu_need_ingredient INTEGER,
    menu_money INTEGER NOT NULL,
    menu_visable INTEGER DEFAULT 1 not null,  -- 1: 顯示, 0: 不顯示
    menu_image BLOB not null,
    FOREIGN KEY (menu_need_ingredient) REFERENCES Menu_need_ingredient(join_sn)
);