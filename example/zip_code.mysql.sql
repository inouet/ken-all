
create table zip_code (
  zip_code    varchar(7)   not null,  -- 郵便番号7桁
  zip_type    int          not null,  -- 区分(1: 住所 2: 事業所)
  pref        varchar(20)  not null,  -- 都道府県名
  city        varchar(255) not null,  -- 市区町村名
  town        varchar(255),           -- 町域名
  street      varchar(255),           -- 小字名、丁目、番地等
  name        varchar(255),           -- 大口事業所名
  update_code int                     -- 修正コード
);

-- 修正コード
-- 住所
--  10: 変更なし 
--  11: 市政・区政・町政・分区・政令指定都市施行
--  12: 住居表示の実施
--  13: 区画整理
--  14: 郵便区調整等
--  15: 訂正
--  16: 廃止（廃止データのみ使用）
-- 事業所
--  20: 修正なし 
--  21: 新規追加 
--  25: 廃止

