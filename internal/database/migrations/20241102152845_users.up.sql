create table users (
  id varchar(50) not null,
  username varchar(255) not null unique,
  "displayName" varchar(255) not null,
  email varchar(255) not null unique,
  picture varchar(255),
  "createdAt" timestamp default current_timestamp,
  "updatedAt" timestamp not null
);
