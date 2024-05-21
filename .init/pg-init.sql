CREATE TABLE IF NOT EXISTS account_balance (
    user_id SERIAL PRIMARY KEY,
    balance NUMERIC NOT NULL
);

INSERT INTO account_balance 
    (user_id, balance)
SELECT 1, 10000
WHERE 
    NOT EXISTS (
        SELECT user_id FROM account_balance WHERE user_id = 1
    );

INSERT INTO account_balance 
    (user_id, balance)
SELECT 2, 5000
WHERE 
    NOT EXISTS (
        SELECT user_id FROM account_balance WHERE user_id = 2
    );

INSERT INTO account_balance 
    (user_id, balance)
SELECT 3, 10
WHERE 
    NOT EXISTS (
        SELECT user_id FROM account_balance WHERE user_id = 3
    );

INSERT INTO account_balance 
    (user_id, balance)
SELECT 4, 0
WHERE 
    NOT EXISTS (
        SELECT user_id FROM account_balance WHERE user_id = 4
    );

INSERT INTO account_balance 
    (user_id, balance)
SELECT 5, 1000
WHERE 
    NOT EXISTS (
        SELECT user_id FROM account_balance WHERE user_id = 5
    );