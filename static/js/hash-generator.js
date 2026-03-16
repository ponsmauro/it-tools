// Utility to copy text
        async function copyText(elementId) {
            const el = document.getElementById(elementId);
            if (!el.value) return;
            try {
                await navigator.clipboard.writeText(el.value);
            } catch (err) {
                console.error('Failed to copy', err);
            }
        }

        // Convert ArrayBuffer to Hex String
        function bufferToHex(buffer) {
            const hashArray = Array.from(new Uint8Array(buffer));
            return hashArray.map(b => b.toString(16).padStart(2, '0')).join('');
        }

        // Web Crypto API for SHA hashes
        async function generateSHA(text, algorithm) {
            const msgUint8 = new TextEncoder().encode(text);
            const hashBuffer = await crypto.subtle.digest(algorithm, msgUint8);
            return bufferToHex(hashBuffer);
        }

        // Update hashes
        async function updateHashes() {
            const input = document.getElementById('hash-input').value;
            const isUppercase = document.getElementById('uppercase-output').checked;

            if (!input) {
                document.getElementById('out-sha1').value = '';
                document.getElementById('out-sha256').value = '';
                document.getElementById('out-sha384').value = '';
                document.getElementById('out-sha512').value = '';
                return;
            }

            try {
                const [sha1, sha256, sha384, sha512] = await Promise.all([
                    generateSHA(input, 'SHA-1'),
                    generateSHA(input, 'SHA-256'),
                    generateSHA(input, 'SHA-384'),
                    generateSHA(input, 'SHA-512')
                ]);

                const formatOutput = (hash) => isUppercase ? hash.toUpperCase() : hash;

                document.getElementById('out-sha1').value = formatOutput(sha1);
                document.getElementById('out-sha256').value = formatOutput(sha256);
                document.getElementById('out-sha384').value = formatOutput(sha384);
                document.getElementById('out-sha512').value = formatOutput(sha512);
            } catch (err) {
                console.error('Error generating hashes', err);
            }
        }

        // Event Listeners
        document.getElementById('hash-input').addEventListener('input', updateHashes);
        document.getElementById('uppercase-output').addEventListener('change', updateHashes);
        
        document.getElementById('btn-clear').addEventListener('click', () => {
            document.getElementById('hash-input').value = '';
            updateHashes();
        });