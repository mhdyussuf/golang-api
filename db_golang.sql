-- Hapus tabel conversations jika sudah ada
DROP TABLE IF EXISTS conversations;

-- Buat tabel conversations
CREATE TABLE conversations (
  id_conversation BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  created_at DATETIME(3) DEFAULT NULL,
  updated_at DATETIME(3) DEFAULT NULL,
  PRIMARY KEY (id_conversation)
);

-- Hapus tabel users jika sudah ada
DROP TABLE IF EXISTS users;

-- Buat tabel users
CREATE TABLE users (
  id_user BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  username VARCHAR(150) NOT NULL,
  email VARCHAR(100) NOT NULL,
  password VARCHAR(255) NOT NULL,
  created_at DATETIME(3) DEFAULT NULL,
  updated_at DATETIME(3) DEFAULT NULL,
  PRIMARY KEY (id_user),
  UNIQUE (id_user),
  UNIQUE (username),
  UNIQUE (email)
);

-- Hapus tabel messages jika sudah ada
DROP TABLE IF EXISTS messages;

-- Buat tabel messages
CREATE TABLE messages (
  id_message BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  user_id BIGINT UNSIGNED NOT NULL,
  conversation_id BIGINT UNSIGNED NOT NULL,
  content TEXT DEFAULT NULL,
  created_at DATETIME(3) DEFAULT NULL,
  updated_at DATETIME(3) DEFAULT NULL,
  PRIMARY KEY (id_message),
  INDEX fk_users_message (user_id),
  INDEX fk_conversations_messages (conversation_id),
  FOREIGN KEY (conversation_id) REFERENCES conversations (id_conversation) ON DELETE CASCADE ON UPDATE RESTRICT,
  FOREIGN KEY (user_id) REFERENCES users (id_user) ON DELETE CASCADE ON UPDATE RESTRICT
);
