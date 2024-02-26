/* CREATE TABLE "varlıklar" (
  "id" bigserial PRIMARY KEY,
  "hesap_iban" bigint  NOT NULL,
  "bakiye" bigint NOT NULL,
  "olusturma_tarihi" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "hesaplar" (
  "id" bigserial PRIMARY KEY,
  "iban" bigserial UNIQUE NOT NULL,
  "hesap_sahibi_ismi" varchar NOT NULL,
  "bakiye" bigint NOT NULL,
  "para_birimi" varchar NOT NULL,
  "olusturma_tarihi" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "para_transferi" (
  "id" bigserial PRIMARY KEY,
  "gonderen_iban" bigint NOT NULL,
  "alan_iban" bigint NOT NULL,
  "mikdar" bigint NOT NULL,
  "olusturma_tarihi" timestamptz NOT NULL DEFAULT 'now()'
);


CREATE INDEX ON "varlıklar" ("hesap_iban");

CREATE INDEX ON "hesaplar" ("hesap_sahibi_ismi");

CREATE INDEX ON "para_transferi" ("gonderen_iban");

CREATE INDEX ON "para_transferi" ("alan_iban");

CREATE INDEX ON "para_transferi" ("gonderen_iban", "alan_iban");

COMMENT ON COLUMN "varlıklar"."bakiye" IS 'negatif ve pozitif olabilir';

COMMENT ON COLUMN "para_transferi"."mikdar" IS 'pozitif olmali';

ALTER TABLE "varlıklar" ADD FOREIGN KEY ("hesap_iban") REFERENCES "hesaplar" ("iban");

ALTER TABLE "para_transferi" ADD FOREIGN KEY ("gonderen_iban") REFERENCES "hesaplar" ("iban");

ALTER TABLE "para_transferi" ADD FOREIGN KEY ("alan_iban") REFERENCES "hesaplar" ("iban");




 */


CREATE TABLE "kullanıcılar" (
  "kullanıcı_adı" varchar PRIMARY KEY,
  "şifre" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL);

CREATE TABLE "varlıklar" (
  "id" bigserial PRIMARY KEY,
  "hesap_iban" bigint  NOT NULL,
  "bakiye" bigint NOT NULL,
  "olusturma_tarihi" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "hesaplar" (
  "id" bigserial PRIMARY KEY,
  "iban"  bigserial  UNIQUE ,
  "hesap_sahibi_ismi" varchar NOT NULL,
  "bakiye" bigint NOT NULL,
  "para_birimi" varchar NOT NULL,
  "olusturma_tarihi" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "para_transferi" (
	"id" bigserial PRIMARY KEY,
	"gonderen_iban" bigint NOT NULL,
	"alan_iban" bigint NOT NULL,
	"bakiye" bigint NOT NULL,
	"olusturma_tarihi" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE INDEX ON "varlıklar" ("hesap_iban");

CREATE INDEX ON "hesaplar" ("hesap_sahibi_ismi");

CREATE INDEX ON "para_transferi" ("gonderen_iban");

CREATE INDEX ON "para_transferi" ("alan_iban");

CREATE INDEX ON "para_transferi" ("gonderen_iban", "alan_iban");

COMMENT ON COLUMN "varlıklar"."bakiye" IS 'negatif ve pozitif olabilir';

COMMENT ON COLUMN "para_transferi"."bakiye" IS 'pozitif olmali';

ALTER TABLE "varlıklar" ADD FOREIGN KEY ("hesap_iban") REFERENCES "hesaplar" ("iban");

ALTER TABLE "para_transferi" ADD FOREIGN KEY ("gonderen_iban") REFERENCES "hesaplar" ("iban");

ALTER TABLE "para_transferi" ADD FOREIGN KEY ("alan_iban") REFERENCES "hesaplar" ("iban");

CREATE SEQUENCE dizi;

ALTER TABLE "hesaplar" ALTER COLUMN "iban" SET DEFAULT nextval('dizi');

ALTER TABLE para_transferi RENAME COLUMN bakiye TO mikdar;

UPDATE hesaplar set bakiye=1000 where iban=1 ;




ALTER TABLE "hesaplar" ADD FOREIGN KEY ("iban") REFERENCES "kullanıcılar" ("kullanıcı_adı");

CREATE UNIQUE INDEX ON "hesaplar" ("iban", "para_birimi");






