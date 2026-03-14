package killport

import (
	"encoding/json"
	toolregistry "it-tools/internal/tool_registry"
	"log"
	"net/http"
)

func KillPortHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Action string `json:"action"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("JSON decode error: %v", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if req.Action != "open" {
		http.Error(w, "Action must be 'open'", http.StatusBadRequest)
		return
	}

	// Simple HTML fragment (no template parse issues)
	const content = `
<div class="tool-content" style="padding: 20px;">
	<h2>Kill Port</h2>
	<p>Enter port to see process and mock kill.</p>
	<input type="number" id="port" placeholder="8080" style="padding: 8px; width: 150px;">
	<button onclick="simulateKill()" style="padding: 8px 16px; background: #ef4444; color: white; border: none;">Kill</button>
	<div id="result" style="margin-top: 20px; padding: 15px; background: #f5f5f5; font-family: monospace;"></div>
</div>
<script>
function simulateKill() {
	const port = document.getElementById('port').value;
	const result = document.getElementById('result');
	if (!port) {
		result.innerHTML = 'Enter port';
		return;
	}
	
	const pid = Math.floor(Math.random()*90000)+10000;
	result.innerHTML = 'Process: node PID ' + pid + '\\nUser: mpons\\nCommand: node server.js\\n\\nkill -9 ' + pid + ' (mock)\\n✅ Port ' + port + ' killed!';
}
</script>
	`
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(content))
}

func Init() {
	toolregistry.Register("kill-port", KillPortHandler, "")
}
