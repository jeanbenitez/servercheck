-- ---
-- Globals
-- ---

-- SET SQL_MODE="NO_AUTO_VALUE_ON_ZERO";
-- SET FOREIGN_KEY_CHECKS=0;

-- ---
-- Table 'server_items'
-- 
-- ---

DROP TABLE IF EXISTS "server_items";
		
CREATE TABLE "server_items" (
  "domain" VARCHAR(100) NULL DEFAULT NULL,
  "address" VARCHAR(100) NULL DEFAULT NULL,
  "ssl_grade" VARCHAR(2) NOT NULL,
  "country" VARCHAR(3) NULL DEFAULT NULL,
  "owner" VARCHAR(100) NULL DEFAULT NULL,
  PRIMARY KEY ("domain")
);

-- ---
-- Table 'servers'
-- 
-- ---

DROP TABLE IF EXISTS "servers";
		
CREATE TABLE "servers" (
  "domain" VARCHAR(100) NULL DEFAULT NULL,
  "servers_changed" BOOL NOT NULL DEFAULT false,
  "ssl_grade" VARCHAR(2) NOT NULL,
  "previous_ssl_grade" VARCHAR(2) NOT NULL,
  "logo" VARCHAR NULL DEFAULT NULL,
  "is_down" BOOL NOT NULL DEFAULT false,
  PRIMARY KEY ("domain")
);

-- ---
-- Foreign Keys 
-- ---

ALTER TABLE "server_items" ADD FOREIGN KEY (domain) REFERENCES "servers" ("domain");

-- ---
-- Table Properties
-- ---

-- ALTER TABLE "server_items" ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE "servers" ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ---
-- Test Data
-- ---

-- INSERT INTO "server_items" ("domain","address","ssl_grade","country","owner") VALUES
-- ('jeanbenitez.com','fakeaddress123','A+','CO','JEANB');
-- INSERT INTO "servers" ("domain","servers_changed","ssl_grade","previous_ssl_grade","logo","is_down") VALUES
-- ('jeanbenitez.com',false, 'B', 'A+', 'logo.png', false);