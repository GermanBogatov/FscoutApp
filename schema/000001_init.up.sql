CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE Role
(
    role_uuid UUID NOT NULL DEFAULT uuid_generate_v4() unique,
    name varchar(255)
);

insert into role(name) VALUES ('sportsman');
insert into role(name) VALUES ('scout');
insert into role(name) VALUES ('admin');

CREATE TABLE Sportsman
(
  sportsman_uuid UUID NOT NULL DEFAULT uuid_generate_v4() unique,
  name varchar(255) not null,
  surname varchar(255) not null,
  phone varchar(255) not null unique,
  email varchar(255) not null unique,
  password varchar(255) not null,
  academy varchar(255),
  country_uuid int,
  address varchar(255),
  city varchar(255),
  state varchar(255),
  index int,
  birthday timestamp,
  height varchar(255),
  weight varchar(255),
  strong_leg varchar(255),
  time_create timestamp,
  role_uuid uuid references Role(role_uuid) on delete cascade not null
);

