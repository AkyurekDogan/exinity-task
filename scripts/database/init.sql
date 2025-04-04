CREATE TABLE public.symbol_data (
	time_stamp timestamp NOT NULL,	
	symbol varchar(50) NOT NULL,
	open_price numeric(20,10) NOT NULL,
	high_price numeric(20,10) NOT NULL,
	low_price numeric(20,10) NOT NULL,
	close_price numeric(20,10) NOT NULL,
	volume numeric(20,10) NOT NULL
);