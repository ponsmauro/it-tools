const input = document.getElementById('input-text');
        const output = document.getElementById('output-text');
        
        input.addEventListener('input', updateLive);
        
        function updateLive() {
            updateStats();
            output.value = input.value;
        }
        
        function updateStats() {
            const text = input.value;
            const lines = text.split('\n');
            const words = text.trim().split(/\s+/).filter(w => w.length > 0);
            
            document.getElementById('chars').textContent = text.length;
            document.getElementById('words').textContent = words.length;
            document.getElementById('lines').textContent = lines.length;
            document.getElementById('empty-lines').textContent = lines.filter(l => l.trim() === '').length;
        }
        
        function reverseLines() {
            const lines = input.value.split('\n').reverse().join('\n');
            input.value = lines;
            updateLive();
        }
        
        function sortLines() {
            const lines = input.value.split('\n').sort().join('\n');
            input.value = lines;
            updateLive();
        }
        
        function uniqueLines() {
            const lines = input.value.split('\n').filter((line, index, self) => 
                self.findIndex(l => l === line) === index
            ).join('\n');
            input.value = lines;
            updateLive();
        }
        
        function clearText() {
            input.value = '';
            output.value = '';
            updateStats();
        }
        
        function transformCase(type) {
            let result = input.value;
            
            switch(type) {
                case 'upper':
                    result = result.toUpperCase();
                    break;
                case 'lower':
                    result = result.toLowerCase();
                    break;
                case 'title':
                    result = result.replace(/\w\S*/g, txt => 
                        txt.charAt(0).toUpperCase() + txt.substr(1).toLowerCase()
                    );
                    break;
                case 'sentence':
                    result = result.toLowerCase().replace(/(^\s*\w|[.!?]\s*\w)/g, c => 
                        c.toUpperCase()
                    );
                    break;
            }
            
            input.value = result;
            updateLive();
        }
        
        function trimWhitespace(type) {
            let result = input.value;
            
            switch(type) {
                case 'start':
                    result = result.replace(/^\s+/gm, '');
                    break;
                case 'end':
                    result = result.replace(/\s+$/gm, '');
                    break;
                case 'both':
                    result = result.replace(/^\s+|\s+$/gm, '');
                    break;
            }
            
            input.value = result;
            updateLive();
        }
        
        function findAndReplace() {
            const findText = document.getElementById('find-text').value;
            const replaceText = document.getElementById('replace-text').value;
            const caseSensitive = document.getElementById('case-sensitive').checked;
            
            if (!findText) return;
            
            let result = input.value;
            let flags = 'g';
            if (!caseSensitive) flags += 'i';
            
            const regex = new RegExp(escapeRegExp(findText), flags);
            result = result.replace(regex, replaceText);
            
            input.value = result;
            updateLive();
        }
        
        function escapeRegExp(string) {
            return string.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
        }
        
        function addLineNumbers() {
            const lines = input.value.split('\n');
            const numberedLines = lines.map((line, index) => 
                `${index + 1}. ${line}`
            );
            
            input.value = numberedLines.join('\n');
            updateLive();
        }
        
        function copyOutput() {
            try {
                navigator.clipboard.writeText(output.value).then(() => {
                    const btn = event.target;
                    const original = btn.textContent;
                    btn.textContent = '✅ Copied!';
                    btn.style.background = '#10b981';
                    setTimeout(() => {
                        btn.textContent = original;
                        btn.style.background = '';
                    }, 1500);
                }).catch(err => {
                    console.error('Failed to copy: ', err);
                    copyToClipboardFallback(output.value);
                });
            } catch (err) {
                console.error('Error accessing clipboard: ', err);
                copyToClipboardFallback(output.value);
            }
        }
        
        function copyToClipboardFallback(text) {
            const textArea = document.createElement('textarea');
            textArea.value = text;
            textArea.style.position = 'fixed';
            document.body.appendChild(textArea);
            textArea.focus();
            textArea.select();
            
            try {
                const successful = document.execCommand('copy');
                const msg = successful ? 'Copied!' : 'Failed to copy';
                console.log(msg);
            } catch (err) {
                console.error('Fallback: Could not copy text: ', err);
            }
            
            document.body.removeChild(textArea);
        }