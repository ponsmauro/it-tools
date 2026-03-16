(function () {
    const inputEl = document.getElementById('gzip-input');
    const outputEl = document.getElementById('gzip-output');
    const processBtn = document.getElementById('process-btn');
    const clearBtn = document.getElementById('clear-btn');
    const copyBtn = document.getElementById('copy-btn');
    const compressTab = document.getElementById('compress-tab');
    const decompressTab = document.getElementById('decompress-tab');
    const statsEl = document.getElementById('gzip-stats');
    
    let mode = 'compress';

    function switchMode(newMode) {
      mode = newMode;
      if (mode === 'compress') {
        compressTab.classList.add('active');
        decompressTab.classList.remove('active');
        processBtn.textContent = 'Compress';
        inputEl.placeholder = 'Paste text to compress...';
      } else {
        compressTab.classList.remove('active');
        decompressTab.classList.add('active');
        processBtn.textContent = 'Decompress';
        inputEl.placeholder = 'Paste Base64 encoded gzip data...';
      }
      inputEl.value = '';
      outputEl.value = '';
      statsEl.style.display = 'none';
    }

    function formatBytes(bytes) {
      if (bytes === 0) return '0 Bytes';
      const k = 1024;
      const sizes = ['Bytes', 'KB', 'MB', 'GB'];
      const i = Math.floor(Math.log(bytes) / Math.log(k));
      return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i];
    }

    function process() {
      const input = inputEl.value;
      if (!input) return;

      try {
        if (mode === 'compress') {
          // Compress
          const compressed = pako.gzip(input);
          // Convert to Base64 in chunks to avoid call stack overflow for large inputs
          let binary = '';
          const chunkSize = 8192;
          for (let i = 0; i < compressed.length; i += chunkSize) {
            binary += String.fromCharCode(...compressed.subarray(i, i + chunkSize));
          }
          const base64 = btoa(binary);
          outputEl.value = base64;
          
          // Stats
          const origSize = new Blob([input]).size;
          const compSize = compressed.length;
          const ratio = ((1 - (compSize / origSize)) * 100).toFixed(1);
          
          document.getElementById('stat-original').textContent = formatBytes(origSize);
          document.getElementById('stat-compressed').textContent = formatBytes(compSize);
          document.getElementById('stat-ratio').textContent = ratio > 0 ? `-${ratio}%` : `+${Math.abs(ratio)}%`;
          statsEl.style.display = 'block';
          
        } else {
          // Decompress
          // Decode Base64
          const binaryString = atob(input);
          const bytes = new Uint8Array(binaryString.length);
          for (let i = 0; i < binaryString.length; i++) {
            bytes[i] = binaryString.charCodeAt(i);
          }
          
          // Decompress
          const decompressed = pako.ungzip(bytes, { to: 'string' });
          outputEl.value = decompressed;
          statsEl.style.display = 'none';
        }
      } catch (err) {
        outputEl.value = `Error: ${err.message}`;
        statsEl.style.display = 'none';
      }
    }

    compressTab.addEventListener('click', () => switchMode('compress'));
    decompressTab.addEventListener('click', () => switchMode('decompress'));
    processBtn.addEventListener('click', process);
    
    clearBtn.addEventListener('click', () => {
      inputEl.value = '';
      outputEl.value = '';
      statsEl.style.display = 'none';
    });
    
    copyBtn.addEventListener('click', async () => {
      if (!outputEl.value) return;
      try {
        await navigator.clipboard.writeText(outputEl.value);
        const originalText = copyBtn.textContent;
        copyBtn.textContent = 'Copied!';
        setTimeout(() => copyBtn.textContent = originalText, 2000);
      } catch (err) {
        console.error('Failed to copy', err);
      }
    });
  })();