CREATE TABLE IF NOT EXISTS users(
  id uuid PRIMARY KEY,
  auth0_id text UNIQUE NOT NULL,
  email text UNIQUE NOT NULL,
  username text UNIQUE NOT NULL,
  first_name text NOT NULL,
  last_name text NOT NULL,
  is_deleted boolean DEFAULT FALSE NOT NULL,
  deleted_at timestamp with time zone,
  created_at timestamp with time zone DEFAULT timezone('utc'::text, now()) NOT NULL,
  updated_at timestamp with time zone)
