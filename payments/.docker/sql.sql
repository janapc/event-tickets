CREATE TABLE IF NOT EXISTS payments (
                                        id UUID PRIMARY KEY,
    user_email TEXT NOT NULL,
    status TEXT NOT NULL,
    event_id TEXT NOT NULL,
    amount NUMERIC(12, 2) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);

CREATE TABLE IF NOT EXISTS  transactions (
                                             id UUID PRIMARY KEY,
                                             status TEXT NOT NULL,
                                             payment_id UUID NOT NULL UNIQUE,
                                             reason TEXT,

                                             created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                                             updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

                                             CONSTRAINT fk_payment
                                             FOREIGN KEY(payment_id)
    REFERENCES payments(id)
    ON DELETE CASCADE
    );