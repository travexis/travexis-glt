package main

import (
    "crypto/sha256"
    "encoding/hex"
    "encoding/json"
    "fmt"
    "os"
    "strings"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: calc_proof.exe <input.json>")
        os.Exit(1)
    }
    
    data, _ := os.ReadFile(os.Args[1])
    var m map[string]any
    json.Unmarshal(data, &m)
    
    payload := m["payload"].(map[string]any)
    
    canonData := map[string]any{
        "case_id": payload["case_id"],
        "profile": payload["profile"],
        "action":  payload["action"],
        "payload": payload["payload"],
    }
    
    canon, _ := json.Marshal(canonData)
    h := sha256.Sum256([]byte(strings.TrimSpace(string(canon))))
    proof := "sha256:" + hex.EncodeToString(h[:])
    
    fmt.Println(proof)
}
