create table if not exists users   (
	user_id serial primary key,
	login varchar(50) not null unique,
	password_hash text not null,
	password_salt text not null
);

create table if not exists subscription_types (
	subscription_type_id serial primary key,
	type_name varchar(50) not null unique,
	duration integer not null unique,
	price integer not null
);

create table if not exists subscriptions (
	subscription_id serial primary key,
	status varchar(20) not null,
	start_date timestamp with time zone not null,
	end_date timestamp with time zone not null,
	subscription_type_id integer references subscription_types(subscription_type_id) not null,
	user_id integer references users(user_id) not null
);

create table if not exists payments (
	payment_id serial primary key,
	token text not null unique,
	subscription_type_name varchar(50) not null,
	amount integer not null,
	start_time timestamp with time zone not null,
	end_time timestamp with time zone not null,
	status varchar(30) not null, 
	user_id integer references users(user_id) not null
);
