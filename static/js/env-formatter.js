(function () {
    const inputEl = document.getElementById('env-input');
    const outputEl = document.getElementById('env-output');
    const errorsEl = document.getElementById('env-errors');
    const statusEl = document.getElementById('env-status');
    const formatBtn = document.getElementById('format-btn');
    const clearBtn = document.getElementById('clear-btn');
    const copyBtn = document.getElementById('copy-btn');

    function setStatus(message) {
      statusEl.textContent = message;
    }

    function escapeHtml(text) {
      return text
        .replaceAll('&', '&amp;')
        .replaceAll('<', '<')
        .replaceAll('>', '>')
        .replaceAll('"', '"')
        .replaceAll("'", '&#39;');
    }

    function parseEnvLine(line, lineNumber) {
      const trimmed = line.trim();
      if (trimmed === '' || trimmed.startsWith('#')) {
        return { skip: true };
      }

      const equalIndex = trimmed.indexOf('=');
      if (equalIndex <= 0) {
        return {
          skip: false,
          error: 'Line ' + lineNumber + ': missing "=" separator.'
        };
      }

      const key = trimmed.slice(0, equalIndex).trim();
      const value = trimmed.slice(equalIndex + 1).trim();

      const keyPattern = /^[A-Za-z_][A-Za-z0-9_]*$/;
      if (!keyPattern.test(key)) {
        return {
          skip: false,
          error: 'Line ' + lineNumber + ': invalid key "' + key + '".'
        };
      }

      return {
        skip: false,
        key: key,
        value: value
      };
    }

    function formatEnv() {
      const raw = inputEl.value || '';
      const lines = raw.split(/\r?\n/);
      const entries = [];
      const errors = [];

      for (let i = 0; i < lines.length; i += 1) {
        const parsed = parseEnvLine(lines[i], i + 1);
        if (parsed.skip) {
          continue;
        }
        if (parsed.error) {
          errors.push(parsed.error);
          continue;
        }
        entries.push({ key: parsed.key, value: parsed.value });
      }

      entries.sort(function (a, b) {
        return a.key.localeCompare(b.key);
      });

      const formatted = entries.map(function (entry) {
        return entry.key + '=' + entry.value;
      }).join('\n');

      outputEl.value = formatted;
      errorsEl.innerHTML = errors.map(function (error) {
        return '<li>' + escapeHtml(error) + '</li>';
      }).join('');

      if (errors.length > 0 && formatted.length > 0) {
        setStatus('Formatted with validation warnings.');
      } else if (errors.length > 0) {
        setStatus('No valid entries found. Please fix input lines.');
      } else {
        setStatus('Formatting completed successfully.');
      }
    }

    function clearAll() {
      inputEl.value = '';
      outputEl.value = '';
      errorsEl.innerHTML = '';
      setStatus('Cleared.');
    }

    async function copyOutput() {
      if (!outputEl.value) {
        setStatus('Nothing to copy.');
        return;
      }

      try {
        await navigator.clipboard.writeText(outputEl.value);
        setStatus('Formatted output copied to clipboard.');
      } catch (error) {
        setStatus('Clipboard copy failed. Please copy manually.');
      }
    }

    formatBtn.addEventListener('click', formatEnv);
    clearBtn.addEventListener('click', clearAll);
    copyBtn.addEventListener('click', copyOutput);
  })();