CREATE TABLE Users (
  "id" serial NOT NULL,
  "F_name" text NOT NULL,
  "L_name" text NOT NULL,
  "M_name" text NOT NULL,
  PRIMARY KEY (id)
);
CREATE TABLE Address (
  "id" serial NOT NULL,
  "houseNo" integer,
  "streetName" text NOT NULL,
  "pincode" numeric (6,0),
  "state" text,
  "country" text DEFAULT 'India',
  PRIMARY KEY (id)
);
CREATE TABLE User_Add(
  "id" serial NOT NULL PRIMARY KEY,
  "user_id" integer,
  "add_id" integer,
  FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE SET NULL,
  FOREIGN KEY (add_id) REFERENCES Address(id) ON DELETE SET NULL
);
INSERT INTO Address ("id", "houseNo", "streetName", "pincode", "state", "country")
VALUES ('1', '23', 'Sarita', '395002', 'Gujrat', '');
INSERT INTO Users ("id", "F_name", "L_name", "M_name")
VALUES ('1', 'sfd', 'erfer', 'dfer');
INSERT INTO User_Add ("id", "user_id", "add_id")
VALUES ('1', '1', '1');
INSERT INTO Address ("id", "houseNo", "streetName", "pincode", "state", "country")
VALUES ('', '67', 'Varachha', '395002', 'Gujrat', 'India');
SELECT * FROM Users WHERE id = 1;
UPDATE users SET "F_name" = 'aaa', "M_name" = 'bbb', "L_name" = 'ccc' WHERE id = 1;
SELECT "add_id" FROM user_add WHERE "user_id" = 1;
...........................................................sqljoin...............................
select * from "A";
select * from "B";
select p.id, q.name as restaurants, r.name as hotels from "A" as p
inner join "B" as q 
on p.r_id = q.id
inner join "B" as r
on p.h_id = r.id;
or
select "A".id, B1.name1 as res, B1.name2 as hotel 
from "A"
left join (select a.id as id1, a.name as name1, b.id as id2, b.name as name2 from "B" a, "B" b) as B1 
on "A".r_id = B1.id1 and "A".h_id = B1.id2;
