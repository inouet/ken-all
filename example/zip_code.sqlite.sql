
create table zip_code (
  zip_code      text    not null,  -- 郵便番号7桁
  zip_type      integer not null,  -- 区分
  pref          text    not null,  -- 都道府県名
  city          text    not null,  -- 市区町村名
  town          text,              -- 町域名
  street        text,              -- 小字名、丁目、番地等
  name          text,              -- 大口事業所名
  update_code   integer            -- 修正コード
);

