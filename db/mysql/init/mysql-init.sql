CREATE DATABASE IF NOT EXISTS keybox;
CREATE DATABASE IF NOT EXISTS keybox_test;
GRANT ALL PRIVILEGES ON `keybox`.* TO `keybox_user`@`%`;
GRANT ALL PRIVILEGES ON `keybox_test`.* TO `keybox_user`@`%`;
