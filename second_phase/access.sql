CREATE USER blazkowicz WITH PASSWORD 'William1939';

CREATE ROLE read_access LOGIN;

GRANT CONNECT ON DATABASE "db-project" TO read_access;
GRANT USAGE ON SCHEMA public TO read_access;
GRANT SELECT ON ALL TABLES IN SCHEMA public TO read_access;

ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT ON TABLES TO read_access;

GRANT read_access TO Blazkowicz;

