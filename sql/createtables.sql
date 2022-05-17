create table product_user (
  user_id int PRIMARY KEY,
  user_name text,
  channel_id text
);

create table reminder (
  remind_id int PRIMARY KEY,
  title text,
  reminder_message text,
  time_to_remind timestamp
);
