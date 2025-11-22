-- Pan Search Platform Database Design - Minimal Version
-- Database: pan_search
-- Character Set: utf8mb4
-- Collation: utf8mb4_unicode_ci

-- Create database
CREATE DATABASE IF NOT EXISTS `pan_search` 
CHARACTER SET utf8mb4 
COLLATE utf8mb4_unicode_ci;

USE `pan_search`;

-- 1. Categories table
CREATE TABLE `categories` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'Category ID',
  `value` VARCHAR(50) NOT NULL COMMENT 'Category value',
  `label` VARCHAR(50) NOT NULL COMMENT 'Category label',
  `icon` VARCHAR(20) NOT NULL COMMENT 'Icon',
  `sort_order` INT NOT NULL DEFAULT 0 COMMENT 'Sort order',
  `is_active` TINYINT(1) NOT NULL DEFAULT 1 COMMENT 'Is active',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Created time',
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Updated time',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_value` (`value`),
  KEY `idx_sort_order` (`sort_order`),
  KEY `idx_is_active` (`is_active`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Categories table';

-- 2. Resources table
CREATE TABLE `resources` (
  `id` VARCHAR(32) NOT NULL COMMENT 'Resource ID',
  `title` VARCHAR(255) NOT NULL COMMENT 'Resource title',
  `description` TEXT COMMENT 'Resource description',
  `size` VARCHAR(20) NOT NULL COMMENT 'File size',
  `type` VARCHAR(50) NOT NULL COMMENT 'Resource type',
  `category_id` INT UNSIGNED NOT NULL COMMENT 'Category ID',
  `source` VARCHAR(100) NOT NULL COMMENT 'Source platform',
  `download_url` VARCHAR(500) NOT NULL COMMENT 'Download URL',
  `extract_code` VARCHAR(20) COMMENT 'Extract code',
  `file_count` INT UNSIGNED NOT NULL DEFAULT 1 COMMENT 'File count',
  `view_count` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'View count',
  `download_count` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Download count',
  `valid` TINYINT(1) NOT NULL DEFAULT 1 COMMENT 'Is valid',
  `expire_time` DATETIME COMMENT 'Expire time',
  `upload_time` DATETIME NOT NULL COMMENT 'Upload time',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Created time',
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Updated time',
  PRIMARY KEY (`id`),
  KEY `idx_category_id` (`category_id`),
  KEY `idx_type` (`type`),
  KEY `idx_source` (`source`),
  KEY `idx_upload_time` (`upload_time`),
  KEY `idx_view_count` (`view_count`),
  KEY `idx_download_count` (`download_count`),
  KEY `idx_valid` (`valid`),
  KEY `idx_expire_time` (`expire_time`),
  FULLTEXT KEY `ft_title_description` (`title`, `description`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Resources table';

-- 3. Resource tags table
CREATE TABLE `resource_tags` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'Tag ID',
  `resource_id` VARCHAR(32) NOT NULL COMMENT 'Resource ID',
  `tag_name` VARCHAR(50) NOT NULL COMMENT 'Tag name',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Created time',
  PRIMARY KEY (`id`),
  KEY `idx_resource_id` (`resource_id`),
  KEY `idx_tag_name` (`tag_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Resource tags table';

-- 4. Help requests table
CREATE TABLE `help_requests` (
  `id` VARCHAR(32) NOT NULL COMMENT 'Request ID',
  `resource_name` VARCHAR(255) NOT NULL COMMENT 'Resource name',
  `resource_type` VARCHAR(50) COMMENT 'Resource type',
  `description` TEXT COMMENT 'Description',
  `contact` VARCHAR(100) NOT NULL COMMENT 'Contact info',
  `contact_type` ENUM('email', 'phone') NOT NULL COMMENT 'Contact type',
  `status` ENUM('pending', 'processing', 'completed', 'rejected') NOT NULL DEFAULT 'pending' COMMENT 'Status',
  `admin_notes` TEXT COMMENT 'Admin notes',
  `completed_time` DATETIME COMMENT 'Completed time',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Created time',
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Updated time',
  PRIMARY KEY (`id`),
  KEY `idx_status` (`status`),
  KEY `idx_contact_type` (`contact_type`),
  KEY `idx_created_at` (`created_at`),
  KEY `idx_resource_type` (`resource_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Help requests table';

-- 5. Download records table
CREATE TABLE `download_records` (
  `id` VARCHAR(32) NOT NULL COMMENT 'Download record ID',
  `resource_id` VARCHAR(32) NOT NULL COMMENT 'Resource ID',
  `user_id` VARCHAR(32) COMMENT 'User ID',
  `user_agent` TEXT COMMENT 'User agent',
  `ip_address` VARCHAR(45) COMMENT 'IP address',
  `download_time` DATETIME NOT NULL COMMENT 'Download time',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Created time',
  PRIMARY KEY (`id`),
  KEY `idx_resource_id` (`resource_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_download_time` (`download_time`),
  KEY `idx_ip_address` (`ip_address`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Download records table';

-- 6. Search records table
CREATE TABLE `search_records` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'Search record ID',
  `keyword` VARCHAR(255) NOT NULL COMMENT 'Search keyword',
  `category` VARCHAR(50) COMMENT 'Category',
  `sort_by` VARCHAR(50) COMMENT 'Sort by',
  `user_id` VARCHAR(32) COMMENT 'User ID',
  `ip_address` VARCHAR(45) COMMENT 'IP address',
  `user_agent` TEXT COMMENT 'User agent',
  `result_count` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Result count',
  `search_time` DATETIME NOT NULL COMMENT 'Search time',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Created time',
  PRIMARY KEY (`id`),
  KEY `idx_keyword` (`keyword`),
  KEY `idx_category` (`category`),
  KEY `idx_search_time` (`search_time`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Search records table';

-- 7. Users table (reserved for future extension)
CREATE TABLE `users` (
  `id` VARCHAR(32) NOT NULL COMMENT 'User ID',
  `username` VARCHAR(50) NOT NULL COMMENT 'Username',
  `email` VARCHAR(100) NOT NULL COMMENT 'Email',
  `phone` VARCHAR(20) COMMENT 'Phone',
  `password_hash` VARCHAR(255) NOT NULL COMMENT 'Password hash',
  `avatar` VARCHAR(500) COMMENT 'Avatar URL',
  `status` ENUM('active', 'inactive', 'banned') NOT NULL DEFAULT 'active' COMMENT 'User status',
  `last_login_at` DATETIME COMMENT 'Last login time',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Created time',
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Updated time',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_username` (`username`),
  UNIQUE KEY `uk_email` (`email`),
  UNIQUE KEY `uk_phone` (`phone`),
  KEY `idx_status` (`status`),
  KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Users table';

-- 8. System configs table
CREATE TABLE `system_configs` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'Config ID',
  `config_key` VARCHAR(100) NOT NULL COMMENT 'Config key',
  `config_value` TEXT NOT NULL COMMENT 'Config value',
  `config_type` ENUM('string', 'number', 'boolean', 'json') NOT NULL DEFAULT 'string' COMMENT 'Config type',
  `description` VARCHAR(255) COMMENT 'Description',
  `is_public` TINYINT(1) NOT NULL DEFAULT 0 COMMENT 'Is public',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Created time',
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Updated time',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_config_key` (`config_key`),
  KEY `idx_is_public` (`is_public`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='System configs table';

-- Insert initial data

-- Insert categories data
INSERT INTO `categories` (`value`, `label`, `icon`, `sort_order`) VALUES
('all', '全部', '📁', 0),
('movie', '电影', '🎬', 1),
('tv', '电视剧', '📺', 2),
('music', '音乐', '🎵', 3),
('software', '软件', '💻', 4),
('document', '文档', '📄', 5),
('other', '其他', '📦', 6);

-- Insert system configs
INSERT INTO `system_configs` (`config_key`, `config_value`, `config_type`, `description`, `is_public`) VALUES
('site_title', '网盘资源搜索', 'string', '网站标题', 1),
('site_description', '专业的网盘资源搜索平台', 'string', '网站描述', 1),
('page_size', '10', 'number', '默认分页大小', 1),
('search_debounce', '500', 'number', '搜索防抖时间(ms)', 0),
('upload_size_limit', '10', 'number', '上传文件大小限制(MB)', 0),
('customer_service_email', 'support@pansearch.com', 'string', '客服邮箱', 1),
('customer_service_phone', '400-123-4567', 'string', '客服电话', 1),
('icp_beian', '京ICP备12345678号', 'string', '备案信息', 1);

-- Show table structure info
SHOW TABLES;

-- Show table records count
SELECT 
    TABLE_NAME as '表名',
    TABLE_ROWS as '记录数',
    DATA_LENGTH as '数据大小(B)',
    INDEX_LENGTH as '索引大小(B)',
    CREATE_TIME as '创建时间'
FROM information_schema.TABLES 
WHERE TABLE_SCHEMA = 'pan_search'
ORDER BY TABLE_NAME;

-- Database connection info memo
-- Host: localhost
-- Port: 3306
-- Database: pan_search
-- Username: root
-- Password: (empty)
-- Charset: utf8mb4