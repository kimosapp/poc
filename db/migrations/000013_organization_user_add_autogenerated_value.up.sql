ALTER TABLE "organization_users"
    ALTER COLUMN id SET DEFAULT uuid_generate_v4()