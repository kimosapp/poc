INSERT INTO "roles"
(id, name, description, editable, created_at, updated_at, created_by, organization_id, level)
VALUES ('b2e1eea3-dd6e-49af-bd16-2bf14bca4492', 'Super Admin', 'Admin platform', false, date ('now'), date ('now'),
        'a6c56de9-1fc0-4215-99a0-97753bb712d8',
        null, 1);
INSERT INTO "roles"
(id, name, description, editable, created_at, updated_at, created_by, organization_id, level)
VALUES ('b2e1eea3-dd6e-49af-bd16-2bf14bca4493', 'Customer Success',
        'Can read but cant not modify settings of the cusotmers or see critical information', false, date ('now'),
        date ('now'), 'a6c56de9-1fc0-4215-99a0-97753bb712d8',
        null, 1);
INSERT INTO "roles"
(id, name, description, editable, created_at, updated_at, created_by, organization_id, level)
VALUES ('b2e1eea3-dd6e-49af-bd16-2bf14bca4494', 'Organization Admin', 'Organization Admin', false, date ('now'), date ('now'),
        'a6c56de9-1fc0-4215-99a0-97753bb712d8',
        null, 2);
INSERT INTO "roles"
(id, name, description, editable, created_at, updated_at, created_by, organization_id, level)
VALUES ('b2e1eea3-dd6e-49af-bd16-2bf14bca4495', 'Organization Member', 'Organization Member', false, date ('now'),
        date ('now'),
        'a6c56de9-1fc0-4215-99a0-97753bb712d8',
        null, 2);