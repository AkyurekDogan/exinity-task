CREATE TABLE symbol_data (
	symbol varchar(50) NOT NULL,
	open_time numeric(20, 10) NOT NULL,
	open numeric(20, 10) NOT NULL,
	high numeric(20, 10) NOT NULL,
	low numeric(20, 10) NOT NULL,
	close numeric(20, 10) NOT NULL,
	volume numeric(20, 10) NOT null,
	close_time numeric(20, 10) NOT NULL
);