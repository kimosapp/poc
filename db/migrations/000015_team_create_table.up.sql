CREATE TABLE IF NOT EXISTS "teams"
(
    id              UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name            VARCHAR(255),
    image_url       VARCHAR(255),
    slug            VARCHAR(255),
    about           TEXT,
    organization_id UUID        NOT NULL,
    created_at      TIMESTAMPTZ NOT NULL,
    updated_at      TIMESTAMPTZ NOT NULL,
    deleted_at      TIMESTAMPTZ
);

ALTER TABLE "teams"
    ADD CONSTRAINT fk_team_organization
        FOREIGN KEY (organization_id)
            REFERENCES "organizations" (id)
ON DELETE CASCADE;

CREATE INDEX idx_teams_deleted_at ON "teams"(deleted_at);
CREATE INDEX idx_teams_organization_id ON "teams"(organization_id);
CREATE INDEX idx_teams_name_id ON "teams"(name);