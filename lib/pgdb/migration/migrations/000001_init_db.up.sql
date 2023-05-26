--用戶
create table users (
	id uuid primary key NOT NULL DEFAULT uuid_generate_v4(),
	name varchar not null,
    uid varchar not null,
	pwd varchar not null,
	currency varchar not null,
	create_at timestamptz not null default now()
);

--帳目類別
create table entry_types (
	id uuid primary key not null default uuid_generate_v4(),
	name varchar not null,
	icon varchar
);

--帳戶
create table accounts (
	id uuid primary key not null default uuid_generate_v4(),
    user_id uuid not null,
	name varchar not null,
	icon varchar,
    constraint acct_fk foreign key (user_id) references users(id) on delete set null
);

--專案
create table projects (
	id uuid primary key not null default uuid_generate_v4(),
    user_id uuid not null,
	name varchar not null,
	icon varchar,
    constraint prj_fk foreign key (user_id) references users(id) on delete set null
);

--商家
create table stores (
	id uuid primary key not null default uuid_generate_v4(),
    user_id uuid not null,
	name varchar not null,
	icon varchar,
    constraint store_fk foreign key (user_id) references users(id) on delete set null
);

--帳目
create table entries (
	id uuid primary key not null default uuid_generate_v4(),
	user_id uuid not null,
	entry_time timestamptz not null,
	behavior int not null, --金流行為: Pay, Income, Transfer
	amount bigint not null,
	entry_type uuid not null,
	account uuid not null,
	project uuid not null,
	store uuid not null,
	note varchar not null,
	constraint entries_fk_user foreign key (user_id) references users (id) on delete set null,
	constraint entries_fk_type foreign key (entry_type) references entry_types (id) on delete set null,
	constraint entries_fk_acct foreign key (account) references accounts (id) on delete set null,
	constraint entries_fk_prj foreign key (project) references projects (id) on delete set null,
	constraint entries_fk_store foreign key (store) references stores (id) on delete set null
);

create index on entries (user_id);
create index on entries (entry_time);
create index on users (uid, pwd);
create index on accounts (user_id);
create index on projects (user_id);
create index on stores (user_id);