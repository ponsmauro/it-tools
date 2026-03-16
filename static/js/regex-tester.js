const regexInput = document.getElementById('regex-input');
        const flagsInput = document.getElementById('flags-input');
        const testString = document.getElementById('test-string');
        const highlightContent = document.getElementById('highlight-content');
        const matchesOutput = document.getElementById('matches-output');
        const errorMsg = document.getElementById('regex-error');

        // Sync scrolling between textarea and backdrop
        testString.addEventListener('scroll', () => {
            document.getElementById('highlight-backdrop').scrollTop = testString.scrollTop;
        });

        function escapeHtml(unsafe) {
            return (unsafe || '').toString()
                .replace(/&/g, "&amp;")
                .replace(/</g, "&lt;")
                .replace(/>/g, "&gt;")
                .replace(/"/g, "&quot;")
                .replace(/'/g, "&#039;");
        }

        function updateRegex() {
            const pattern = regexInput.value;
            const flags = flagsInput.value;
            const text = testString.value;

            if (!pattern) {
                highlightContent.textContent = text;
                matchesOutput.innerHTML = `<div class="no-matches">{{ tr "tools.regex-tester.no-match" }}</div>`;
                errorMsg.style.display = 'none';
                return;
            }

            try {
                const regex = new RegExp(pattern, flags);
                errorMsg.style.display = 'none';

                // Find matches
                let match;
                const matches = [];
                
                // If global flag is not set, we only get the first match
                if (!regex.global) {
                    match = regex.exec(text);
                    if (match) matches.push(match);
                } else {
                    // Reset lastIndex just in case
                    const MAX_MATCHES = 500;
                    regex.lastIndex = 0;
                    while ((match = regex.exec(text)) !== null && matches.length < MAX_MATCHES) {
                        matches.push(match);
                        // Prevent infinite loop with zero-length matches
                        if (match.index === regex.lastIndex) {
                            regex.lastIndex++;
                        }
                    }
                }

                // Render matches list
                if (matches.length === 0) {
                    matchesOutput.innerHTML = `<div class="no-matches">{{ tr "tools.regex-tester.no-match" }}</div>`;
                    highlightContent.textContent = text;
                } else {
                    let html = '';
                    if (matches.length === 500) {
                        html += `<div class="no-matches" style="color: var(--text-secondary); margin-bottom: 8px;">Showing first 500 matches</div>`;
                    }
                    matches.forEach((m, i) => {
                        html += `<div class="match-item">`;
                        html += `<div class="match-header">Match ${i + 1}: "${escapeHtml(m[0])}" (Index: ${m.index})</div>`;
                        
                        // Render capture groups
                        for (let j = 1; j < m.length; j++) {
                            if (m[j] !== undefined) {
                                html += `<div class="match-group">Group ${j}: "${escapeHtml(m[j])}"</div>`;
                            }
                        }
                        html += `</div>`;
                    });
                    matchesOutput.innerHTML = html;

                    // Render highlights
                    let highlightedText = '';
                    let lastIndex = 0;

                    matches.forEach(m => {
                        // Add text before match
                        highlightedText += escapeHtml(text.substring(lastIndex, m.index));
                        // Add highlighted match
                        highlightedText += `<span class="match-highlight">${escapeHtml(m[0])}</span>`;
                        lastIndex = m.index + m[0].length;
                    });
                    // Add remaining text
                    highlightedText += escapeHtml(text.substring(lastIndex));

                    // Handle trailing newlines properly for the backdrop
                    if (text.endsWith('\n')) {
                        highlightedText += '<br>';
                    }

                    highlightContent.innerHTML = highlightedText;
                }

            } catch (err) {
                errorMsg.textContent = err.message;
                errorMsg.style.display = 'block';
                highlightContent.textContent = text;
                matchesOutput.innerHTML = `<div class="no-matches">{{ tr "tools.regex-tester.error" }}</div>`;
            }
        }

        regexInput.addEventListener('input', updateRegex);
        flagsInput.addEventListener('input', updateRegex);
        testString.addEventListener('input', updateRegex);

        // Initial update
        updateRegex();