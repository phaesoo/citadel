CREATE DATABASE IF NOT EXISTS citadel;
CREATE DATABASE IF NOT EXISTS citadel_test;
GRANT ALL PRIVILEGES ON `citadel`.* TO `citadel_user`@`%`;
GRANT ALL PRIVILEGES ON `citadel_test`.* TO `citadel_user`@`%`;
