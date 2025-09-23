CREATE DATABASE IF NOT EXISTS records;
USE records;


CREATE TABLE snakes (
    suid        CHAR(36) NOT NULL PRIMARY KEY DEFAULT (UUID()),
    owner_uuid  CHAR(36) NOT NULL,
    sid        VARCHAR(100) NOT NULL,
    species     VARCHAR(100) NOT NULL,
    sex         CHAR(1),
    age   CHAR(3),
    genes       CHAR(255),
    notes       CHAR(255)

);

CREATE TABLE feeding (
    uuid        CHAR(36) NOT NULL PRIMARY KEY,
    snake_uuid  CHAR(36) NOT NULL,
    sid        VARCHAR(100) NOT NULL,
    feed_date   DATE NOT NULL,
    prey_type   VARCHAR(100),
    prey_size    VARCHAR(50),
    notes       TEXT,
    FOREIGN KEY (snake_uuid) REFERENCES snakes(suid)
);

CREATE TABLE health (
    suid        CHAR(36) NOT NULL,
    owner_uuid  CHAR(36) NOT NULL,
    sid         VARCHAR(100) NOT NULL,
    check_date  DATE NOT NULL,
    weight      VARCHAR(36),
    length      VARCHAR(36),
    topic     VARCHAR(50),
    notes       TEXT,
    PRIMARY KEY (suid, check_date),
    FOREIGN KEY (suid) REFERENCES snakes(suid)
);

-- breeding

CREATE TABLE breeding (
    breeding_uuid        CHAR(36) NOT NULL PRIMARY KEY DEFAULT (UUID()),
    owner_uuid           CHAR(36) NOT NULL,
    suid    CHAR(36) NOT NULL,
    mate_suid      CHAR(36) NOT NULL,
    sid    CHAR(36) NOT NULL,
    mate_sid      CHAR(36) NOT NULL,
    breeding_year  DATE,
    weight  CHAR(36),
    cooling_start  DATE,
    cooling_end  DATE,
    warming_start  DATE,
    warming_end  DATE,
    pairing_date  DATE,
    gravid_date  DATE,
    lay_date DATE,
    clutch_size DATE,
    clutch_survive VARCHAR(36),
    outcome        VARCHAR(50),
    notes          TEXT,
     FOREIGN KEY (suid)      REFERENCES snakes(suid),
    FOREIGN KEY (mate_suid) REFERENCES snakes(suid)

);

-- offspring