CREATE TABLE IF NOT EXISTS users
(
    id              VARCHAR(36) NOT NULL,
    first_name      VARCHAR(64) NOT NULL,
    middle_name     VARCHAR(64) DEFAULT NULL,
    last_name       VARCHAR(64) NOT NULL,
    type            VARCHAR(10) NOT NULL, -- 'customer', 'driver'
    status          VARCHAR(10) NOT NULL, -- 'active', 'inactive'
    created_at      TIMESTAMP   DEFAULT (now() at time zone 'utc'),
    updated_at      TIMESTAMP   DEFAULT NULL,

    PRIMARY KEY (id)
);