-- ---
-- Globals
-- ---

-- SET SQL_MODE="NO_AUTO_VALUE_ON_ZERO";
-- SET FOREIGN_KEY_CHECKS=0;

-- ---
-- Table 'servers'
-- 
-- ---

DROP TABLE IF EXISTS "servers";
		
CREATE TABLE "servers" (
  "domain" VARCHAR(100) NULL DEFAULT NULL,
  "address" VARCHAR(100) NULL DEFAULT NULL,
  "ssl_grade" VARCHAR(2) NOT NULL,
  "country" VARCHAR(3) NULL DEFAULT NULL,
  "owner" VARCHAR(100) NULL DEFAULT NULL,
  INDEX KEY ("domain")
);

-- ---
-- Table 'domains'
-- 
-- ---

DROP TABLE IF EXISTS "domains";
		
CREATE TABLE "domains" (
  "domain" VARCHAR(100) NULL DEFAULT NULL,
  "servers_changed" BOOL NOT NULL DEFAULT false,
  "ssl_grade" VARCHAR(2) NOT NULL,
  "previous_ssl_grade" VARCHAR(2) NOT NULL,
  "logo" VARCHAR NULL DEFAULT NULL,
  "title" VARCHAR NULL DEFAULT NULL,
  "is_down" BOOL NOT NULL DEFAULT false
);

-- ---
-- Foreign Keys 
-- ---

ALTER TABLE "servers" ADD FOREIGN KEY (domain) REFERENCES "domains" ("domain");

-- ---
-- Table Properties
-- ---

-- ALTER TABLE "servers" ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE "domains" ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ---
-- Test Data
-- ---

-- INSERT INTO "servers" ("domain","address","ssl_grade","country","owner") VALUES
-- ('jeanbenitez.com','fakeaddress123','A+','CO','JEANB');
-- INSERT INTO "domains" ("domain","servers_changed","ssl_grade","previous_ssl_grade","logo","is_down") VALUES
-- ('jeanbenitez.com',false, 'B', 'A+', 'logo.png', false);