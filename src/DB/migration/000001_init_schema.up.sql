CREATE TABLE Addresses (
  country varchar NOT NULL,
  city varchar NOT NULL,
  street varchar NOT NULL,
  PRIMARY KEY (country, city, street)
);

CREATE TABLE University (
  u_id bigserial PRIMARY KEY,
  university_name varchar UNIQUE NOT NULL,
  abbreviation varchar NOT NULL,
  email_extension varchar NOT NULL,
  country varchar NOT NULL,
  city varchar NOT NULL,
  street varchar NOT NULL,
  FOREIGN KEY (country, city, street) REFERENCES Addresses ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE Student (
  s_id bigserial PRIMARY KEY,
  user_name varchar UNIQUE NOT NULL,
  hashed_password varchar UNIQUE NOT NULL,
  full_name varchar NOT NULL,
  email varchar UNIQUE NOT NULL,
  description varchar,
  avatar varchar,
  u_id bigserial NOT NULL,
  CONSTRAINT fk_student FOREIGN KEY (u_id) REFERENCES University(u_id)  ON UPDATE CASCADE ON DELETE CASCADE,
  credit int NOT NULL
);

CREATE TABLE Subscription (
    sub_id bigserial PRIMARY KEY,
    s_id bigserial NOT NULL,
    plan_type varchar NOT NULL,
    sub_expire_time varchar NOT NULL,
    sub_start_time varchar NOT NULL,
    FOREIGN KEY (s_id) REFERENCES Student (s_id)  ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE Course (
  c_id bigserial PRIMARY KEY,
  course_name varchar NOT NULL,
  semester varchar NOT NULL,
  description varchar NOT NULL,
  u_id bigserial NOT NULL,
  FOREIGN KEY (u_id) REFERENCES University (u_id)  ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE Student_Courses (
  s_id bigserial,
  c_id bigserial,
  PRIMARY KEY (s_id, c_id),
  FOREIGN KEY (s_id) REFERENCES Student (s_id) ON UPDATE CASCADE ON DELETE CASCADE,
  FOREIGN KEY (c_id) REFERENCES Course (c_id) ON UPDATE CASCADE ON DELETE CASCADE
);


CREATE TABLE Mentor (
  m_id bigserial PRIMARY KEY,
  c_id bigserial NOT NULL,
  username varchar NOT NULL,
  full_name varchar NOT NULL,
  hashed_password varchar UNIQUE NOT NULL,
  email varchar NOT NULL,
  description varchar NOT NULL,
  evaluation_count int NOT NULL,
  score int NOT NULL,
  balance int NOT NULL,
  u_id bigserial NOT NULL,
  country varchar NOT NULL,
  city varchar NOT NULL,
  street varchar NOT NULL,
  FOREIGN KEY (country, city, street) REFERENCES Addresses ON UPDATE CASCADE ON DELETE CASCADE,
  FOREIGN KEY (c_id) REFERENCES Course (c_id)  ON UPDATE CASCADE ON DELETE CASCADE,
  FOREIGN KEY (u_id) REFERENCES University (u_id)  ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE INDEX "full_address" ON Addresses ("country", "city", "street");

CREATE UNIQUE INDEX ON Course ("c_id", "semester");
