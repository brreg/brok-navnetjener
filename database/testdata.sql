/*
  Denne filen gjør 3 forskjellige ting:
  1. Oppretter schema og tabell for wallets
  2. Legger inn testdata for 3 forskjellige foretak som ligger på The Graph
  3. Legger inn testdata for 1 foretak som ligger på en lokal node
      Den lokalen noden startes via brøk repoet https://github.com/brreg/brok/
*/

-- Create database
CREATE DATABASE brok;

-- Connect to database
\c brok;

-- Create schema
CREATE SCHEMA navnetjener;

-- Create table for wallets
CREATE TABLE navnetjener.wallets (
  "id" bigserial,
  "created_at" timestamptz,
  "updated_at" timestamptz,
  "deleted_at" timestamptz,
  "owner_person_first_name" varchar(255),
  "owner_person_last_name" varchar(255),
  "owner_person_birth_date" varchar(6),
  "owner_person_fnr" varchar(11),
  "owner_company_name" varchar(255),
  "owner_company_orgnr" varchar(9),
  "cap_table_orgnr" varchar(9) NOT NULL,
  "wallet_address" varchar(42) NOT NULL UNIQUE,
  PRIMARY KEY ("id")
);

-- Create an index on the cap_table_orgnr column
CREATE INDEX idx_cap_table_orgnr ON navnetjener.wallets (cap_table_orgnr);

-- Create an index on the owner_person_fnr column
CREATE INDEX idx_owner_person_fnr ON navnetjener.wallets (owner_person_fnr);

-- Create an index on the owner_company_orgnr column
CREATE INDEX idx_owner_company_orgnr ON navnetjener.wallets (owner_company_orgnr);

-- Create an index on the wallet_address column
CREATE INDEX idx_wallet_address ON navnetjener.wallets (wallet_address);


/*
----------------------------------------------------------------------------------------------------
  bruker en eldre versjon av The Graph med mye testdata
  https://api.thegraph.com/subgraphs/name/broklab/captable_dev_11/ 
*/

/* 
  Aksjeeiere for foretaket LIVLIG SKYFRI TIGER AS
  orgnr 310780472

  WalletAddresser:
  0x0c6598f08872e79c8195882e059acf8cee7eb468
  0x2f6fa0b1adf996a1795dabc4b06e66b18ff867a5
  0x39d1786d6c23955830146b3658c6f028507c0fbe
  0x642c9271f215901ef01fde16f38b1257c97528a6
  0x93ec94e79347c9a7b7b1c4685a24322bbbbfb26f
  0x9865006528b2883bcc2a343a1b0c6168540d0847
  0xaf4f666604609d92fb64b0ac2e3b5455026b6113
  0xdc4bdc357910e51c22471a51be33fa840e850647
  0xe0457130d824848853d398f1e358c31c415eaba4
  0xeae18965741f8d221d488512468f5f334198c1f1
*/
INSERT INTO navnetjener.wallets (owner_person_first_name, owner_person_last_name, cap_table_orgnr, owner_person_fnr, owner_person_birth_date, wallet_address) VALUES
('Elise', 'Berg', '310780472', '21058000000', '210580', '0x0c6598f08872e79c8195882e059acf8cee7eb468'),
('Lars', 'Myhre', '310780472', '14039000010', '140390', '0x2f6fa0b1adf996a1795dabc4b06e66b18ff867a5'),
('Nina', 'Pedersen', '310780472', '15097600020', '150976', '0x39d1786d6c23955830146b3658c6f028507c0fbe'),
('Johan', 'Aas', '310780472', '11066500030', '110665', '0x642c9271f215901ef01fde16f38b1257c97528a6'),
('Ingeborg', 'Knutsen', '310780472', '08097900040', '080979', '0x93ec94e79347c9a7b7b1c4685a24322bbbbfb26f'),
('Maria', 'Sunde', '310780472', '01028800050', '010288', '0x9865006528b2883bcc2a343a1b0c6168540d0847'),
('Knut', 'Hagen', '310780472', '09000000060', '090000', '0xaf4f666604609d92fb64b0ac2e3b5455026b6113'),
('Ellen', 'Eide', '310780472', '30037200070', '300372', '0xdc4bdc357910e51c22471a51be33fa840e850647'),
('Peter', 'Lang', '310780472', '27098900080', '270989', '0xe0457130d824848853d398f1e358c31c415eaba4'),
('Oscar', 'Nilsen', '310780472', '25029500090', '250295', '0xeae18965741f8d221d488512468f5f334198c1f1');

/*
  FORSTÅELSESFULL KOSTBAR TIGER AS
  310812277

  0x1a5c49f5a6398b84a68746d5d1172cb38f71104e
  0xd1db4aac0a1d1bc5e71c0d90da1aae4c4354a3e4
*/
INSERT INTO navnetjener.wallets (owner_person_first_name, owner_person_last_name, cap_table_orgnr, owner_person_fnr, owner_person_birth_date, wallet_address) VALUES
('Elise', 'Berg', '310812277', '21058000000', '210580', '0x1a5c49f5a6398b84a68746d5d1172cb38f71104e'),
('Lars', 'Myhre', '310812277', '14039000010', '140390', '0xd1db4aac0a1d1bc5e71c0d90da1aae4c4354a3e4');


/*
  SPESIFIKK NORMAL TIGER AS
  310767859

  0x4f4441a36e5870018a9481fd7dab9d326f71f1fe
  0x635e62f8e16087875bea87dd26d7845104ccb1e1
  0xcf1f029280db9169c15841962f2282e57f04640f
*/
INSERT INTO navnetjener.wallets (owner_person_first_name, owner_person_last_name, cap_table_orgnr, owner_person_fnr, owner_person_birth_date, wallet_address) VALUES
('Emma', 'Olsen', '310767859', '22047500120', '220475', '0x4f4441a36e5870018a9481fd7dab9d326f71f1fe'),
('Hans', 'Iversen', '310767859', '08078100130', '080781', '0x635e62f8e16087875bea87dd26d7845104ccb1e1'),
('Sara', 'Johansen', '310767859', '12109000140', '121090', '0xcf1f029280db9169c15841962f2282e57f04640f');


/*
----------------------------------------------------------------------------------------------------
  Data til den lokal noden
*/

/*
  For aksjonerer som er personer

  Ryddig Bobil AS
  815493000

  0xbbb12c73703a8dc9ae2569e1c7ad699a5ac8c782
  0xcc6aa2c0d12716916e19012e954a0630fa25e097
  0x8be848ce9ebba1e304e6daa1d6b1b40f17e478fd
*/
INSERT INTO navnetjener.wallets (owner_person_first_name, owner_person_last_name, cap_table_orgnr, owner_person_fnr, owner_person_birth_date, wallet_address) VALUES
('Elise', 'Berg', '815493000', '21058000000', '210580', '0xbbb12c73703a8dc9ae2569e1c7ad699a5ac8c782'),
('Lars', 'Myhre', '815493000', '14039000001', '140390', '0xcc6aa2c0d12716916e19012e954a0630fa25e097'),
('Nina', 'Pedersen', '815493000', '15097600002', '150976', '0x8be848ce9ebba1e304e6daa1d6b1b40f17e478fd');

/*
  For aksjonerer som er foretak

  Ryddig Bobil AS
  815493000

  0xf04eb77c73c11d4b9ec610cf8ce6b51b7f78929b
  0xee879e18569a12687489a8cc48b53292ea2907c6
  0xc15451645ba50375580f673647c3ac34aad22e62
*/
INSERT INTO navnetjener.wallets (owner_company_name, owner_company_orgnr, cap_table_orgnr, wallet_address) VALUES
('Verdig Brygge AS', '310780472', '815493000', '0xf04eb77c73c11d4b9ec610cf8ce6b51b7f78929b'),
('Stilig Fisk AS', '310812277', '815493000', '0xee879e18569a12687489a8cc48b53292ea2907c6'),
('Gnien Dugnad AS', '310767859', '815493000', '0xc15451645ba50375580f673647c3ac34aad22e62');