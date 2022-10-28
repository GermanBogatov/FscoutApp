CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE Roles
(
    role_uuid UUID NOT NULL DEFAULT uuid_generate_v4() unique,
    name varchar(255)
);

insert into Roles(name) VALUES ('sportsman');
insert into Roles(name) VALUES ('scout');
insert into Roles(name) VALUES ('admin');

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
  zipcode int,
  birthday timestamp,
  gender varchar(255),
  height varchar(255),
  weight varchar(255),
  strong_leg varchar(255),
  time_create timestamp,
  role_uuid uuid references Roles(role_uuid) on delete cascade not null
);

CREATE TABLE Scout
(
    scout_uuid UUID NOT NULL DEFAULT uuid_generate_v4() unique,
    name varchar(255) not null,
    surname varchar(255) not null,
    phone varchar(255) not null unique,
    email varchar(255) not null unique,
    password varchar(255) not null,
    company varchar(255),
    country_uuid int,
    address varchar(255),
    city varchar(255),
    state varchar(255),
    zipcode int,
    gender varchar(255),
    vat_number varchar(255),
    passport varchar(255),
    time_create timestamp,
    role_uuid uuid references Roles(role_uuid) on delete cascade not null,
    confirmed bool
);
