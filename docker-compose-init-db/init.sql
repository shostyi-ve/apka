CREATE TABLE IF NOT EXISTS cities
(
    id           BIGINT,
    latitude     DOUBLE PRECISION NOT NULL,
    longitude    DOUBLE PRECISION NOT NULL,
    city         TEXT             NOT NULL,
    country      TEXT             NOT NULL,
    region_code  TEXT             NOT NULL,
    prediction   JSONB,
    last_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS  history_weather
(
    id_city      INT,
    weather      JSONB,
    last_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id_city),
    FOREIGN KEY (id_city) REFERENCES cities (id)
);

CREATE TABLE IF NOT EXISTS  predictions
(
    id_city      INT,
    prediction   JSONB,
    last_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id_city),
    FOREIGN KEY (id_city) REFERENCES cities (id)
);
