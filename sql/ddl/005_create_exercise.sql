CREATE TABLE TRACKER.EXERCISE (
    NAME VARCHAR(50) PRIMARY KEY NOT NULL,
    TYPE VARCHAR(100),
    VARIATION VARCHAR(100),
    CRET_TS TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    UPDT_TS TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);