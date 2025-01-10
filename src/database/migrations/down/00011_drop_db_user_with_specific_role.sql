REVOKE read_access FROM Blazkowicz;

ALTER DEFAULT PRIVILEGES IN SCHEMA public REVOKE SELECT ON TABLES FROM read_access;
REVOKE SELECT ON ALL TABLES IN SCHEMA public FROM read_access;
REVOKE USAGE ON SCHEMA public FROM read_access;
REVOKE CONNECT ON DATABASE "db-project" FROM read_access;

DROP ROLE read_access;

DROP USER Blazkowicz;
