(function () {
    const inputEl = document.getElementById('csv-input');
    const parseBtn = document.getElementById('parse-btn');
    const clearBtn = document.getElementById('clear-btn');
    const tableContainer = document.getElementById('csv-table-container');

    // CSV parser with quote-aware parsing
    function parseCSV(text) {
      const lines = text.split('\n');
      const result = [];
      for (let i = 0; i < lines.length; i++) {
        const line = lines[i].trim();
        if (!line) continue;
        
        const row = [];
        let inQuotes = false;
        let currentValue = '';
        
        for (let j = 0; j < line.length; j++) {
          const char = line[j];
          if (char === '"') {
            inQuotes = !inQuotes;
          } else if (char === ',' && !inQuotes) {
            row.push(currentValue);
            currentValue = '';
          } else {
            currentValue += char;
          }
        }
        row.push(currentValue);
        result.push(row);
      }
      return result;
    }

    function renderTable() {
      const text = inputEl.value.trim();
      if (!text) {
        tableContainer.innerHTML = '<div class="csv-placeholder">Enter CSV data and click Parse</div>';
        return;
      }

      try {
        const data = parseCSV(text);
        if (data.length === 0) {
          tableContainer.innerHTML = '<div class="csv-placeholder">No valid data found</div>';
          return;
        }

        let html = '<table class="csv-table">';
        
        // Header
        html += '<thead><tr>';
        for (let i = 0; i < data[0].length; i++) {
          html += `<th>${escapeHtml(data[0][i])}</th>`;
        }
        html += '</tr></thead>';
        
        // Body
        html += '<tbody>';
        for (let i = 1; i < data.length; i++) {
          html += '<tr>';
          for (let j = 0; j < data[i].length; j++) {
            html += `<td>${escapeHtml(data[i][j] || '')}</td>`;
          }
          html += '</tr>';
        }
        html += '</tbody></table>';
        
        tableContainer.innerHTML = html;
      } catch (err) {
        const errDiv = document.createElement('div');
        errDiv.style.color = 'var(--error-text)';
        errDiv.textContent = 'Error parsing CSV: ' + err.message;
        tableContainer.replaceChildren(errDiv);
      }
    }

    function escapeHtml(text) {
      return String(text)
        .replace(/&/g, '&amp;')
        .replace(/</g, '<')
        .replace(/>/g, '>')
        .replace(/"/g, '"')
        .replace(/'/g, '&#039;');
    }

    parseBtn.addEventListener('click', renderTable);
    
    clearBtn.addEventListener('click', () => {
      inputEl.value = '';
      tableContainer.innerHTML = '<div class="csv-placeholder">Enter CSV data and click Parse</div>';
    });
  })();