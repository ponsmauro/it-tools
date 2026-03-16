// Character sets
        const charSets = {
            uppercase: 'ABCDEFGHIJKLMNOPQRSTUVWXYZ',
            lowercase: 'abcdefghijklmnopqrstuvwxyz',
            numbers: '0123456789',
            symbols: '!@#$%^&*()_+-=[]{}|;:,.<>?'
        };
        
        const similarChars = 'iIlL1oO0';
        const ambiguousChars = '{}[]()/\\\'"`~,;:.<>';
        let currentPassword = '';
        
        // Initialize
        document.addEventListener('DOMContentLoaded', function() {
            generatePassword();
            
            // Add event listeners to checkboxes to ensure at least one is checked
            const checkboxes = document.querySelectorAll('.checkboxes input[type="checkbox"]');
            checkboxes.forEach(checkbox => {
                checkbox.addEventListener('change', function() {
                    const anyChecked = Array.from(checkboxes).some(cb => cb.checked);
                    if (!anyChecked) {
                        this.checked = true;
                        showStatus('At least one character set must be selected', 'error');
                    }
                });
            });
        });
        
        function updateLength() {
            const length = document.getElementById('length').value;
            document.getElementById('length-value').textContent = length;
        }
        
        function getSelectedCharset() {
            let charset = '';
            const customChars = document.getElementById('custom-chars').value;
            
            // If custom charset is provided, use it exclusively
            if (customChars) {
                return customChars;
            }
            
            // Otherwise build charset from selected options
            if (document.getElementById('uppercase').checked) {
                charset += charSets.uppercase;
            }
            if (document.getElementById('lowercase').checked) {
                charset += charSets.lowercase;
            }
            if (document.getElementById('numbers').checked) {
                charset += charSets.numbers;
            }
            if (document.getElementById('symbols').checked) {
                charset += charSets.symbols;
            }
            
            // Remove similar characters if option is checked
            if (document.getElementById('exclude-similar').checked) {
                charset = Array.from(charset)
                    .filter(char => !similarChars.includes(char))
                    .join('');
            }
            
            // Remove ambiguous characters if option is checked
            if (document.getElementById('exclude-ambiguous').checked) {
                charset = Array.from(charset)
                    .filter(char => !ambiguousChars.includes(char))
                    .join('');
            }
            
            return charset;
        }
        
        function generateRandomPassword(length, charset) {
            // Use cryptographically secure random values
            const randomValues = new Uint32Array(length * 2); // Extra values in case we need to skip repeats
            crypto.getRandomValues(randomValues);
            
            let password = '';
            let index = 0;
            const noRepeats = document.getElementById('no-repeats').checked;
            
            for (let i = 0; i < length; i++) {
                if (index >= randomValues.length) {
                    // If we somehow run out of random values, generate more
                    crypto.getRandomValues(randomValues);
                    index = 0;
                }
                
                const randomIndex = randomValues[index++] % charset.length;
                const nextChar = charset.charAt(randomIndex);
                
                // If no-repeats is enabled and this character would be a repeat, try again
                if (noRepeats && password.length > 0 && nextChar === password[password.length - 1]) {
                    // Try a different character, but don't increment i so we still get the right length
                    i--;
                    continue;
                }
                
                password += nextChar;
            }
            
            return password;
        }
        
        function generatePassword() {
            const length = parseInt(document.getElementById('length').value);
            const charset = getSelectedCharset();
            
            if (charset.length === 0) {
                showStatus('Please select at least one character set', 'error');
                return;
            }
            
            currentPassword = generateRandomPassword(length, charset);
            document.getElementById('password').value = currentPassword;
            document.getElementById('copy-btn').disabled = false;
            
            calculatePasswordStrength(currentPassword);
            
            // Clear the password list
            document.getElementById('password-list').innerHTML = '';
        }
        
        function generateMultiple() {
            const length = parseInt(document.getElementById('length').value);
            const charset = getSelectedCharset();
            const count = 5;
            
            if (charset.length === 0) {
                showStatus('Please select at least one character set', 'error');
                return;
            }
            
            const list = document.getElementById('password-list');
            list.replaceChildren();

            for (let i = 0; i < count; i++) {
                const password = generateRandomPassword(length, charset);
                const item = document.createElement('div');
                item.className = 'password-item';
                item.dataset.password = password;
                const code = document.createElement('code');
                code.textContent = password;
                item.appendChild(code);
                item.addEventListener('click', () => selectPassword(item.dataset.password, item));
                list.appendChild(item);
            }
        }
        
        function selectPassword(password, selectedItem) {
            currentPassword = password;
            document.getElementById('password').value = password;
            document.getElementById('copy-btn').disabled = false;
            calculatePasswordStrength(password);
            
            document.querySelectorAll('.password-item').forEach(item => {
                item.classList.toggle('selected', item === selectedItem);
            });
        }
        
        function showStatus(message, type) {
            const strengthText = document.getElementById('strength-text');
            const original = strengthText.textContent;
            strengthText.textContent = message;
            strengthText.className = 'strength-text status-' + type;
            setTimeout(() => {
                strengthText.textContent = original;
                strengthText.className = 'strength-text';
            }, 2500);
        }

        function copyPassword() {
            const password = document.getElementById('password').value;
            if (!password) return;
            
            const btn = document.getElementById('copy-btn');
            const original = btn.textContent;

            navigator.clipboard.writeText(password).then(() => {
                btn.textContent = '✓ Copied!';
                btn.classList.add('state-copied');
                setTimeout(() => {
                    btn.textContent = original;
                    btn.classList.remove('state-copied');
                }, 1500);
            }).catch(err => {
                console.error('Clipboard API error:', err);
                showStatus('Copy failed — please copy manually', 'error');
            });
        }
        
        function calculatePasswordStrength(password) {
            // Enhanced password strength calculation
            let strength = 0;
            
            // Length contribution (up to 40 points)
            strength += Math.min(40, password.length * 2);
            
            // Character variety contribution (up to 60 points)
            const hasUpper = /[A-Z]/.test(password);
            const hasLower = /[a-z]/.test(password);
            const hasNumber = /[0-9]/.test(password);
            const hasSymbol = /[^A-Za-z0-9]/.test(password);
            
            const variety = (hasUpper ? 15 : 0) + 
                           (hasLower ? 15 : 0) + 
                           (hasNumber ? 15 : 0) + 
                           (hasSymbol ? 15 : 0);
            
            strength += variety;
            
            // Penalize for patterns and repetitions
            const patterns = [
                /abc/i, /bcd/i, /cde/i, /def/i, /efg/i, /fgh/i, /ghi/i, /hij/i, // Sequential letters
                /123/, /234/, /345/, /456/, /567/, /678/, /789/, // Sequential numbers
                /(.)\1\1/ // Three of the same character in a row
            ];
            
            for (const pattern of patterns) {
                if (pattern.test(password)) {
                    strength -= 10;
                }
            }
            
            // Ensure strength is between 0 and 100
            strength = Math.max(0, Math.min(100, strength));
            
            // Update the strength meter
            const strengthBar = document.getElementById('strength-bar');
            const strengthText = document.getElementById('strength-text');
            
            strengthBar.style.width = strength + '%';
            
            if (strength < 40) {
                strengthBar.style.background = 'var(--error-text)';
                strengthText.textContent = 'Weak';
            } else if (strength < 70) {
                strengthBar.style.background = '#f59e0b';
                strengthText.textContent = 'Moderate';
            } else if (strength < 90) {
                strengthBar.style.background = '#10b981';
                strengthText.textContent = 'Strong';
            } else {
                strengthBar.style.background = '#059669';
                strengthText.textContent = 'Very Strong';
            }
            
            // Add detailed strength info
            let details = [];
            if (password.length < 12) details.push('Consider using at least 12 characters');
            if (!hasUpper) details.push('Add uppercase letters');
            if (!hasLower) details.push('Add lowercase letters');
            if (!hasNumber) details.push('Add numbers');
            if (!hasSymbol) details.push('Add symbols');
            
            if (details.length > 0) {
                strengthText.innerHTML = `${strengthText.textContent} <span class="strength-tips">(${details.join(', ')})</span>`;
            }
        }