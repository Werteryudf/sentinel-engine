CREATE TABLE users(
    user_id  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(50) UNIQUE NOT NULL,
    user_password VARCHAR NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    is_active BOOLEAN DEFAULT TRUE
);

CREATE TABLE assets(
    asset_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    ticker VARCHAR UNIQUE NOT NULL,
    name_asset VARCHAR NOT NULL,
    asset_type VARCHAR NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE prices(
    price_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    asset_id UUID REFERENCES assets(asset_id) ON DELETE CASCADE,
    price NUMERIC(18,8),
    volume NUMERIC(18,2),
    recorded_at TIMESTAMPTZ DEFAULT NOW()
);






