ALTER TABLE
user_invitations
ADD
column expiry TIMESTAMP(0) WITH TIME ZONE NOT NULL;