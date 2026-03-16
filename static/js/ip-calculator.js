(function () {
    const inputEl = document.getElementById('ip-input');
    const calcBtn = document.getElementById('calc-btn');
    const errorEl = document.getElementById('ip-error');
    
    const resIp = document.getElementById('res-ip');
    const resNetwork = document.getElementById('res-network');
    const resRange = document.getElementById('res-range');
    const resBroadcast = document.getElementById('res-broadcast');
    const resHosts = document.getElementById('res-hosts');
    const resUsable = document.getElementById('res-usable');
    const resMask = document.getElementById('res-mask');

    function ip2long(ip) {
      return ip.split('.').reduce((ipInt, octet) => (ipInt << 8) + parseInt(octet, 10), 0) >>> 0;
    }

    function long2ip(ipInt) {
      return (
        ((ipInt >>> 24) & 255) + '.' +
        ((ipInt >>> 16) & 255) + '.' +
        ((ipInt >>> 8) & 255) + '.' +
        (ipInt & 255)
      );
    }

    function calculate() {
      const input = inputEl.value.trim();
      errorEl.style.display = 'none';
      
      try {
        let ipStr = input;
        let cidr = 24;
        
        if (input.includes('/')) {
          const parts = input.split('/');
          ipStr = parts[0];
          cidr = parseInt(parts[1], 10);
        }
        
        if (!/^(\d{1,3}\.){3}\d{1,3}$/.test(ipStr)) {
          throw new Error('Invalid IP address format');
        }
        
        if (isNaN(cidr) || cidr < 0 || cidr > 32) {
          throw new Error('Invalid CIDR notation (must be 0-32)');
        }
        
        const ipParts = ipStr.split('.').map(Number);
        if (ipParts.some(p => p < 0 || p > 255)) {
          throw new Error('IP octets must be between 0 and 255');
        }
        
        const ipLong = ip2long(ipStr);
        const maskLong = cidr === 0 ? 0 : (~0 << (32 - cidr)) >>> 0;
        const networkLong = (ipLong & maskLong) >>> 0;
        const broadcastLong = (networkLong | ~maskLong) >>> 0;
        
        const totalHosts = cidr === 32 ? 1 : Math.pow(2, 32 - cidr);
        const usableHosts = cidr >= 31 ? 0 : totalHosts - 2;
        
        let firstHost = '-';
        let lastHost = '-';
        
        if (usableHosts > 0) {
          firstHost = long2ip(networkLong + 1);
          lastHost = long2ip(broadcastLong - 1);
        }
        
        resIp.textContent = ipStr;
        resNetwork.textContent = long2ip(networkLong);
        resRange.textContent = usableHosts > 0 ? `${firstHost} - ${lastHost}` : 'N/A';
        resBroadcast.textContent = long2ip(broadcastLong);
        resHosts.textContent = totalHosts.toLocaleString();
        resUsable.textContent = usableHosts.toLocaleString();
        resMask.textContent = long2ip(maskLong);
        
      } catch (err) {
        errorEl.textContent = err.message;
        errorEl.style.display = 'block';
        
        resIp.textContent = '-';
        resNetwork.textContent = '-';
        resRange.textContent = '-';
        resBroadcast.textContent = '-';
        resHosts.textContent = '-';
        resUsable.textContent = '-';
        resMask.textContent = '-';
      }
    }

    calcBtn.addEventListener('click', calculate);
    inputEl.addEventListener('keypress', (e) => {
      if (e.key === 'Enter') calculate();
    });
    
    // Initial calculation
    calculate();
  })();