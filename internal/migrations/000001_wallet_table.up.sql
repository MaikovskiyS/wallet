CREATE SCHEMA IF NOT EXISTS postgres;

CREATE TABLE IF NOT EXISTS wallets(
   id serial PRIMARY KEY,
   name VARCHAR (50) NOT NULL,
   balance NUMERIC NOT NULL
);
CREATE TABLE IF NOT EXISTS transactions(
   id serial PRIMARY KEY,
   transaction_type INT NOT NULL,
   wallet_id INT NOT NULL,
   amount NUMERIC NOT NULL,
   updated_balance NUMERIC NOT NULL
);

CREATE FUNCTION check_balance() RETURNS trigger AS $check_balance$
    BEGIN
        IF NEW.balance < 0 THEN
            RAISE EXCEPTION '% cannot have a negative balance', NEW.balance;
        END IF;
        RETURN NEW;
    END;
$check_balance$ LANGUAGE plpgsql;

CREATE TRIGGER check_balance BEFORE UPDATE ON wallets
    FOR EACH ROW EXECUTE FUNCTION check_balance();