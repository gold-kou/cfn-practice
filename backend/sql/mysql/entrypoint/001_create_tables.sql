CREATE TABLE `messages` (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'サロゲートキー',
    `message` VARCHAR(255) NOT NULL COMMENT 'メッセージ'
)COMMENT 'メッセージテーブル';