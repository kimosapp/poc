CREATE
EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "organizations"
(
    id                       uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name                     VARCHAR(255),
    created_by               uuid,
    slug                     VARCHAR(255),
    billing_email            VARCHAR(255),
    url                      VARCHAR(255),
    about                    TEXT,
    logo_url                 VARCHAR(255),
    background_image_url     VARCHAR(255),
    plan                     VARCHAR(255),
    current_period_starts_at TIMESTAMP WITH TIME ZONE,
    current_period_ends_at   TIMESTAMP WITH TIME ZONE,
    subscription_id          VARCHAR(255),
    status                   VARCHAR(255),
    timezone                 VARCHAR(255),
    created_at               TIMESTAMP WITH TIME ZONE        NOT NULL,
    updated_at               TIMESTAMP WITH TIME ZONE        NOT NULL,
    deleted_at               TIMESTAMP WITH TIME ZONE
);

CREATE UNIQUE INDEX lower_organizations_unique_idx
    ON "organizations" (LOWER(slug::TEXT));

CREATE TABLE "users"
(
    id                             uuid    DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    email                          VARCHAR(255),
    organization_id                uuid                               NOT NULL,
    accept_terms_and_conditions    BOOLEAN DEFAULT FALSE              NOT NULL,
    hash                           VARCHAR(255),
    first_name                     VARCHAR(255),
    last_name                      VARCHAR(255),
    created_at                     TIMESTAMP WITH TIME ZONE           NOT NULL,
    updated_at                     TIMESTAMP WITH TIME ZONE           NOT NULL,
    accept_terms_and_conditions_at TIMESTAMP WITH TIME ZONE           NOT NULL,
    photo_url                      VARCHAR(255),
    timezone                       VARCHAR(255),
    phone                          VARCHAR(255),
    role                           VARCHAR(255),
    status                         varchar(255),
    created_by_user_id             uuid,
    request_password               BOOLEAN DEFAULT TRUE               NOT NULL,
    has_password                   BOOLEAN DEFAULT FALSE              NOT NULL,
    password_expired_at            TIMESTAMP WITH TIME ZONE,
    invited_at                     timestamp with time zone,
    is_active                      boolean default true               NOT NULL,
    email_verified_at              TIMESTAMP WITH TIME ZONE,
    is_locked                      BOOLEAN DEFAULT FALSE              NOT NULL,
    bad_attempts                   INTEGER DEFAULT 0                  NOT NULL,
    last_login                     TIMESTAMP WITH TIME ZONE,
    deleted_at                     TIMESTAMP WITH TIME ZONE
);


CREATE INDEX idx_organizations_id
    ON "organizations" (id);

CREATE INDEX email_unique_idx
    ON "users" (LOWER(email::TEXT));

CREATE INDEX idx_users_id
    ON "users" (id);

CREATE INDEX users_email
    ON "users" (email);

CREATE UNIQUE INDEX users_email_tenant
    ON "users" (email, organization_id);