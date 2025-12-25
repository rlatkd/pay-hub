

-- 통합 결제 트랜잭션
CREATE TABLE core_transaction (
    tx_id VARCHAR(255) PRIMARY KEY,
    external_ref_id VARCHAR(255) NOT NULL,
    method_type VARCHAR(20) NOT NULL,
    amount DECIMAL(18, 2) NOT NULL,
    currency VARCHAR(10) DEFAULT 'KRW',
    status VARCHAR(50) NOT NULL,
    created_at DATETIME(6) DEFAULT CURRENT_TIMESTAMP(6),
    updated_at DATETIME(6) DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6)
);

-- PG 승인 상세 정보
CREATE TABLE core_approval (
    approval_id VARCHAR(255) PRIMARY KEY,
    tx_id VARCHAR(255) UNIQUE NOT NULL,
    approval_provider VARCHAR(50) NOT NULL,
    approve_no VARCHAR(100),
    card_company VARCHAR(50),
    raw_response JSON,
    approved_at DATETIME(6),
    CONSTRAINT fk_core_pg_tx FOREIGN KEY (tx_id) REFERENCES core_transaction(tx_id)
);


-- 취소 및 환불 기록
CREATE TABLE core_refund (
    refund_id VARCHAR(255) PRIMARY KEY,
    tx_id VARCHAR(255) NOT NULL,
    refund_pg_tid VARCHAR(255),
    amount DECIMAL(18, 2) NOT NULL,
    reason TEXT,
    status VARCHAR(50) NOT NULL,
    refunded_at DATETIME(6) DEFAULT CURRENT_TIMESTAMP(6),
    CONSTRAINT fk_core_refund_tx FOREIGN KEY (tx_id) REFERENCES core_transaction(tx_id)
);

-- CREATE TABLE core_payment_failure (
--     failure_id BIGINT AUTO_INCREMENT PRIMARY KEY,
--     tx_id VARCHAR(255) NOT NULL,
--     error_code VARCHAR(100),
--     error_message TEXT,
--     failed_at DATETIME(6) DEFAULT CURRENT_TIMESTAMP(6),
--     CONSTRAINT fk_core_fail_tx FOREIGN KEY (tx_id) REFERENCES core_transaction(tx_id)
-- );
