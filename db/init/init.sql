CREATE USER customer WITH ENCRYPTED PASSWORD '1';
CREATE USER bonus WITH ENCRYPTED PASSWORD '1';
CREATE DATABASE customer;
CREATE DATABASE bonus;
GRANT ALL PRIVILEGES ON DATABASE customer TO customer;


-- need for migrations (issue https://github.com/golang-migrate/migrate/issues/826)
ALTER DATABASE customer OWNER TO customer;
ALTER DATABASE bonus OWNER TO bonus;
--GRANT "customer" TO pg_read_server_files;