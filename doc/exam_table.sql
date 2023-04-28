/*
Navicat MySQL Data Transfer

Source Server         : localhost liguanghong
Source Server Version : 80031
Source Host           : localhost:3306
Source Database       : exam

Target Server Type    : MYSQL
Target Server Version : 80031
File Encoding         : 65001

Date: 2023-04-28 16:57:51
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for answer
-- ----------------------------
DROP TABLE IF EXISTS `answer`;
CREATE TABLE `answer` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `plan_id` int unsigned NOT NULL DEFAULT '0' COMMENT '考试id',
  `paper_id` int unsigned NOT NULL DEFAULT '0' COMMENT '试卷id',
  `user_id` int unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `grade_id` int unsigned NOT NULL DEFAULT '0' COMMENT '成绩id',
  `submit_time` int unsigned NOT NULL DEFAULT '0' COMMENT '提交时间',
  `memo` varchar(255) NOT NULL DEFAULT '' COMMENT '备注信息',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='考生试卷答题表';

-- ----------------------------
-- Table structure for answer_item
-- ----------------------------
DROP TABLE IF EXISTS `answer_item`;
CREATE TABLE `answer_item` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `answer_id` int unsigned NOT NULL DEFAULT '0' COMMENT '答题表id',
  `question_id` int unsigned NOT NULL DEFAULT '0' COMMENT '试题id',
  `option_id` int unsigned NOT NULL DEFAULT '0' COMMENT '答案id',
  `content` text NOT NULL COMMENT '答案内容',
  `check` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '评分状态\r\n 0：未评分\r\n 1：已评分',
  `score` int unsigned NOT NULL DEFAULT '0' COMMENT '得分',
  `memo` varchar(255) NOT NULL DEFAULT '' COMMENT '备注信息',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='考生试题答案表';

-- ----------------------------
-- Table structure for class
-- ----------------------------
DROP TABLE IF EXISTS `class`;
CREATE TABLE `class` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '班级名称',
  `status` int unsigned NOT NULL DEFAULT '0' COMMENT '状态\r\n 0：禁用\r\n 1：正常',
  `desc` varchar(255) NOT NULL DEFAULT '' COMMENT '备注信息',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='班级表';

-- ----------------------------
-- Table structure for class_user_rel
-- ----------------------------
DROP TABLE IF EXISTS `class_user_rel`;
CREATE TABLE `class_user_rel` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `class_id` int unsigned NOT NULL DEFAULT '0' COMMENT '班级id',
  `user_id` int unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_user_role` (`user_id`,`class_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='班级考生表';

-- ----------------------------
-- Table structure for grade
-- ----------------------------
DROP TABLE IF EXISTS `grade`;
CREATE TABLE `grade` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `plan_id` int unsigned NOT NULL DEFAULT '0' COMMENT '考试id',
  `paper_id` int unsigned NOT NULL DEFAULT '0' COMMENT '试卷id',
  `user_id` int unsigned NOT NULL DEFAULT '0' COMMENT '考生id',
  `score` int unsigned NOT NULL DEFAULT '0' COMMENT '得分',
  `objective_score` int unsigned NOT NULL DEFAULT '0' COMMENT '客观题得分',
  `subjective_score` int unsigned NOT NULL DEFAULT '0' COMMENT '主观题得分',
  `status` int unsigned NOT NULL DEFAULT '0' COMMENT '状态\r\n 0：待参加考试\r\n 1：待交卷\r\n 2：已交卷待评分\r\n 3：部分评分\r\n 4：评分完成\r\n 5：考试取消',
  `start_time` int unsigned NOT NULL DEFAULT '0' COMMENT '开始考试时间',
  `end_time` int unsigned NOT NULL DEFAULT '0' COMMENT '结束考试时间',
  `duration` int unsigned NOT NULL DEFAULT '0' COMMENT '考试时长',
  `memo` varchar(255) NOT NULL DEFAULT '' COMMENT '备注信息',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='考试成绩表';

-- ----------------------------
-- Table structure for knowledge
-- ----------------------------
DROP TABLE IF EXISTS `knowledge`;
CREATE TABLE `knowledge` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '知识点名称',
  `desc` varchar(255) NOT NULL DEFAULT '' COMMENT '知识点描述',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='知识点表';

-- ----------------------------
-- Table structure for paper
-- ----------------------------
DROP TABLE IF EXISTS `paper`;
CREATE TABLE `paper` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '试卷名称',
  `subject_id` int unsigned NOT NULL DEFAULT '0' COMMENT '科目id',
  `knowledge_ids` varchar(2048) NOT NULL DEFAULT '' COMMENT '知识点id集合，多个id用”,“分隔',
  `score` int unsigned NOT NULL DEFAULT '0' COMMENT '试卷总分',
  `pass_score` int unsigned NOT NULL DEFAULT '0' COMMENT '及格分',
  `difficulty` decimal(8,6) unsigned NOT NULL DEFAULT '0.000000' COMMENT '难度',
  `choice_single_num` int unsigned NOT NULL DEFAULT '0' COMMENT '单选题数量',
  `choice_single_score` int unsigned NOT NULL DEFAULT '0' COMMENT '单选题分值',
  `choice_multi_num` int unsigned NOT NULL DEFAULT '0' COMMENT '多选题数量',
  `choice_multi_score` int unsigned NOT NULL DEFAULT '0' COMMENT '多选题分值',
  `judge_num` int unsigned NOT NULL DEFAULT '0' COMMENT '判断题数量',
  `judge_score` int unsigned NOT NULL DEFAULT '0' COMMENT '判断题分值',
  `blank_single_num` int unsigned NOT NULL DEFAULT '0' COMMENT '填空题数量',
  `blank_single_score` int unsigned NOT NULL DEFAULT '0' COMMENT '填空题分值',
  `blank_multi_num` int unsigned NOT NULL DEFAULT '0' COMMENT '多项填空题数量',
  `blank_multi_score` int unsigned NOT NULL DEFAULT '0' COMMENT '多项填空题分值',
  `answer_single_num` int unsigned NOT NULL DEFAULT '0' COMMENT '简答题数量',
  `answer_single_score` int unsigned NOT NULL DEFAULT '0' COMMENT '简答题分值',
  `answer_multi_num` int unsigned NOT NULL DEFAULT '0' COMMENT '多项简答题数量',
  `answer_multi_score` int unsigned NOT NULL DEFAULT '0' COMMENT '多项简答题分值',
  `status` int unsigned NOT NULL DEFAULT '0' COMMENT '状态\r\n 0：草稿\r\n 1：已发布',
  `create_time` int unsigned NOT NULL DEFAULT '0' COMMENT '添加时间',
  `update_time` int unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `memo` varchar(255) NOT NULL DEFAULT '' COMMENT '备注信息',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=63 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='试卷表';

-- ----------------------------
-- Table structure for paper_question
-- ----------------------------
DROP TABLE IF EXISTS `paper_question`;
CREATE TABLE `paper_question` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `paper_id` int unsigned NOT NULL DEFAULT '0' COMMENT '试卷id',
  `origin_id` int unsigned NOT NULL DEFAULT '0' COMMENT '原始试题id',
  `subject_id` int unsigned NOT NULL DEFAULT '0' COMMENT '科目id',
  `name` varchar(1024) NOT NULL DEFAULT '' COMMENT '题干',
  `type` int unsigned NOT NULL DEFAULT '1' COMMENT '题目类型\r\n 1：单选题\r\n 2：多选题\r\n 3：判断题\r\n 4：填空题\r\n 5：多项填空题\r\n 6：简答题\r\n 7：多项简答题\r\n 8：文件题\r\n 9：多项文件题',
  `content` text NOT NULL COMMENT '内容',
  `tips` text NOT NULL COMMENT '提示',
  `analysis` text NOT NULL COMMENT '解析',
  `difficulty` decimal(8,6) unsigned NOT NULL DEFAULT '0.000000' COMMENT '难度',
  `knowledge_ids` varchar(2048) NOT NULL DEFAULT '' COMMENT '知识点id集合，多个id用”,“分隔',
  `score` int unsigned NOT NULL DEFAULT '0' COMMENT '推荐分数',
  `update_time` int unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `memo` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `subject_id` (`subject_id`) USING BTREE,
  KEY `type` (`type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=994 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='试卷试题表';

-- ----------------------------
-- Table structure for paper_question_option
-- ----------------------------
DROP TABLE IF EXISTS `paper_question_option`;
CREATE TABLE `paper_question_option` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '试题id',
  `question_id` int unsigned NOT NULL DEFAULT '0' COMMENT '试题id',
  `tag` varchar(10) NOT NULL DEFAULT '' COMMENT '选项标签（A、B、C、...）',
  `content` text NOT NULL COMMENT '内容',
  `is_right` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否正确答案\r\n 0：否\r\n 1：是',
  `memo` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3307 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='试卷试题选项表';

-- ----------------------------
-- Table structure for permission
-- ----------------------------
DROP TABLE IF EXISTS `permission`;
CREATE TABLE `permission` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `type` int unsigned NOT NULL DEFAULT '1' COMMENT '权限类型\r\n 1：菜单权限\r\n 2：页面权限\r\n 3：组件权限\r\n 4：操作权限\r\n 5：按钮权限\r\n 6：数据权限',
  `pid` int unsigned NOT NULL DEFAULT '0' COMMENT '父权限id',
  `code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '权限代码',
  `status` int unsigned NOT NULL DEFAULT '0' COMMENT '状态\r\n 0：禁用\r\n 1：正常',
  `path` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '路由地址',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '路由名称',
  `component` varchar(255) NOT NULL DEFAULT '' COMMENT '路由组件',
  `redirect` varchar(1024) NOT NULL DEFAULT '' COMMENT '重定向',
  `meta_title` varchar(50) NOT NULL DEFAULT '' COMMENT '名称',
  `meta_icon` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '图标',
  `meta_extra_icon` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '右侧图标',
  `meta_show_link` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否显示\r\n 0：隐藏\r\n 1：显示',
  `meta_show_parent` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否显示父级菜单\r\n 0：隐藏\r\n 1：显示',
  `meta_keep_alive` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '缓存\r\n 0：关闭\r\n 1：开启',
  `meta_frame_src` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '内嵌的`iframe`链接',
  `meta_frame_loading` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '`iframe`页是否开启首次加载动画\r\n 0：关闭\r\n 1：开启',
  `meta_hidden_tag` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否不添加信息到标签页\r\n 0：添加\r\n 1：不添加',
  `meta_rank` int unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `create_time` int unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `memo` varchar(255) NOT NULL DEFAULT '' COMMENT '备注信息',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='权限表';

-- ----------------------------
-- Table structure for plan
-- ----------------------------
DROP TABLE IF EXISTS `plan`;
CREATE TABLE `plan` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '考试名称',
  `paper_id` int unsigned NOT NULL DEFAULT '0' COMMENT '试卷id',
  `start_time` int unsigned NOT NULL DEFAULT '0' COMMENT '开始考试试卷',
  `end_time` int unsigned NOT NULL DEFAULT '0' COMMENT '结束考试时间',
  `duration` int unsigned NOT NULL DEFAULT '0' COMMENT '考试时长',
  `publish_time` int unsigned NOT NULL DEFAULT '0' COMMENT '发布时间',
  `status` int unsigned NOT NULL DEFAULT '0' COMMENT '状态\r\n 0：待发布\r\n 1：已发布\r\n 2：已取消\r\n 3：已结束',
  `query_grade` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '考试是否可查询成绩\r\n 0：否\r\n 1：是',
  `create_time` int unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `memo` varchar(255) NOT NULL DEFAULT '' COMMENT '备注信息',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='考试计划表';

-- ----------------------------
-- Table structure for plan_class_rel
-- ----------------------------
DROP TABLE IF EXISTS `plan_class_rel`;
CREATE TABLE `plan_class_rel` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `plan_id` int unsigned NOT NULL DEFAULT '0' COMMENT '考试计划id',
  `class_id` int unsigned NOT NULL DEFAULT '0' COMMENT '班级id',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_plan_class` (`plan_id`,`class_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='班级考生表';

-- ----------------------------
-- Table structure for question
-- ----------------------------
DROP TABLE IF EXISTS `question`;
CREATE TABLE `question` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `subject_id` int unsigned NOT NULL DEFAULT '0' COMMENT '科目id',
  `name` varchar(1024) NOT NULL DEFAULT '' COMMENT '题干',
  `type` int unsigned NOT NULL DEFAULT '1' COMMENT '题目类型\r\n 1：单选题\r\n 2：多选题\r\n 3：判断题\r\n 4：填空题\r\n 5：多项填空题\r\n 6：简答题\r\n 7：多项简答题\r\n 8：文件题\r\n 9：多项文件题',
  `content` text NOT NULL COMMENT '内容',
  `tips` text NOT NULL COMMENT '提示',
  `analysis` text NOT NULL COMMENT '解析',
  `difficulty` decimal(8,6) unsigned NOT NULL DEFAULT '0.000000' COMMENT '难度',
  `knowledge_ids` varchar(2048) NOT NULL DEFAULT '' COMMENT '知识点id集合，多个id用”,“分隔',
  `score` int unsigned NOT NULL DEFAULT '0' COMMENT '推荐分数',
  `status` int unsigned NOT NULL DEFAULT '0' COMMENT '状态\r\n 0：禁用\r\n 1：正常',
  `create_time` int unsigned NOT NULL DEFAULT '0' COMMENT '添加时间',
  `update_time` int unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `memo` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `subject_id` (`subject_id`) USING BTREE,
  KEY `type` (`type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2915 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='试题表';

-- ----------------------------
-- Table structure for question_option
-- ----------------------------
DROP TABLE IF EXISTS `question_option`;
CREATE TABLE `question_option` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '试题id',
  `question_id` int unsigned NOT NULL DEFAULT '0' COMMENT '试题id',
  `tag` varchar(10) NOT NULL DEFAULT '' COMMENT '选项标签（A、B、C、...）',
  `content` text NOT NULL COMMENT '内容',
  `is_right` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否正确答案\r\n 0：否\r\n 1：是',
  `memo` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11325 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='试题选项表';

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '角色名称',
  `code` varchar(50) NOT NULL DEFAULT '' COMMENT '角色代码',
  `seq` int unsigned NOT NULL DEFAULT '0' COMMENT '序号',
  `status` int unsigned NOT NULL DEFAULT '0' COMMENT '状态\r\n 0：禁用\r\n 1：正常',
  `create_time` int unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `memo` varchar(255) NOT NULL DEFAULT '' COMMENT '备注信息',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_code` (`code`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色表';

-- ----------------------------
-- Table structure for role_permission_rel
-- ----------------------------
DROP TABLE IF EXISTS `role_permission_rel`;
CREATE TABLE `role_permission_rel` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int unsigned NOT NULL DEFAULT '0' COMMENT '角色id',
  `permission_id` int unsigned NOT NULL DEFAULT '0' COMMENT '权限id',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_role_permission` (`role_id`,`permission_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色权限表';

-- ----------------------------
-- Table structure for subject
-- ----------------------------
DROP TABLE IF EXISTS `subject`;
CREATE TABLE `subject` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '科目名称',
  `desc` varchar(255) NOT NULL DEFAULT '' COMMENT '科目描述',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='科目表';

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(50) NOT NULL DEFAULT '' COMMENT '密码',
  `type` int unsigned NOT NULL DEFAULT '0' COMMENT '账号类型\r\n 1：管理员账号\r\n 2：教师账号\r\n 3：考生账号',
  `true_name` varchar(50) NOT NULL DEFAULT '' COMMENT '真实姓名',
  `mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',
  `email` varchar(255) NOT NULL DEFAULT '' COMMENT '邮箱地址',
  `status` int unsigned NOT NULL DEFAULT '0' COMMENT '状态\r\n 0：禁用\r\n 1：正常',
  `create_time` int unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `memo` varchar(255) NOT NULL DEFAULT '' COMMENT '备注信息',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_name` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户表';

-- ----------------------------
-- Table structure for user_role_rel
-- ----------------------------
DROP TABLE IF EXISTS `user_role_rel`;
CREATE TABLE `user_role_rel` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `role_id` int unsigned NOT NULL DEFAULT '0' COMMENT '角色id',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_user_role` (`user_id`,`role_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户角色表';
