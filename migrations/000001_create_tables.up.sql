CREATE TABLE IF NOT EXISTS infinitybottles (
    id BIGSERIAL PRIMARY KEY,
    bottle_name VARCHAR(255) NOT NULL,
    number_of_contributions INT NOT NULL DEFAULT 0,
    empty_start BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP(0) WITH time zone NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP(0) WITH time zone NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS contributions (
    id BIGSERIAL PRIMARY KEY,
    infinitybottle_id BIGINT NOT NULL REFERENCES infinitybottles(id),
    added_at TIMESTAMP(0) WITH time zone NOT NULL DEFAULT NOW(),
    brand_name VARCHAR(255) NOT NULL,
    tags text[] NULL
)
