(function initUrlEncoderTool() {
        const input = document.getElementById('url-encoder-input');
        const output = document.getElementById('url-encoder-output');
        const status = document.getElementById('url-encoder-status');
        const encodeButton = document.getElementById('btn-url-encode');
        const decodeButton = document.getElementById('btn-url-decode');
        const clearButton = document.getElementById('btn-url-clear');
        const copyButton = document.getElementById('btn-url-copy');

        if (!input || !output || !status || !encodeButton || !decodeButton || !clearButton || !copyButton) {
            return;
        }

        function setStatus(message, isError) {
            status.textContent = message;
            status.classList.toggle('error', Boolean(isError));
        }

        encodeButton.addEventListener('click', function () {
            try {
                output.value = encodeURIComponent(input.value);
                setStatus('Encoded successfully.', false);
            } catch (error) {
                output.value = '';
                setStatus('Encoding failed.', true);
            }
        });

        decodeButton.addEventListener('click', function () {
            try {
                output.value = decodeURIComponent(input.value);
                setStatus('Decoded successfully.', false);
            } catch (error) {
                output.value = '';
                setStatus('Invalid encoded input for decoding.', true);
            }
        });

        clearButton.addEventListener('click', function () {
            input.value = '';
            output.value = '';
            setStatus('Cleared.', false);
            input.focus();
        });

        copyButton.addEventListener('click', async function () {
            try {
                await navigator.clipboard.writeText(output.value);
                setStatus('Output copied to clipboard.', false);
            } catch (error) {
                setStatus('Clipboard copy failed.', true);
            }
        });
    })();