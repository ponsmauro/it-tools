(function () {
    const originalEl = document.getElementById('diff-original');
    const modifiedEl = document.getElementById('diff-modified');
    const outputEl = document.getElementById('diff-output');
    const compareBtn = document.getElementById('compare-btn');
    const clearBtn = document.getElementById('clear-btn');

    function escapeHtml(text) {
      return text
        .replaceAll('&', '&amp;')
        .replaceAll('<', '<')
        .replaceAll('>', '>')
        .replaceAll('"', '"')
        .replaceAll("'", '&#39;');
    }

    function compare() {
      const text1 = originalEl.value;
      const text2 = modifiedEl.value;

      if (!text1 && !text2) {
        outputEl.innerHTML = '<div class="diff-placeholder">Please enter text to compare</div>';
        return;
      }

      try {
        const dmp = new diff_match_patch();
        const diffs = dmp.diff_main(text1, text2);
        dmp.diff_cleanupSemantic(diffs);

        let html = '';
        for (let i = 0; i < diffs.length; i++) {
          const op = diffs[i][0];    // Operation (insert, delete, equal)
          const data = diffs[i][1];  // Text of change
          const escapedData = escapeHtml(data);

          if (op === 1) { // Insert
            html += `<span class="diff-added">${escapedData}</span>`;
          } else if (op === -1) { // Delete
            html += `<span class="diff-removed">${escapedData}</span>`;
          } else { // Equal
            html += `<span>${escapedData}</span>`;
          }
        }

        outputEl.innerHTML = html || '<div class="diff-placeholder">Texts are identical</div>';
      } catch (err) {
        const errDiv = document.createElement('div');
        errDiv.style.color = 'var(--error-text)';
        errDiv.textContent = 'Error: ' + err.message;
        outputEl.replaceChildren(errDiv);
      }
    }

    compareBtn.addEventListener('click', compare);
    
    clearBtn.addEventListener('click', () => {
      originalEl.value = '';
      modifiedEl.value = '';
      outputEl.innerHTML = '<div class="diff-placeholder">Click "Compare" to see differences</div>';
    });
  })();