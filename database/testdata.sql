-- Create table for wallets
CREATE TABLE navnetjener.wallets ("id" bigserial,"created_at" timestamptz,"updated_at" timestamptz,"deleted_at" timestamptz,"first_name" varchar(255) NOT NULL,"last_name" varchar(255) NOT NULL,"orgnr" bigint NOT NULL,"pnr" text NOT NULL,"birth_date" text NOT NULL,"wallet_address" varchar(42) NOT NULL UNIQUE,PRIMARY KEY ("id"))

-- Create an index on the orgnr column
CREATE INDEX idx_orgnr ON navnetjener.wallets (orgnr);

-- Create an index on the pnr column
CREATE INDEX idx_pnr ON navnetjener.wallets (pnr);

-- Create an index on the wallet_address column
CREATE INDEX idx_wallet_address ON navnetjener.wallets (wallet_address);

/* 
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
INSERT INTO navnetjener.wallets (first_name, last_name, orgnr, pnr, birth_date, wallet_address) VALUES
('Elise', 'Berg', 310780472, '2105800000', '210580', '0x0c6598f08872e79c8195882e059acf8cee7eb468'),
('Lars', 'Myhre', 310780472, '1403900001', '140390', '0x2f6fa0b1adf996a1795dabc4b06e66b18ff867a5'),
('Nina', 'Pedersen', 310780472, '1509760002', '150976', '0x39d1786d6c23955830146b3658c6f028507c0fbe'),
('Johan', 'Aas', 310780472, '1106650003', '110665', '0x642c9271f215901ef01fde16f38b1257c97528a6'),
('Ingeborg', 'Knutsen', 310780472, '0809790004', '080979', '0x93ec94e79347c9a7b7b1c4685a24322bbbbfb26f'),
('Maria', 'Sunde', 310780472, '0102880005', '010288', '0x9865006528b2883bcc2a343a1b0c6168540d0847'),
('Knut', 'Hagen', 310780472, '0900000006', '090000', '0xaf4f666604609d92fb64b0ac2e3b5455026b6113'),
('Ellen', 'Eide', 310780472, '3003720007', '300372', '0xdc4bdc357910e51c22471a51be33fa840e850647'),
('Peter', 'Lang', 310780472, '2709890008', '270989', '0xe0457130d824848853d398f1e358c31c415eaba4'),
('Oscar', 'Nilsen', 310780472, '2502950009', '250295', '0xeae18965741f8d221d488512468f5f334198c1f1');

/*
  FORSTÅELSESFULL KOSTBAR TIGER AS
  310812277

  0x1a5c49f5a6398b84a68746d5d1172cb38f71104e
  0xd1db4aac0a1d1bc5e71c0d90da1aae4c4354a3e4
*/
INSERT INTO navnetjener.wallets (first_name, last_name, orgnr, pnr, birth_date, wallet_address) VALUES
('Elise', 'Berg', 310812277, '2105800000', '210580', '0x1a5c49f5a6398b84a68746d5d1172cb38f71104e'),
('Lars', 'Myhre', 310812277, '1403900001', '140390', '0xd1db4aac0a1d1bc5e71c0d90da1aae4c4354a3e4');


/*
  SPESIFIKK NORMAL TIGER AS
  310767859

  0x4f4441a36e5870018a9481fd7dab9d326f71f1fe
  0x635e62f8e16087875bea87dd26d7845104ccb1e1
  0xcf1f029280db9169c15841962f2282e57f04640f
*/
INSERT INTO navnetjener.wallets (first_name, last_name, orgnr, pnr, birth_date, wallet_address) VALUES
('Emma', 'Olsen', 310767859, '2204750012', '220475', '0x4f4441a36e5870018a9481fd7dab9d326f71f1fe'),
('Hans', 'Iversen', 310767859, '0807810013', '080781', '0x635e62f8e16087875bea87dd26d7845104ccb1e1'),
('Sara', 'Johansen', 310767859, '1210900014', '121090', '0xcf1f029280db9169c15841962f2282e57f04640f');
