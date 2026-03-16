(function() {
            const input = document.getElementById('jwt-input');
            const headerOutput = document.getElementById('jwt-header-output');
            const payloadOutput = document.getElementById('jwt-payload-output');
            const signatureOutput = document.getElementById('jwt-signature-output');
            const errorMsg = document.getElementById('jwt-error');
            const clearBtn = document.getElementById('btn-clear');

            function base64UrlDecode(str) {
                // Normalize base64url to base64 and add padding
                str = str.replace(/-/g, '+').replace(/_/g, '/');
                while (str.length % 4) {
                    str += '=';
                }

                // Decode base64 to bytes and convert to UTF-8 string via TextDecoder
                const bytes = Uint8Array.from(atob(str), c => c.charCodeAt(0));
                return new TextDecoder().decode(bytes);
            }

            function formatJSON(str) {
                try {
                    const obj = JSON.parse(str);
                    return JSON.stringify(obj, null, 2);
                } catch (e) {
                    return str;
                }
            }

            function decodeJWT() {
                const token = input.value.trim();
                
                if (!token) {
                    headerOutput.textContent = '';
                    payloadOutput.textContent = '';
                    signatureOutput.textContent = '';
                    errorMsg.style.display = 'none';
                    return;
                }

                const parts = token.split('.');
                
                if (parts.length !== 3) {
                    headerOutput.textContent = '';
                    payloadOutput.textContent = '';
                    signatureOutput.textContent = '';
                    errorMsg.style.display = 'block';
                    return;
                }

                try {
                    const header = formatJSON(base64UrlDecode(parts[0]));
                    const payload = formatJSON(base64UrlDecode(parts[1]));
                    const signature = parts[2]; // Signature is not decoded, just displayed

                    headerOutput.textContent = header;
                    payloadOutput.textContent = payload;
                    signatureOutput.textContent = signature;
                    errorMsg.style.display = 'none';
                } catch (e) {
                    headerOutput.textContent = '';
                    payloadOutput.textContent = '';
                    signatureOutput.textContent = '';
                    errorMsg.style.display = 'block';
                }
            }

            function copyToClipboard(targetId) {
                const element = document.getElementById(targetId);
                if (!element || !element.textContent) return;
                
                navigator.clipboard.writeText(element.textContent).then(() => {
                    // Optional: show a brief success indication
                    const btn = document.querySelector(`[data-target="${targetId}"]`);
                    const originalText = btn.textContent;
                    btn.textContent = '✓';
                    setTimeout(() => {
                        btn.textContent = originalText;
                    }, 1000);
                }).catch(err => {
                    console.error('Failed to copy text: ', err);
                });
            }

            // Event Listeners
            input.addEventListener('input', decodeJWT);
            
            clearBtn.addEventListener('click', () => {
                input.value = '';
                decodeJWT();
            });

            document.querySelectorAll('.btn-copy-small').forEach(btn => {
                btn.addEventListener('click', (e) => {
                    copyToClipboard(e.target.getAttribute('data-target'));
                });
            });
        })();