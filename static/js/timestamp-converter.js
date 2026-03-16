// Utility to copy text
        async function copyText(elementId) {
            const el = document.getElementById(elementId);
            const text = el.tagName === 'INPUT' ? el.value : el.textContent;
            if (!text) return;
            try {
                await navigator.clipboard.writeText(text);
                // Optional: show a brief tooltip or toast
            } catch (err) {
                console.error('Failed to copy', err);
            }
        }

        // Current Timestamp
        function updateCurrentTimestamp() {
            const now = Date.now();
            document.getElementById('current-ts-ms').textContent = now;
            document.getElementById('current-ts-sec').textContent = Math.floor(now / 1000);
        }

        // Timestamp to Date
        function convertTsToDate() {
            const input = document.getElementById('ts-input').value.trim();
            if (!input) return;

            let ts = parseInt(input, 10);
            if (isNaN(ts)) return;

            // Auto-detect seconds vs milliseconds
            // If it's less than 10000000000 (Nov 2286), assume seconds
            if (ts < 10000000000) {
                ts *= 1000;
            }

            const date = new Date(ts);
            
            if (isNaN(date.getTime())) {
                document.getElementById('ts-out-local').value = 'Invalid Date';
                document.getElementById('ts-out-utc').value = 'Invalid Date';
                document.getElementById('ts-out-iso').value = 'Invalid Date';
                return;
            }

            document.getElementById('ts-out-local').value = date.toString();
            document.getElementById('ts-out-utc').value = date.toUTCString();
            document.getElementById('ts-out-iso').value = date.toISOString();
        }

        // Date to Timestamp
        function convertDateToTs() {
            const input = document.getElementById('date-input').value.trim();
            if (!input) return;

            const date = new Date(input);
            
            if (isNaN(date.getTime())) {
                document.getElementById('date-out-sec').value = 'Invalid Date';
                document.getElementById('date-out-ms').value = 'Invalid Date';
                return;
            }

            const ms = date.getTime();
            document.getElementById('date-out-ms').value = ms;
            document.getElementById('date-out-sec').value = Math.floor(ms / 1000);
        }

        // Event Listeners
        document.getElementById('btn-ts-to-date').addEventListener('click', convertTsToDate);
        document.getElementById('ts-input').addEventListener('keypress', (e) => {
            if (e.key === 'Enter') convertTsToDate();
        });
        document.getElementById('btn-clear-ts').addEventListener('click', () => {
            document.getElementById('ts-input').value = '';
            document.getElementById('ts-out-local').value = '';
            document.getElementById('ts-out-utc').value = '';
            document.getElementById('ts-out-iso').value = '';
        });

        document.getElementById('btn-date-to-ts').addEventListener('click', convertDateToTs);
        document.getElementById('date-input').addEventListener('keypress', (e) => {
            if (e.key === 'Enter') convertDateToTs();
        });
        document.getElementById('btn-clear-date').addEventListener('click', () => {
            document.getElementById('date-input').value = '';
            document.getElementById('date-out-sec').value = '';
            document.getElementById('date-out-ms').value = '';
        });

        // Initialize — store interval ID so it can be cleared if the tool is unloaded via AJAX
        updateCurrentTimestamp();
        const _tsIntervalId = setInterval(updateCurrentTimestamp, 1000);

        // Clean up interval when the tool content is replaced by AJAX navigation
        document.addEventListener('toolunload', function onToolUnload() {
            clearInterval(_tsIntervalId);
            document.removeEventListener('toolunload', onToolUnload);
        }, { once: true });