-- ROLLBACK;
BEGIN;
create table if not exists subjects
(
  id uuid default gen_random_uuid() not null constraint subjects_pk primary key
);

create unique index if not exists unique_subjects_id on subjects (id);

create table if not exists users
(
  id              uuid PRIMARY KEY constraint fk_subject_user_id references subjects on update cascade on delete restrict,
  identifier      varchar not null,
  password_hashed varchar(2048) not null
);

create unique index if not exists unique_users_identifier on users (identifier);

create table if not exists roles
(
  id   uuid PRIMARY KEY constraint fk_subject_role_id references subjects on update cascade on delete restrict,
  name varchar
);

create unique index if not exists unique_roles_name on roles (name);

create table if not exists permissions
(
  id   uuid default gen_random_uuid() not null constraint permissions_pk primary key,
  name varchar
);

create unique index if not exists unique_permissions_name on permissions (name);

create table if not exists users_roles
(
  id      uuid default gen_random_uuid() not null constraint users_roles_pk primary key,
  user_id uuid                           not null constraint fk_user_id references users on update cascade on delete restrict,
  role_id uuid                           not null constraint fk_roles_id references roles on update cascade on delete restrict
);

create unique index if not exists unique_users_roles_user_role on users_roles (user_id, role_id);

create table if not exists applications
(
  id   uuid default gen_random_uuid() not null constraint applications_pk primary key,
  name varchar                       not null
);

create unique index if not exists unique_applications_name on applications (name);

create table if not exists polices
(
  id            uuid default gen_random_uuid() not null constraint polices_pk primary key,
  aplication_id uuid constraint fk_aplication_id references applications on update cascade on delete restrict,
  permission_id uuid constraint fk_permission_id references permissions  on update cascade on delete restrict,
  subject_id    uuid constraint fk_subject_id    references subjects     on update cascade on delete restrict,
  created_at    timestamp with time zone default CURRENT_TIMESTAMP,
  updated_at    timestamp with time zone default CURRENT_TIMESTAMP
);

INSERT INTO applications (name) VALUES ('default');
COMMIT;
