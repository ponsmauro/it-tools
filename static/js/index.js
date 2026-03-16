let activeToolLink = null;

        document.addEventListener('DOMContentLoaded', () => {
            const searchInput = document.getElementById('sidebar-search');
            const toolLinks = document.querySelectorAll('.tool-link');
            
            searchInput.addEventListener('input', (e) => {
                const query = e.target.value.toLowerCase();
                toolLinks.forEach(link => {
                    const text = link.textContent.toLowerCase();
                    link.style.display = text.includes(query) ? 'block' : 'none';
                });
            });
        });

        // Toggle sidebar function removed as requested

        async function loadToolContent(toolId) {
            // Notify current tool content it is being replaced (allows cleanup of intervals, etc.)
            document.dispatchEvent(new CustomEvent('toolunload'));

            // Clear previous active
            if (activeToolLink) activeToolLink.classList.remove('active');
            
            // Set active
            const toolLink = document.querySelector(`[data-tool="${toolId}"]`);
            toolLink.classList.add('active');
            activeToolLink = toolLink;
            
            // Show loading
            const contentArea = document.getElementById('content-area');
            const loadingDiv = document.createElement('div');
            loadingDiv.className = 'loading';
            loadingDiv.textContent = 'Loading ' + toolId.toUpperCase() + '...';
            contentArea.replaceChildren(loadingDiv);
            
            // AJAX call
            try {
                const response = await fetch(`/tools/${toolId}`, {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({ action: 'open' })
                });
                
                if (response.ok) {
                    const html = await response.text();
                    contentArea.innerHTML = html;
                } else {
                    contentArea.innerHTML = '<div class="loading" style="color: var(--error-text);">Error loading tool</div>';
                }
            } catch (error) {
                contentArea.innerHTML = '<div class="loading" style="color: var(--error-text);">Network error</div>';
                console.error('Load error:', error);
            }
        }