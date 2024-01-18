CREATE USER customer WITH ENCRYPTED PASSWORD '1';
CREATE DATABASE customer;
GRANT ALL PRIVILEGES ON DATABASE customer TO customer;


-- need for migrations (issue https://github.com/golang-migrate/migrate/issues/826)
ALTER DATABASE customer OWNER TO customer;
GRANT "customer" TO pg_read_server_files;