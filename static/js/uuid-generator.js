let uuids = [];
        
        function updateCountValue() {
            const countValue = document.getElementById('count-slider').value;
            document.getElementById('count-value').textContent = countValue;
        }
        
        function formatUUID(uuid) {
            const uppercase = document.getElementById('uppercase').checked;
            const hyphens = document.getElementById('hyphens').checked;
            const braces = document.getElementById('braces').checked;
            
            // Start with the raw UUID
            let formattedUUID = uuid;
            
            // First handle case (uppercase/lowercase)
            if (uppercase) {
                formattedUUID = formattedUUID.toUpperCase();
            } else {
                formattedUUID = formattedUUID.toLowerCase();
            }
            
            // Then handle hyphens (remove if not wanted)
            if (!hyphens) {
                formattedUUID = formattedUUID.replace(/-/g, '');
            }
            
            // Finally add braces if wanted
            if (braces) {
                formattedUUID = '{' + formattedUUID + '}';
            }
            
            return formattedUUID;
        }
        
        function generate() {
            // Ensure count is between 1 and 50
            const count = Math.max(1, Math.min(50, parseInt(document.getElementById('count-slider').value) || 1));
            
            // Update slider value if it was out of bounds
            document.getElementById('count-slider').value = count;
            document.getElementById('count-value').textContent = count;
            
            uuids = [];
            
            for (let i = 0; i < count; i++) {
                const rawUUID = crypto.randomUUID();
                const formattedUUID = formatUUID(rawUUID);
                uuids.unshift(formattedUUID);
            }
            
            renderList();
            updateStats();
            document.getElementById('copy-btn').disabled = false;
        }
        
        function renderList() {
            const container = document.getElementById('uuid-list');
            container.innerHTML = uuids.map((uuid, index) => 
                `<div class="uuid-item" onclick="copyUUID(${index}, this)" title="Click to copy">
                    <code>${uuid}</code>
                </div>`
            ).join('');
        }
        
        function copyUUID(index, element) {
            navigator.clipboard.writeText(uuids[index]).then(() => {
                element.classList.add('state-copied');
                element.title = 'Copied!';
                setTimeout(() => {
                    element.classList.remove('state-copied');
                    element.title = 'Click to copy';
                }, 1500);
            }).catch(err => {
                console.error('Failed to copy: ', err);
            });
        }
        
        function copyAll() {
            const text = uuids.join('\n');
            navigator.clipboard.writeText(text).then(() => {
                const btn = document.getElementById('copy-btn');
                const original = btn.textContent;
                btn.textContent = '✅ Copied!';
                btn.classList.add('state-copied');
                setTimeout(() => {
                    btn.textContent = original;
                    btn.classList.remove('state-copied');
                }, 1500);
            }).catch(err => {
                console.error('Failed to copy all: ', err);
            });
        }
        
        function clearList() {
            uuids = [];
            document.getElementById('uuid-list').innerHTML = '';
            document.getElementById('stats').style.display = 'none';
            document.getElementById('copy-btn').disabled = true;
        }
        
        function updateStats() {
            document.getElementById('count').textContent = uuids.length;
            document.getElementById('stats').style.display = 'block';
        }