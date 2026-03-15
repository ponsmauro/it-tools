import os
import subprocess
import re
from pathlib import Path

def run_cmd(cmd):
    try:
        result = subprocess.run(cmd, shell=True, capture_output=True, text=True)
        return result.stdout.strip()
    except Exception as e:
        return str(e)

def analyze_backend():
    print("Analyzing Backend...")
    # Count Go files and LOC
    go_files = run_cmd("find . -name '*.go' | wc -l")
    go_loc = run_cmd("find . -name '*.go' -exec cat {} + | wc -l")
    
    # Run tests and get coverage
    coverage_out = run_cmd("go test ./... -coverprofile=coverage.out")
    coverage_match = re.search(r'coverage: (\d+\.\d+)%', coverage_out)
    coverage = float(coverage_match.group(1)) if coverage_match else 0.0
    
    # Run go vet
    vet_out = run_cmd("go vet ./...")
    vet_issues_count = len([line for line in vet_out.split('\n') if line.strip()]) if vet_out else 0
    
    # Calculate score (base 100, -5 per vet issue, + coverage weight)
    score = max(0, min(100, (coverage * 0.8) + (20 - vet_issues_count * 5)))
    
    recommendations = []
    if coverage < 80:
        recommendations.append(f"Test coverage is low ({coverage}%). Aim for at least 80% by writing more unit tests for your Go packages.")
    if vet_issues_count > 0:
        recommendations.append(f"Found {vet_issues_count} issues with 'go vet'. Run 'go vet ./...' and fix the reported warnings to ensure code correctness.")
        recommendations.append(f"Vet Output:\n{vet_out}")
    if score == 100:
        recommendations.append("Backend code looks great! Keep up the good work.")

    return {
        "files": int(go_files),
        "loc": int(go_loc),
        "coverage": coverage,
        "vet_issues": vet_issues_count,
        "score": round(score, 1),
        "recommendations": recommendations
    }

def analyze_frontend():
    print("Analyzing Frontend...")
    # Count HTML/CSS/JS files and LOC
    html_files = run_cmd("find static/templates -name '*.html' | wc -l")
    html_loc = run_cmd("find static/templates -name '*.html' -exec cat {} + | wc -l")
    css_files = run_cmd("find static/css -name '*.css' | wc -l")
    css_loc = run_cmd("find static/css -name '*.css' -exec cat {} + | wc -l")
    
    total_files = int(html_files) + int(css_files)
    total_loc = int(html_loc) + int(css_loc)
    
    # Check for inline styles and scripts in HTML
    inline_styles_count = int(run_cmd("grep -r 'style=' static/templates | wc -l"))
    inline_styles_files = run_cmd("grep -rl 'style=' static/templates")
    
    style_tags_count = int(run_cmd("grep -r '<style>' static/templates | wc -l"))
    style_tags_files = run_cmd("grep -rl '<style>' static/templates")
    
    script_tags_count = int(run_cmd("grep -r '<script>' static/templates | wc -l"))
    script_tags_files = run_cmd("grep -rl '<script>' static/templates")
    
    # Calculate score (base 100, penalize inline styles and embedded scripts/styles)
    penalty = (inline_styles_count * 0.5) + (style_tags_count * 2) + (script_tags_count * 1)
    score = max(0, min(100, 100 - penalty))
    
    recommendations = []
    if inline_styles_count > 0:
        files = [f.replace('static/templates/', '') for f in inline_styles_files.split('\n') if f]
        recommendations.append(f"Found {inline_styles_count} inline styles (style=\"...\"). Inline styles make CSS hard to maintain and override. Move these to static/css/style.css using utility classes. Affected files: {', '.join(files)}")
    
    if style_tags_count > 0:
        files = [f.replace('static/templates/', '') for f in style_tags_files.split('\n') if f]
        recommendations.append(f"Found {style_tags_count} embedded <style> tags. For better caching and separation of concerns, move CSS to static/css/style.css. Affected files: {', '.join(files)}")
        
    if script_tags_count > 0:
        files = [f.replace('static/templates/', '') for f in script_tags_files.split('\n') if f]
        recommendations.append(f"Found {script_tags_count} embedded <script> tags. Consider moving complex JavaScript logic to external .js files in a static/js/ directory to improve maintainability and enable Content Security Policy (CSP). Affected files: {', '.join(files)}")

    if score == 100:
        recommendations.append("Frontend code looks clean! Good separation of concerns.")

    return {
        "files": total_files,
        "loc": total_loc,
        "inline_styles": inline_styles_count,
        "style_tags": style_tags_count,
        "script_tags": script_tags_count,
        "score": round(score, 1),
        "recommendations": recommendations
    }

def generate_html(backend_data, frontend_data):
    total_score = round((backend_data['score'] + frontend_data['score']) / 2, 1)
    
    html = f"""<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>IT Tools - Code Quality Report</title>
    <style>
        :root {{
            --bg: #0f172a;
            --card: #1e293b;
            --text: #f1f5f9;
            --text-muted: #94a3b8;
            --primary: #3b82f6;
            --success: #10b981;
            --warning: #f59e0b;
            --danger: #ef4444;
            --border: #334155;
        }}
        body {{
            font-family: system-ui, -apple-system, sans-serif;
            background-color: var(--bg);
            color: var(--text);
            margin: 0;
            padding: 40px 20px;
            line-height: 1.5;
        }}
        .container {{
            max-width: 800px;
            margin: 0 auto;
        }}
        .header {{
            text-align: center;
            margin-bottom: 40px;
        }}
        .total-score {{
            font-size: 48px;
            font-weight: bold;
            color: {get_color(total_score)};
        }}
        .tabs {{
            display: flex;
            border-bottom: 1px solid var(--border);
            margin-bottom: 20px;
        }}
        .tab {{
            padding: 12px 24px;
            cursor: pointer;
            color: var(--text-muted);
            border-bottom: 2px solid transparent;
            font-weight: 500;
        }}
        .tab.active {{
            color: var(--primary);
            border-bottom-color: var(--primary);
        }}
        .panel {{
            display: none;
            background: var(--card);
            border: 1px solid var(--border);
            border-radius: 8px;
            padding: 24px;
        }}
        .panel.active {{
            display: block;
        }}
        .metric {{
            display: flex;
            justify-content: space-between;
            padding: 12px 0;
            border-bottom: 1px solid var(--border);
        }}
        .metric:last-child {{
            border-bottom: none;
        }}
        .metric-label {{
            color: var(--text-muted);
        }}
        .metric-value {{
            font-weight: bold;
        }}
        .score-badge {{
            display: inline-block;
            padding: 4px 12px;
            border-radius: 999px;
            font-weight: bold;
            background: rgba(255,255,255,0.1);
        }}
        .recommendations {{
            margin-top: 24px;
            padding-top: 24px;
            border-top: 1px solid var(--border);
        }}
        .recommendations h3 {{
            margin-top: 0;
            color: var(--text);
        }}
        .recommendation-item {{
            background: rgba(255,255,255,0.05);
            padding: 12px;
            border-radius: 6px;
            margin-bottom: 12px;
            border-left: 4px solid var(--warning);
            font-size: 14px;
        }}
        .recommendation-item pre {{
            background: var(--bg);
            padding: 8px;
            border-radius: 4px;
            overflow-x: auto;
            font-size: 12px;
            margin-top: 8px;
        }}
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>Code Quality Report</h1>
            <div class="total-score">{total_score}/100</div>
            <p style="color: var(--text-muted)">Overall Project Score</p>
        </div>

        <div class="tabs">
            <div class="tab active" onclick="switchTab('frontend', this)">Frontend</div>
            <div class="tab" onclick="switchTab('backend', this)">Backend</div>
        </div>

        <div id="frontend" class="panel active">
            <div style="display: flex; justify-content: space-between; align-items: center; margin-bottom: 20px;">
                <h2>Frontend Metrics</h2>
                <div class="score-badge" style="color: {get_color(frontend_data['score'])}">{frontend_data['score']}/100</div>
            </div>
            <div class="metric">
                <span class="metric-label">Total Files (HTML/CSS)</span>
                <span class="metric-value">{frontend_data['files']}</span>
            </div>
            <div class="metric">
                <span class="metric-label">Lines of Code</span>
                <span class="metric-value">{frontend_data['loc']}</span>
            </div>
            <div class="metric">
                <span class="metric-label">Inline Styles (style="")</span>
                <span class="metric-value" style="color: {get_color(100 - frontend_data['inline_styles']*2)}">{frontend_data['inline_styles']}</span>
            </div>
            <div class="metric">
                <span class="metric-label">Embedded Style Tags</span>
                <span class="metric-value" style="color: {get_color(100 - frontend_data['style_tags']*5)}">{frontend_data['style_tags']}</span>
            </div>
            <div class="metric">
                <span class="metric-label">Embedded Script Tags</span>
                <span class="metric-value" style="color: {get_color(100 - frontend_data['script_tags']*5)}">{frontend_data['script_tags']}</span>
            </div>
            
            <div class="recommendations">
                <h3>Recommendations</h3>
                {''.join(f'<div class="recommendation-item">{r}</div>' for r in frontend_data['recommendations'])}
            </div>
        </div>

        <div id="backend" class="panel">
            <div style="display: flex; justify-content: space-between; align-items: center; margin-bottom: 20px;">
                <h2>Backend Metrics</h2>
                <div class="score-badge" style="color: {get_color(backend_data['score'])}">{backend_data['score']}/100</div>
            </div>
            <div class="metric">
                <span class="metric-label">Total Go Files</span>
                <span class="metric-value">{backend_data['files']}</span>
            </div>
            <div class="metric">
                <span class="metric-label">Lines of Code</span>
                <span class="metric-value">{backend_data['loc']}</span>
            </div>
            <div class="metric">
                <span class="metric-label">Test Coverage</span>
                <span class="metric-value" style="color: {get_color(backend_data['coverage'])}">{backend_data['coverage']}%</span>
            </div>
            <div class="metric">
                <span class="metric-label">Go Vet Issues</span>
                <span class="metric-value" style="color: {get_color(100 - backend_data['vet_issues']*10)}">{backend_data['vet_issues']}</span>
            </div>
            
            <div class="recommendations">
                <h3>Recommendations</h3>
                {''.join(f'<div class="recommendation-item">{r.replace(chr(10), "<br>")}</div>' if not r.startswith('Vet Output:') else f'<div class="recommendation-item"><pre>{r}</pre></div>' for r in backend_data['recommendations'])}
            </div>
        </div>
    </div>

    <script>
        function switchTab(tabId, element) {{
            document.querySelectorAll('.tab').forEach(t => t.classList.remove('active'));
            document.querySelectorAll('.panel').forEach(p => p.classList.remove('active'));
            
            if (element) {{
                element.classList.add('active');
            }} else {{
                document.querySelector(`.tab[onclick="switchTab('${{tabId}}', this)"]`).classList.add('active');
            }}
            document.getElementById(tabId).classList.add('active');
        }}
    </script>
</body>
</html>"""
    return html

def get_color(score):
    if score >= 80: return "var(--success)"
    if score >= 60: return "var(--warning)"
    return "var(--danger)"

if __name__ == "__main__":
    backend_data = analyze_backend()
    frontend_data = analyze_frontend()
    
    html_content = generate_html(backend_data, frontend_data)
    
    output_path = "/Users/mpons/git-repo/ponsmauro/it-tools-code-quality-gemini.html"
    with open(output_path, "w") as f:
        f.write(html_content)
        
    print(f"Report generated successfully at: {output_path}")
