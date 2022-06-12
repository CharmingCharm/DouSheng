# DouSheng项目

## 技术选型

数据库：MySQL + MinIO

微服务：kitex + ETCD + thrift

主要框架：gin + GORM + kitex

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

通过ETCD进行服务的注册发现，通过jaeger进行分布式链路追踪，通过MinIO上传并存储视频数据。

通过gin在api目录下对每个功能设置了一个路由，对应相应的请求在路由handler处通过jwt中间件进行用户身份的认证，接着对请求进行解析将其通过RPC调用对应的微服务，得到结果后返回给客户端。
```
.
├─idl
│      action.thrift
│      base.thrift
│      user.thrift
│      video.thrift
│      
├─internal
│  │  .DS_Store
│  │  
│  ├─action
│  │  │  build.sh
│  │  │  handler.go
│  │  │  main.go
│  │  │  
│  │  ├─db
│  │  │      comment.go
│  │  │      favorite.go
│  │  │      init.go
│  │  │      relation.go
│  │  │      
│  │  ├─rpc
│  │  │      init.go
│  │  │      user.go
│  │  │      video.go
│  │  │      
│  │  ├─script
│  │  │      bootstrap.sh
│  │  │      
│  │  └─service
│  │          checkFavorite.go
│  │          checkRelation.go
│  │          getCommentList.go
│  │          getFavoriteVideos.go
│  │          getUserFollowerList.go
│  │          getUserFollowList.go
│  │          updateComment.go
│  │          updateFavorite.go
│  │          updateRelationship.go
│  │          
│  ├─api
│  │  │  main.go
│  │  │  router.go
│  │  │  
│  │  ├─controller
│  │  │      comment.go
│  │  │      favorite.go
│  │  │      feed.go
│  │  │      publish.go
│  │  │      relation.go
│  │  │      user.go
│  │  │      
│  │  ├─ostg
│  │  │      minio.go
│  │  │      
│  │  └─rpc
│  │          action.go
│  │          init.go
│  │          user.go
│  │          video.go
│  │          
│  ├─user
│  │  │  build.sh
│  │  │  handler.go
│  │  │  main.go
│  │  │  
│  │  ├─db
│  │  │      init.go
│  │  │      user.go
│  │  │      
│  │  ├─rpc
│  │  │      action.go
│  │  │      init.go
│  │  │      
│  │  ├─script
│  │  │      bootstrap.sh
│  │  │      
│  │  └─service
│  │          checkUser.go
│  │          createUser.go
│  │          getUserInfo.go
│  │          updateUserFollow.go
│  │          
│  └─video
│      │  build.sh
│      │  handler.go
│      │  main.go
│      │  
│      ├─db
│      │      init.go
│      │      video.go
│      │      
│      ├─rpc
│      │      action.go
│      │      init.go
│      │      user.go
│      │      
│      ├─script
│      │      bootstrap.sh
│      │      
│      └─service
│              getPublishedVideos.go
│              getVideoList.go
│              loadVideos.go
│              publishVideo.go
│              updateCommentCount.go
│              updateFavoriteCount.go
│              
├─kitex_gen
│  ├─action
│  │  │  action.go
│  │  │  k-action.go
│  │  │  k-consts.go
│  │  │  
│  │  └─actionservice
│  │          actionservice.go
│  │          client.go
│  │          invoker.go
│  │          server.go
│  │          
│  ├─base
│  │      base.go
│  │      k-base.go
│  │      k-consts.go
│  │      
│  ├─user
│  │  │  k-consts.go
│  │  │  k-user.go
│  │  │  user.go
│  │  │  
│  │  └─userservice
│  │          client.go
│  │          invoker.go
│  │          server.go
│  │          userservice.go
│  │          
│  └─video
│      │  k-consts.go
│      │  k-video.go
│      │  video.go
│      │  
│      └─videoservice
│              client.go
│              invoker.go
│              server.go
│              videoservice.go
│              
├─pkg
│  ├─constants
│  │      constants.go
│  │      
│  ├─middleware
│  │      jwt.go
│  │      
│  ├─response
│  │      response.go
│  │      
│  ├─send
│  │      send.go
│  │      
│  └─status
│          status.go
│          
└─public
        bear.mp4
        data
```

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
