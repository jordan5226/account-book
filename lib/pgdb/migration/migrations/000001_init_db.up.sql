--用戶
create table if not exists users (
	id uuid primary key not null,
	name varchar not null,
    uid varchar not null,
	pwd varchar not null,
	currency varchar not null,
	create_at timestamptz not null default now()
);

--帳目類別
create table if not exists types (
	id uuid primary key not null,
	name varchar not null,
	icon varchar
);

--帳戶
create table if not exists accounts (
	id uuid primary key not null,
    user_id uuid not null,
	name varchar not null,
	icon varchar,
    constraint acct_fk foreign key (user_id) references users(id) on delete set null
);

--專案
create table if not exists projects (
	id uuid primary key not null,
    user_id uuid not null,
	name varchar not null,
	icon varchar,
    constraint prj_fk foreign key (user_id) references users(id) on delete set null
);

--商家
create table if not exists stores (
	id uuid primary key not null,
    user_id uuid not null,
	name varchar not null,
	icon varchar,
    constraint store_fk foreign key (user_id) references users(id) on delete set null
);

--帳目
create table if not exists entries (
	id uuid primary key not null,
	user_id uuid not null,
	time timestamptz not null,
	behavior int not null, --金流行為: Pay, Income, Transfer
	amount bigint not null,
	type uuid not null,
	account uuid not null,
	project uuid not null,
	store uuid not null,
	note varchar not null,
	constraint entries_fk_user foreign key (user_id) references users (id) on delete set null,
	constraint entries_fk_type foreign key (type) references types (id) on delete set null,
	constraint entries_fk_acct foreign key (account) references accounts (id) on delete set null,
	constraint entries_fk_prj foreign key (project) references projects (id) on delete set null,
	constraint entries_fk_store foreign key (store) references stores (id) on delete set null
);

create index on entries (user_id);
create index on entries (time);
create index on users (uid, pwd);
create index on accounts (user_id);
create index on projects (user_id);
create index on stores (user_id);