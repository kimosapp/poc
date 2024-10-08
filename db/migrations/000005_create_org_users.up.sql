create table "organization_users"
(
    id                 uuid                     not null
        primary key,
    organization_id    uuid                     not null
        constraint "organization_users_organization_id_fkey"
            references "organizations"
            on update cascade on delete cascade,
    user_id            uuid                     not null,
    role               text,
    created_at         timestamp with time zone not null,
    updated_at         timestamp with time zone not null,
    status             varchar(255),
    role_id            uuid
        constraint organization_users_permission_role_fkey
            references "roles",
    invite_sent_at     timestamp with time zone,
    created_by_user_id uuid,
    is_active          boolean default true
);

create index organization__users_org_id
    on "organization_users" (organization_id);

create unique index organization__users_org_id_user_id_unique
    on "organization_users" (organization_id, user_id);

create index organization__users_organization_id
    on "organization_users" (organization_id);

create index organization__users_user_id
    on "organization_users" (user_id);