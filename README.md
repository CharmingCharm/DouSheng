# DouSheng项目

## 技术选型

数据库：mysql + minio
微服务：kitex + ETCD
主要框架：gin + gorm + kitex

## 数据表

```sql
CREATE TABLE IF NOT EXISTS  `user` (
    `id` BIGINT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(32) UNIQUE NOT NULL COMMENT '用户名',
    `password` VARCHAR(32) NOT NULL COMMENT '密码',
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

CREATE TABLE IF NOT EXISTS `video` (
    `id` BIGINT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `title` VARCHAR(255) NOT NULL COMMENT '视频标题',
    `author_id` BIGINT(10) UNSIGNED NOT NULL COMMENT '发布者用户id',
    `play_url` VARCHAR(255) NOT NULL COMMENT '播放地址',
    `cover_url` VARCHAR(255) NOT NULL COMMENT '封面地址',
    `created_on` DATETIME NOT NULL COMMENT '创建时间',
    INDEX `author`(author_id),
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='视频表';

CREATE TABLE IF NOT EXISTS `favorite` (
    `id` BIGINT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT(10) UNSIGNED NOT NULL COMMENT '点赞用户id',
    `video_id` BIGINT(10) UNSIGNED NOT NULL COMMENT '点赞视频id',
    INDEX `user_video`(user_id, video_id),
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='点赞表';

CREATE TABLE IF NOT EXISTS `relation` (
    `id` BIGINT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT(10) UNSIGNED NOT NULL COMMENT '被关注用户id',
    `follower_id` BIGINT(10) UNSIGNED NOT NULL COMMENT '发起关注请求用户id',
    INDEX `user_follower`(user_id, follower_id),
    INDEX `follower_user`(follower_id, user_id),
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户关系表';

CREATE TABLE IF NOT EXISTS `comment` (
    `id` BIGINT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT(10) UNSIGNED NOT NULL COMMENT '评论用户id',
    `video_id` BIGINT(10) UNSIGNED NOT NULL COMMENT '评论视频id',
    `comment_text` text COMMENT '评论内容',
    `created_on` DATETIME NOT NULL COMMENT '创建时间',
    INDEX `video`(video_id),
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='评论表';

CREATE TABLE IF NOT EXISTS `user_count` (
    `id` BIGINT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT(10) UNSIGNED UNIQUE NOT NULL COMMENT '用户id',
    `follow_count` BIGINT(10) UNSIGNED NOT NULL COMMENT '关注数',
    `follower_count` BIGINT(10) UNSIGNED NOT NULL COMMENT '粉丝数',
    INDEX `user`(user_id),
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户数量信息表';

CREATE TABLE IF NOT EXISTS `video_count` (
    `id` BIGINT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `video_id` BIGINT(10) UNSIGNED UNIQUE NOT NULL COMMENT '视频id',
    `favorite_count` BIGINT(10) UNSIGNED NOT NULL COMMENT '点赞数',
    `comment_count` BIGINT(10) UNSIGNED NOT NULL COMMENT '评论数',
    INDEX `video`(video_id),
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='视频数量信息表';
```

## 完成的功能

- Rigister
- Login
- UserInfo
- Publish
- PublishList
- Feed
- Favorite
- FavoriteList
- Comment
- CommentList
- Follow
- FollowList
- FollowerList

## 项目架构

采用微服务的项目架构，拆分为User、Video、Action三个微服务。使用Thrift作为RPC的通信协议。

通过ETCD进行服务的注册发现，通过jaeger进行分布式链路追踪，通过minio上传并存储视频数据。

通过gin在api目录下对每个功能设置了一个路由，对应相应的请求在路由handler处通过jwt中间件进行用户身份的认证，接着对请求进行解析将其通过RPC调用对应的微服务，得到结果后返回给客户端。

## 如何运行

- Start 5 terminals in DouSheng
- Terminal 1: docker-compose up
- Terminal 2:

1. cd internal/user
2. sh build.sh
3. sh output/bootstrap.sh

- Terminal 3:

1. cd internal/video
2. sh build.sh
3. sh output/bootstrap.sh

- Terminal 4:

1. cd internal/action
2. sh build.sh
3. sh output/bootstrap.sh

- Terminal 5:

1. cd internal/api
2. go build && ./api
