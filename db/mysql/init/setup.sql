CREATE DATABASE IF NOT EXISTS simple_chat_dev;
CREATE USER IF NOT EXISTS 'simple_chat_dev'@'%' IDENTIFIED BY 'simple_chat_pw';
GRANT ALL PRIVILEGES ON simple_chat_dev.* TO 'simple_chat_dev'@'%';
FLUSH PRIVILEGES;