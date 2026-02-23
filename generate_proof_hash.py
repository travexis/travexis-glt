import hashlib
import json

# Payload 数据
payload = {
    "ledger_head": { "ledger_seq": 41, "line_hash": "abc123" },
    "ledger_seq": 42,
    "prev_hash": "abc123",
    "entry_type": "append",
    "amount_usd": 1200.5,
    "cfu_impact": 54.4,
    "responsible_party": "operator",
    "evidence_ref": "PP-20260219-173252Z"
}

# 规范化 payload 为 JSON 字符串
payload_str = json.dumps(payload, separators=(',', ':'))

# 计算 SHA256 哈希值
proof_hash = hashlib.sha256(payload_str.encode('utf-8')).hexdigest()

# 输出 proof_hash
print("proof_hash:", proof_hash)