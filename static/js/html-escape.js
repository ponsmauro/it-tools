(function initHtmlEscapeTool() {
        const input = document.getElementById('html-escape-input');
        const output = document.getElementById('html-escape-output');
        const status = document.getElementById('html-escape-status');
        const escapeButton = document.getElementById('btn-escape');
        const unescapeButton = document.getElementById('btn-unescape');
        const clearButton = document.getElementById('btn-clear');
        const copyButton = document.getElementById('btn-copy');

        if (!input || !output || !status || !escapeButton || !unescapeButton || !clearButton || !copyButton) {
            return;
        }

        const entityMap = {
            '&': '&amp;',
            '<': '<',
            '>': '>',
            '"': '"',
            "'": '&#39;'
        };

        const reverseEntityMap = {
            '&amp;': '&',
            '<': '<',
            '>': '>',
            '"': '"',
            '&#39;': "'"
        };

        function setStatus(message, isError) {
            status.textContent = message;
            status.classList.toggle('error', Boolean(isError));
        }

        function escapeHtml(value) {
            return value.replace(/[&<>"']/g, function (match) {
                return entityMap[match];
            });
        }

        function unescapeHtml(value) {
            return value.replace(/&(amp|lt|gt|quot|#39);/g, function (match) {
                return reverseEntityMap[match] || match;
            });
        }

        escapeButton.addEventListener('click', function () {
            output.value = escapeHtml(input.value);
            setStatus('Escaped successfully.', false);
        });

        unescapeButton.addEventListener('click', function () {
            output.value = unescapeHtml(input.value);
            setStatus('Unescaped successfully.', false);
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