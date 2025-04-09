-- Create the symbol_data table with the specified definition
CREATE TABLE symbol_data (
    symbol VARCHAR(50) NOT NULL,
    open_time INT8 NOT NULL,
    "open" NUMERIC(30, 20) NOT NULL,
    high NUMERIC(30, 20) NOT NULL,
    low NUMERIC(30, 20) NOT NULL,
    "close" NUMERIC(30, 20) NOT NULL,
    volume NUMERIC(40, 20) NOT NULL,
    close_time INT8 NOT NULL
);

