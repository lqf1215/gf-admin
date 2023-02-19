/*
 Navicat Premium Data Transfer

 Source Server         : localhost-pgsql
 Source Server Type    : PostgreSQL
 Source Server Version : 150002 (150002)
 Source Host           : localhost:5432
 Source Catalog        : gf-school
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 150002 (150002)
 File Encoding         : 65001

 Date: 19/02/2023 06:07:48
*/


-- ----------------------------
-- Table structure for co_company_dept
-- ----------------------------
DROP TABLE IF EXISTS "public"."co_company_dept";
CREATE TABLE "public"."co_company_dept" (
  "id" int8 NOT NULL,
  "parent_id" int8,
  "ancestors" varchar(45) COLLATE "pg_catalog"."default",
  "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "email" varchar(64) COLLATE "pg_catalog"."default",
  "state" int4,
  "sort" int4,
  "mobile" varchar(16) COLLATE "pg_catalog"."default",
  "leader" varchar(64) COLLATE "pg_catalog"."default",
  "created_by" varchar(64) COLLATE "pg_catalog"."default",
  "updated_by" varchar(64) COLLATE "pg_catalog"."default",
  "created_at" timestamp(6),
  "updated_at" timestamp(6),
  "deleted_at" timestamp(6)
)
;
COMMENT ON COLUMN "public"."co_company_dept"."id" IS 'ID';
COMMENT ON COLUMN "public"."co_company_dept"."parent_id" IS '父部门id';
COMMENT ON COLUMN "public"."co_company_dept"."ancestors" IS '祖级列表';
COMMENT ON COLUMN "public"."co_company_dept"."name" IS '部门名称';
COMMENT ON COLUMN "public"."co_company_dept"."email" IS '邮箱';
COMMENT ON COLUMN "public"."co_company_dept"."state" IS '部门状态（0正常 1停用）';
COMMENT ON COLUMN "public"."co_company_dept"."sort" IS '显示顺序';
COMMENT ON COLUMN "public"."co_company_dept"."mobile" IS '联系电话';
COMMENT ON COLUMN "public"."co_company_dept"."leader" IS '负责人';
COMMENT ON COLUMN "public"."co_company_dept"."created_by" IS '创建人';
COMMENT ON COLUMN "public"."co_company_dept"."updated_by" IS '修改人';

-- ----------------------------
-- Table structure for co_company_post
-- ----------------------------
DROP TABLE IF EXISTS "public"."co_company_post";
CREATE TABLE "public"."co_company_post" (
  "id" int8 NOT NULL,
  "code" varchar(45) COLLATE "pg_catalog"."default" NOT NULL,
  "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "state" int4,
  "sort" int4,
  "remark" varchar(64) COLLATE "pg_catalog"."default",
  "created_by" varchar(64) COLLATE "pg_catalog"."default",
  "updated_by" varchar(64) COLLATE "pg_catalog"."default",
  "created_at" timestamp(6),
  "updated_at" timestamp(6),
  "deleted_at" timestamp(6)
)
;
COMMENT ON COLUMN "public"."co_company_post"."id" IS 'ID';
COMMENT ON COLUMN "public"."co_company_post"."code" IS '岗位编码';
COMMENT ON COLUMN "public"."co_company_post"."name" IS '岗位名称';
COMMENT ON COLUMN "public"."co_company_post"."state" IS '状态（0正常 1停用）';
COMMENT ON COLUMN "public"."co_company_post"."sort" IS '显示顺序';
COMMENT ON COLUMN "public"."co_company_post"."remark" IS '备注';
COMMENT ON COLUMN "public"."co_company_post"."created_by" IS '创建人';
COMMENT ON COLUMN "public"."co_company_post"."updated_by" IS '修改人';
