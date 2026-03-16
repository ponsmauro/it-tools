const inputEl = document.getElementById('qr-input');
        const qrContainer = document.getElementById('qr-container');
        const btnDownload = document.getElementById('btn-download');
        let qrcode = null;

        function generateQR() {
            const text = inputEl.value.trim();
            
            // Clear previous
            qrContainer.innerHTML = '';
            btnDownload.disabled = true;

            if (!text) return;

            try {
                qrcode = new QRCode(qrContainer, {
                    text: text,
                    width: 256,
                    height: 256,
                    colorDark : "#000000",
                    colorLight : "white",
                    correctLevel : QRCode.CorrectLevel.H
                });

                // Enable download after a short delay to allow image generation
                setTimeout(() => {
                    const img = qrContainer.querySelector('img');
                    if (img && img.src) {
                        btnDownload.disabled = false;
                    }
                }, 100);
            } catch (err) {
                console.error('Error generating QR code', err);
            }
        }

        inputEl.addEventListener('input', generateQR);

        document.getElementById('btn-clear').addEventListener('click', () => {
            inputEl.value = '';
            generateQR();
        });

        btnDownload.addEventListener('click', () => {
            const img = qrContainer.querySelector('img');
            if (!img || !img.src) return;

            const link = document.createElement('a');
            link.download = 'qrcode.png';
            link.href = img.src;
            document.body.appendChild(link);
            link.click();
            document.body.removeChild(link);
        });

        // Initial generation if there's a value
        if (inputEl.value) {
            generateQR();
        }