const inputEl = document.getElementById('json-input');
        const outputEl = document.getElementById('json-output');
        const errorEl = document.getElementById('error-message');
        const indentSelect = document.getElementById('indent-select');

        function getIndent() {
            const val = indentSelect.value;
            if (val === 'tab') return '\t';
            return parseInt(val, 10);
        }

        function processJson(minify = false) {
            const raw = inputEl.value.trim();
            if (!raw) {
                outputEl.value = '';
                errorEl.style.display = 'none';
                return;
            }

            try {
                const parsed = JSON.parse(raw);
                const indent = minify ? 0 : getIndent();
                outputEl.value = JSON.stringify(parsed, null, indent);
                errorEl.style.display = 'none';
            } catch (err) {
                outputEl.value = '';
                errorEl.textContent = '{{ tr "tools.json-formatter.invalid" }}: ' + err.message;
                errorEl.style.display = 'block';
            }
        }

        document.getElementById('btn-format').addEventListener('click', () => processJson(false));
        document.getElementById('btn-minify').addEventListener('click', () => processJson(true));
        
        document.getElementById('btn-clear').addEventListener('click', () => {
            inputEl.value = '';
            outputEl.value = '';
            errorEl.style.display = 'none';
        });

        document.getElementById('btn-copy').addEventListener('click', async () => {
            if (!outputEl.value) return;
            try {
                await navigator.clipboard.writeText(outputEl.value);
            } catch (err) {
                console.error('Failed to copy', err);
            }
        });