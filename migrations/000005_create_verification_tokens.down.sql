-- ============================================
-- migrations/000005_create_verification_tokens.down.sql
-- ============================================

DROP TABLE IF EXISTS password_reset_tokens;
DROP TABLE IF EXISTS email_verification_tokens;
