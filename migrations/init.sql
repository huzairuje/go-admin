CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE users (
       user_id VARCHAR(100) PRIMARY KEY DEFAULT uuid_generate_v4(),
       nik VARCHAR(10),
       name VARCHAR(50),
       email VARCHAR(100) UNIQUE,
       password VARCHAR(100),
       is_active INT,
       type_user INT,
       user_role_id VARCHAR(50),
       created_at TIMESTAMPTZ,
       updated_at TIMESTAMPTZ
);
CREATE TABLE web_menu (
      id VARCHAR(100) PRIMARY KEY DEFAULT uuid_generate_v4(),
      name VARCHAR(50),
      parent_id VARCHAR(50),
      route VARCHAR(50),
      icon_class VARCHAR(50),
      is_active VARCHAR(2),
      role VARCHAR(255),
      prefix VARCHAR(50),
      menu_type VARCHAR(20),
      created_at TIMESTAMPTZ default now()
);

CREATE TABLE web_menu_parent (
     id VARCHAR(100) PRIMARY KEY DEFAULT uuid_generate_v4(),
     name VARCHAR(60),
     parent_id VARCHAR(50),
     is_active VARCHAR(255),
     user_role_id VARCHAR(255),
     icon_parent VARCHAR(255),
     "order" INT,
     created_at TIMESTAMPTZ default now()
);

CREATE TABLE menu_role (
       id VARCHAR(100) PRIMARY KEY DEFAULT uuid_generate_v4(),
       menu_id VARCHAR(60),
       user_role_id VARCHAR(50),
       create INT,
       read INT,
       update INT,
       delete INT,
       upload INT,
       created_at TIMESTAMPTZ default now()
);

CREATE TABLE user_role (
       id VARCHAR(100) PRIMARY KEY DEFAULT uuid_generate_v4(),
       name VARCHAR(50),
       description VARCHAR(255),
       created_at TIMESTAMPTZ default now()
);