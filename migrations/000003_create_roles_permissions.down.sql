-- ============================================
-- migrations/000003_create_roles_permissions.down.sql
-- ============================================

DROP TABLE IF EXISTS role_permissions;
DROP TABLE IF EXISTS user_roles;
DROP TABLE IF EXISTS permissions;
DROP TABLE IF EXISTS roles;