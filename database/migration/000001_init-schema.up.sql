CREATE TABLE "akun" (
  "id" bigserial PRIMARY KEY,
  "nama" varchar NOT NULL,
  "saldo" bigint NOT NULL,
  "dibuat_pada" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "catatan" (
  "id" bigserial PRIMARY KEY,
  "id_akun" bigint NOT NULL,
  "jumlah" bigint NOT NULL,
  "dibuat_pada" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfer" (
  "id" bigserial PRIMARY KEY,
  "pengirim" bigint NOT NULL,
  "penerima" bigint NOT NULL,
  "jumlah" bigint NOT NULL,
  "dibuat_pada" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "akun" ("nama");

CREATE INDEX ON "catatan" ("id_akun");

CREATE INDEX ON "transfer" ("pengirim");

CREATE INDEX ON "transfer" ("penerima");

CREATE INDEX ON "transfer" ("pengirim", "penerima");

COMMENT ON COLUMN "catatan"."jumlah" IS 'bisa positif bisa negatif';

COMMENT ON COLUMN "transfer"."jumlah" IS 'harus positif';

ALTER TABLE "catatan" ADD FOREIGN KEY ("id_akun") REFERENCES "akun" ("id");

ALTER TABLE "transfer" ADD FOREIGN KEY ("pengirim") REFERENCES "akun" ("id");

ALTER TABLE "transfer" ADD FOREIGN KEY ("penerima") REFERENCES "akun" ("id");