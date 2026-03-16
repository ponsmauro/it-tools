const colorPicker = document.getElementById('color-picker');
        const colorInput = document.getElementById('color-input');
        const colorPreview = document.getElementById('color-preview');
        
        const outHex = document.getElementById('out-hex');
        const outRgb = document.getElementById('out-rgb');
        const outHsl = document.getElementById('out-hsl');

        // Helper to parse color
        function parseColor(str) {
            const ctx = document.createElement('canvas').getContext('2d');
            ctx.fillStyle = str;
            if (ctx.fillStyle === 'var(--bg-color)000' && str.toLowerCase() !== 'black' && str !== 'var(--bg-color)000' && str !== 'var(--bg-color)') {
                // Invalid color might default to black in some browsers, but let's do a basic check
                // Actually, a better way is to check if assigning it changes the fillStyle
                ctx.fillStyle = '#123456';
                ctx.fillStyle = str;
                if (ctx.fillStyle === '#123456') return null;
            }
            return ctx.fillStyle; // Returns hex
        }

        function hexToRgb(hex) {
            const result = /^#?([a-f\d]{2})([a-f\d]{2})([a-f\d]{2})$/i.exec(hex);
            return result ? {
                r: parseInt(result[1], 16),
                g: parseInt(result[2], 16),
                b: parseInt(result[3], 16)
            } : null;
        }

        function rgbToHsl(r, g, b) {
            r /= 255, g /= 255, b /= 255;
            const max = Math.max(r, g, b), min = Math.min(r, g, b);
            let h, s, l = (max + min) / 2;

            if (max === min) {
                h = s = 0; // achromatic
            } else {
                const d = max - min;
                s = l > 0.5 ? d / (2 - max - min) : d / (max + min);
                switch (max) {
                    case r: h = (g - b) / d + (g < b ? 6 : 0); break;
                    case g: h = (b - r) / d + 2; break;
                    case b: h = (r - g) / d + 4; break;
                }
                h /= 6;
            }

            return {
                h: Math.round(h * 360),
                s: Math.round(s * 100),
                l: Math.round(l * 100)
            };
        }

        function updateColors(hex) {
            if (!hex) return;

            // Update inputs
            colorPicker.value = hex;
            colorPreview.style.backgroundColor = hex;

            // Update outputs
            outHex.value = hex.toUpperCase();
            
            const rgb = hexToRgb(hex);
            if (rgb) {
                outRgb.value = `rgb(${rgb.r}, ${rgb.g}, ${rgb.b})`;
                
                const hsl = rgbToHsl(rgb.r, rgb.g, rgb.b);
                outHsl.value = `hsl(${hsl.h}, ${hsl.s}%, ${hsl.l}%)`;
            }
        }

        function handleInput() {
            const val = colorInput.value.trim();
            if (!val) return;

            const hex = parseColor(val);
            if (hex) {
                updateColors(hex);
            }
        }

        colorPicker.addEventListener('input', (e) => {
            colorInput.value = e.target.value;
            updateColors(e.target.value);
        });

        colorInput.addEventListener('input', handleInput);

        // Copy buttons
        document.querySelectorAll('.btn-copy').forEach(btn => {
            btn.addEventListener('click', async (e) => {
                const targetId = e.target.getAttribute('data-target');
                const input = document.getElementById(targetId);
                if (input && input.value) {
                    try {
                        await navigator.clipboard.writeText(input.value);
                        const originalText = e.target.textContent;
                        e.target.textContent = 'Copied!';
                        setTimeout(() => {
                            e.target.textContent = originalText;
                        }, 2000);
                    } catch (err) {
                        console.error('Failed to copy', err);
                    }
                }
            });
        });

        // Initial update
        updateColors(colorPicker.value);