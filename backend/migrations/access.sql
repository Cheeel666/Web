grant all privileges on database agregator to chel;
grant connect on database agregator to chel;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public to chel;
alter user chel with superuser;