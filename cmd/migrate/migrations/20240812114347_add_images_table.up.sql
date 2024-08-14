create table if not exists images(
  id serial primary key, /* if you were fetching by id, it would probably be better to use a uuid to make it not guessable */
  user_id uuid references auth.users,
  status int not null default 1, /* 1 is pending */
  external_id text, /* cloud storage image id */
  prompt text not null,
  deleted boolean not null default 'false', /* soft delete */
  deleted_at timestamp,
  created_at timestamp not null default now(),
  location text
);
