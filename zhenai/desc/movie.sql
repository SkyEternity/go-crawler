CREATE TABLE `zhenai_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) DEFAULT '' COMMENT '名称',
  `city` varchar(30) DEFAULT '' COMMENT '城市',
  `salary` varchar(30) DEFAULT '' COMMENT '收入',
  `education` varchar(30) DEFAULT '' COMMENT '学历',
  `age` int(10) unsigned DEFAULT '0' COMMENT '年龄',
  `height` int(10) unsigned DEFAULT '0' COMMENT '身高',
  `introduce` varchar(200) DEFAULT '' COMMENT '介绍',
  `cover` varchar(200) DEFAULT '' COMMENT '头像',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='珍爱网人物信息';

