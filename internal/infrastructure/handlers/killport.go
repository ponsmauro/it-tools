package handlers

// KillPortInitialContent returns the initial kill-port UI
const KillPortInitialContent = `
<div style="padding: 40px; max-width: 800px; margin: 0 auto;">
    <h1>Kill Port Tool</h1>
    <p>1. Scan port. 2. Kill process with confirmation.</p>
    <div style="margin: 20px 0;">
        <label>Port: </label>
        <input type="number" id="port-input" placeholder="8080" min="1" max="65535" style="padding: 8px 12px; width: 120px; margin-left: 10px;">
        <button onclick="scanPort()" style="padding: 8px 16px; background: #059669; color: white; border: none; border-radius: 4px; margin-left: 10px; cursor: pointer;">Scan</button>
    </div>
    <div id="scan-result" style="display: none; background: #f8fafc; padding: 20px; border-radius: 8px; font-family: monospace; border: 1px solid #e2e8f0;"></div>
    <div id="process-list" style="display: none; background: #f8fafc; padding: 20px; border-radius: 8px; margin-top: 10px;"></div>
</div>

<div id="confirm-modal" style="display: none; position: fixed; top: 0; left: 0; width: 100%; height: 100%; background: rgba(0,0,0,0.5); z-index: 1000; justify-content: center; align-items: center;" onclick="closeModal()">
    <div style="background: white; padding: 30px; border-radius: 12px; max-width: 400px; text-align: center;" onclick="event.stopPropagation()">
        <h2>Confirm Kill</h2>
        <p id="modal-message"></p>
        <button id="confirm-kill" style="padding: 10px 20px; background: #dc2626; color: white; border: none; border-radius: 6px; margin: 10px; cursor: pointer;">Kill</button>
        <button onclick="closeModal()" style="padding: 10px 20px; background: #6b7280; color: white; border: none; border-radius: 6px; margin: 10px; cursor: pointer;">Cancel</button>
    </div>
</div>

<script>
let currentPort = '';
function scanPort() {
    currentPort = document.getElementById('port-input').value;
    const result = document.getElementById('scan-result');
    const processList = document.getElementById('process-list');
    if (!currentPort || currentPort < 1 || currentPort > 65535) {
        result.innerHTML = 'Enter valid port (1-65535)';
        result.style.display = 'block';
        processList.style.display = 'none';
        return;
    }
    result.innerHTML = 'Scan port ' + currentPort + '\\nRunning mock lsof -ti:' + currentPort + '\\n\\nFound processes:';
    result.style.display = 'block';
    processList.innerHTML = '<div style="padding: 12px; border-bottom: 1px solid #e2e8f0; cursor: pointer;" onclick="showConfirm(12345, \'node Express server\')">&bull; node PID: 12345 User: mpons</div>' + '<div style="padding: 12px; border-bottom: 1px solid #e2e8f0; cursor: pointer;" onclick="showConfirm(67890, \'python Flask app\')">&bull; python PID: 67890 User: mpons</div>';
    processList.style.display = 'block';
}
function showConfirm(pid, command) {
    document.getElementById('modal-message').innerHTML = 'Kill <strong>' + command + '</strong> (PID: ' + pid + ')?';
    document.getElementById('confirm-kill').onclick = function() { confirmKill(pid); };
    document.getElementById('confirm-modal').style.display = 'flex';
}
function confirmKill(pid) {
    document.getElementById('scan-result').innerHTML += '\\n\\nKilled PID ' + pid + '! Port ' + currentPort + ' freed.';
    document.getElementById('process-list').style.display = 'none';
    closeModal();
}
function closeModal() {
    document.getElementById('confirm-modal').style.display = 'none';
}
</script>
`
