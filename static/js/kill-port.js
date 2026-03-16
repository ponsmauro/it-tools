async function scanAndKill() {
            const port = document.getElementById('port').value;
            const result = document.getElementById('result');
            
            if (!port || port < 1 || port > 65535) {
                result.textContent = 'Invalid port';
                result.className = 'result-box error';
                return;
            }
            
            result.textContent = `Scanning port ${port}...`;
            result.className = 'result-box loading';
            
            // Mock delay
            await new Promise(resolve => setTimeout(resolve, 1000));
            
            const pid1 = Math.floor(Math.random() * 90000) + 10000;
            const pid2 = Math.floor(Math.random() * 90000) + 10000;
            
            result.innerHTML = `
                <h3>Processes:</h3>
                <pre>lsof -ti:${port}
COMMAND  PID USER  FD TYPE DEVICE SIZE/OFF NODE NAME
node    ${pid1} mpons 20u IPv6 0x123 0t0 TCP *:${port} (LISTEN)
python  ${pid2} mpons 15u IPv4 0x456 0t0 TCP *:${port} (LISTEN)</pre>
                <p>Killed PIDs ${pid1}, ${pid2}</p>
                <strong>✅ Freed!</strong>
            `;
            result.className = 'result-box success';
        }

        document.addEventListener('DOMContentLoaded', function() {
            const tabs = document.querySelectorAll('.os-tab');
            const panels = document.querySelectorAll('.os-panel');

            tabs.forEach(tab => {
                tab.addEventListener('click', function() {
                    const os = tab.getAttribute('data-os');

                    tabs.forEach(item => item.classList.remove('active'));
                    panels.forEach(item => item.classList.remove('active'));

                    tab.classList.add('active');
                    document.getElementById(`os-${os}`).classList.add('active');
                });
            });
        });