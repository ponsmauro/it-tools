(function initBase64Tool() {
            const input = document.getElementById('base64-input');
            const output = document.getElementById('base64-output');
            const status = document.getElementById('base64-status');
            const processButton = document.getElementById('btn-process');
            const clearButton = document.getElementById('btn-clear');
            const copyButton = document.getElementById('btn-copy');
            const encodeTab = document.getElementById('encode-tab');
            const decodeTab = document.getElementById('decode-tab');
            const fileInput = document.getElementById('file-input');
            const fileInfo = document.getElementById('file-info');
            const fileUploadSection = document.getElementById('file-upload-section');
            
            let currentMode = 'encode';
            let currentFile = null;

            function setStatus(message, isError = false) {
                status.textContent = message;
                status.classList.toggle('error', isError);
                status.classList.add('visible');
                
                setTimeout(() => {
                    status.classList.remove('visible');
                }, 3000);
            }

            function switchMode(mode) {
                currentMode = mode;
                
                if (mode === 'encode') {
                    encodeTab.classList.add('active');
                    decodeTab.classList.remove('active');
                    processButton.textContent = '{{ tr "tools.base64.encode-btn" }}';
                    fileUploadSection.style.display = 'block';
                } else {
                    encodeTab.classList.remove('active');
                    decodeTab.classList.add('active');
                    processButton.textContent = '{{ tr "tools.base64.decode-btn" }}';
                    fileUploadSection.style.display = 'none';
                }
                
                // Clear outputs when switching modes
                output.value = '';
                currentFile = null;
                fileInfo.textContent = '';
            }

            function encodeBase64() {
                try {
                    if (currentFile) {
                        // File is already processed in the file input handler
                        return;
                    }
                    
                    if (!input.value.trim()) {
                        setStatus('{{ tr "tools.base64.no-input" }}', true);
                        return;
                    }
                    
                    const urlSafe = document.getElementById('url-safe').checked;
                    const includePadding = document.getElementById('include-padding').checked;
                    
                    // Standard Base64 encoding
                    let encoded = btoa(input.value);
                    
                    // Apply URL-safe transformation if needed
                    if (urlSafe) {
                        encoded = encoded.replace(/\+/g, '-').replace(/\//g, '_');
                    }
                    
                    // Remove padding if needed
                    if (!includePadding) {
                        encoded = encoded.replace(/=+$/, '');
                    }
                    
                    output.value = encoded;
                    setStatus('{{ tr "tools.base64.encode-success" }}');
                    document.getElementById('btn-download').disabled = true;
                } catch (error) {
                    setStatus('{{ tr "tools.base64.encode-error" }}', true);
                    console.error('Encode error:', error);
                }
            }

            function decodeBase64() {
                try {
                    if (!input.value.trim()) {
                        setStatus('{{ tr "tools.base64.no-input" }}', true);
                        return;
                    }
                    
                    const urlSafe = document.getElementById('url-safe').checked;
                    let encoded = input.value.trim();
                    
                    // Restore standard Base64 characters if URL-safe was used
                    if (urlSafe) {
                        encoded = encoded.replace(/-/g, '+').replace(/_/g, '/');
                    }
                    
                    // Add padding if needed
                    while (encoded.length % 4 !== 0) {
                        encoded += '=';
                    }
                    
                    // Try to decode
                    const decoded = atob(encoded);
                    output.value = decoded;
                    
                    // Enable download button if it looks like binary data
                    const isBinary = containsBinaryData(decoded);
                    document.getElementById('btn-download').disabled = !isBinary;
                    
                    setStatus('{{ tr "tools.base64.decode-success" }}');
                } catch (error) {
                    setStatus('{{ tr "tools.base64.decode-error" }}', true);
                    console.error('Decode error:', error);
                }
            }
            
            function containsBinaryData(str) {
                // Check if the string contains non-printable characters
                for (let i = 0; i < str.length; i++) {
                    const code = str.charCodeAt(i);
                    if (code < 32 && code !== 9 && code !== 10 && code !== 13) {
                        return true;
                    }
                }
                return false;
            }
            
            function downloadDecodedFile() {
                if (!output.value) {
                    setStatus('{{ tr "tools.base64.nothing-to-download" }}', true);
                    return;
                }
                
                try {
                    // Create a Blob from the decoded data
                    const blob = new Blob([output.value], { type: 'application/octet-stream' });
                    
                    // Create a download link
                    const url = URL.createObjectURL(blob);
                    const a = document.createElement('a');
                    a.href = url;
                    a.download = 'decoded_file';
                    
                    // Trigger download
                    document.body.appendChild(a);
                    a.click();
                    
                    // Cleanup
                    setTimeout(() => {
                        document.body.removeChild(a);
                        URL.revokeObjectURL(url);
                    }, 0);
                    
                    setStatus('{{ tr "tools.base64.download-started" }}');
                } catch (error) {
                    setStatus('{{ tr "tools.base64.download-error" }}', true);
                    console.error('Download error:', error);
                }
            }

            function processInput() {
                if (currentMode === 'encode') {
                    encodeBase64();
                } else {
                    decodeBase64();
                }
            }

            function clearAll() {
                input.value = '';
                output.value = '';
                currentFile = null;
                fileInfo.textContent = '';
                fileInput.value = '';
                setStatus('{{ tr "tools.base64.cleared" }}');
                input.focus();
            }

            async function copyOutput() {
                try {
                    if (!output.value) {
                        setStatus('{{ tr "tools.base64.nothing-to-copy" }}', true);
                        return;
                    }
                    
                    try {
                        await navigator.clipboard.writeText(output.value);
                        setStatus('{{ tr "tools.base64.copy-success" }}');
                    } catch (clipboardError) {
                        console.error('Clipboard API error:', clipboardError);
                        copyToClipboardFallback(output.value);
                        setStatus('{{ tr "tools.base64.copy-success" }}');
                    }
                } catch (error) {
                    setStatus('{{ tr "tools.base64.copy-error" }}', true);
                    console.error('Copy error:', error);
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
                    if (!successful) {
                        console.error('Fallback copy failed');
                    }
                } catch (err) {
                    console.error('Fallback copy error:', err);
                }
                
                document.body.removeChild(textArea);
            }

            function handleFileUpload(event) {
                const file = event.target.files[0];
                if (!file) return;
                
                // Check file size (limit to 5MB for browser performance)
                if (file.size > 5 * 1024 * 1024) {
                    setStatus('{{ tr "tools.base64.file-too-large" }}', true);
                    return;
                }
                
                currentFile = file;
                fileInfo.textContent = `${file.name} (${formatFileSize(file.size)})`;
                
                const reader = new FileReader();
                
                reader.onload = function(e) {
                    try {
                        // Get the base64 part from the data URL
                        const base64String = e.target.result.split(',')[1];
                        
                        // Apply URL-safe and padding options
                        const urlSafe = document.getElementById('url-safe').checked;
                        const includePadding = document.getElementById('include-padding').checked;
                        
                        let encodedOutput = base64String;
                        
                        if (urlSafe) {
                            encodedOutput = encodedOutput.replace(/\+/g, '-').replace(/\//g, '_');
                        }
                        
                        if (!includePadding) {
                            encodedOutput = encodedOutput.replace(/=+$/, '');
                        }
                        
                        output.value = encodedOutput;
                        setStatus('{{ tr "tools.base64.file-encoded" }}');
                    } catch (error) {
                        setStatus('{{ tr "tools.base64.file-error" }}', true);
                        console.error('File processing error:', error);
                    }
                };
                
                reader.onerror = function() {
                    setStatus('{{ tr "tools.base64.file-read-error" }}', true);
                };
                
                reader.readAsDataURL(file);
            }

            function formatFileSize(bytes) {
                if (bytes < 1024) return bytes + ' bytes';
                else if (bytes < 1048576) return (bytes / 1024).toFixed(1) + ' KB';
                else return (bytes / 1048576).toFixed(1) + ' MB';
            }

            // Event listeners
            encodeTab.addEventListener('click', () => switchMode('encode'));
            decodeTab.addEventListener('click', () => switchMode('decode'));
            processButton.addEventListener('click', processInput);
            clearButton.addEventListener('click', clearAll);
            copyButton.addEventListener('click', copyOutput);
            fileInput.addEventListener('change', handleFileUpload);
            document.getElementById('btn-download').addEventListener('click', downloadDecodedFile);
            
            // Initialize
            switchMode('encode');
        })();