-- +migrate Up
ALTER TABLE sync.monitored_txs
    ADD COLUMN network_id INTEGER NOT NULL DEFAULT 0,
    DROP CONSTRAINT IF EXISTS monitored_txs_pkey,
    ADD PRIMARY KEY (deposit_id, network_id);

-- +migrate Down
ALTER TABLE sync.monitored_txs
    DROP COLUMN IF EXISTS network_id,
    DROP CONSTRAINT IF EXISTS monitored_txs_pkey,
    ADD PRIMARY KEY (deposit_id);
