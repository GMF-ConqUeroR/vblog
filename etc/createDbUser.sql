# 创建用户
CREATE USER 'blog'@'%' IDENTIFIED BY 'blog123';

# 为该用户附上所有数据库的所有权限
GRANT ALL PRIVILEGES ON *.* TO 'blog'@'%';

# 为该用户附上指定数据库的所有权限
GRANT ALL PRIVILEGES ON blog.* TO 'blog'@'%';

# 为该用户附上指定数据库的指定权限
GRANT SELECT, INSERT, UPDATE ON blog.* TO 'blog'@'%';

FLUSH PRIVILEGES;
# 确认 'blog' 用户是否有从该 IP 地址连接到数据库的权限。可以登录MySQL的root账户，检查用户权限：
SELECT host, user, authentication_string FROM mysql.user WHERE user = 'blog';

CREATE DATABASE `blog` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

-- blog.blog definition

CREATE TABLE `blog` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_at` bigint(20) DEFAULT NULL,
  `update_at` bigint(20) DEFAULT NULL,
  `publish_at` bigint(20) DEFAULT NULL,
  `title` longtext,
  `author` longtext,
  `content` longtext,
  `tags` longtext,
  `status` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;