-- ============================================
-- migrations/000006_create_security_tables.down.sql
-- ============================================

DROP TABLE IF EXISTS audit_logs;
DROP TABLE IF EXISTS login_attempts;