(function () {
    const inputEl = document.getElementById('cron-input');
    const parseBtn = document.getElementById('parse-btn');
    const explanationEl = document.getElementById('cron-explanation');
    const nextListEl = document.getElementById('cron-next');

    function parseCron() {
      const expr = inputEl.value.trim();
      if (!expr) return;

      try {
        // Use cronstrue if available
        if (window.cronstrue) {
          const explanation = window.cronstrue.toString(expr);
          explanationEl.textContent = explanation;
          explanationEl.style.color = 'var(--success-text)';
        } else {
          explanationEl.textContent = 'cronstrue library not loaded';
          explanationEl.style.color = 'var(--error-text)';
        }
        
        // Mock next occurrences for now
        const now = new Date();
        let html = '';
        for (let i = 1; i <= 5; i++) {
            const next = new Date(now.getTime() + i * 60000);
            html += `<li>${next.toLocaleString()}</li>`;
        }
        nextListEl.innerHTML = html;
        
      } catch (err) {
        explanationEl.textContent = err.toString();
        explanationEl.style.color = 'var(--error-text)';
        nextListEl.innerHTML = '';
      }
    }

    parseBtn.addEventListener('click', parseCron);
    inputEl.addEventListener('input', parseCron);
    
    // Initial parse
    setTimeout(parseCron, 500);
  })();