CREATE TABLE "role_permissions"
(
    id              uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    "role_id"       uuid                            NOT NULL,
    "permission_id" uuid                         NOT NULL,
    FOREIGN KEY ("role_id") REFERENCES "roles" ("id") ON DELETE CASCADE,
    FOREIGN KEY ("permission_id") REFERENCES "permissions" ("id") ON DELETE CASCADE
);

-- Create index for role_id and permission_id
CREATE INDEX "role_permissions_role_id_idx" ON "role_permissions" ("role_id");
CREATE INDEX "role_permissions_permission_id_idx" ON "role_permissions" ("permission_id");